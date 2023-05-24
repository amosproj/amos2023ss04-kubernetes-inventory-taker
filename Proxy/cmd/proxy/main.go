package main

import (
	"log"
	"time"

	. "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/cluster"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	var deploymentLister, namespaceLister, podLister, serviceLister cache.Indexer

	DBpool := SetupDBConnection()

	externalConfig := ReadExternalConfig()
	log.Printf("kubeconfig path is set to \"%s\"\n", externalConfig.KubeconfigPath)

	clientset := CreateClientSet(externalConfig.KubeconfigPath)
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	funcs := SetupEventHandlerFuncs(workqueue)

	stopCh := make(chan struct{})
	RegisterEventHandlers(externalConfig.ResourceTypes, informerFactory, funcs, &deploymentLister, &namespaceLister, &podLister, &serviceLister)

	defer close(stopCh)
	informerFactory.Start(stopCh)

	for i := 0; i < 5; i++ {
		go ProcessWorkqueue(DBpool, workqueue, deploymentLister, namespaceLister, podLister, serviceLister)
	}

	// Wait for the workequeue to stop
	<-stopCh
}

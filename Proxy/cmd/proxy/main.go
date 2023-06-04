package main

import (
	"log"
	"time"

	. "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/cluster"
	database "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/persistent"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

func main() {
	var deploymentLister, namespaceLister, podLister, serviceLister cache.Indexer

	dbq := database.SetupDBConnection()

	externalConfig := ReadExternalConfig()
	log.Printf("kubeconfig path is set to \"%s\"\n", externalConfig.KubeconfigPath)
	WriteCluster(externalConfig.KubeconfigPath, dbq)

	clientset := CreateClientSet(externalConfig.KubeconfigPath)
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	funcs := SetupEventHandlerFuncs(workqueue)

	stopCh := make(chan struct{})

	RegisterEventHandlers(externalConfig.ResourceTypes, informerFactory, funcs,
		&deploymentLister, &namespaceLister, &podLister, &serviceLister)

	log.Println(externalConfig.ResourceTypes)

	defer close(stopCh)
	informerFactory.Start(stopCh)

	for i := 0; i < 1; i++ {
		klog.Info("starting worker", "Number", i)
		go ProcessWorkqueue(dbq, workqueue, deploymentLister, namespaceLister, podLister, serviceLister)
	}

	// Wait for the workequeue to stop
	<-stopCh
}

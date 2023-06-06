package main

import (
	"time"

	//nolint:revive,stylecheck
	. "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/cluster"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

func main() {
	bunDB := SetupDBConnection()

	externalConfig := ReadExternalConfig()
	klog.Infof("kubeconfig path is set to \"%s\"\n", externalConfig.KubeconfigPath)
	WriteCluster(externalConfig.KubeconfigPath, bunDB)

	clientset := CreateClientSet(externalConfig.KubeconfigPath)
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	funcs := SetupEventHandlerFuncs(workqueue)

	stopCh := make(chan struct{})
	RegisterEventHandlers(externalConfig.ResourceTypes, informerFactory, funcs)

	defer close(stopCh)
	informerFactory.Start(stopCh)

	for i := 0; i < 1; i++ {
		klog.Info("starting worker ", i)
		go ProcessWorkqueue(bunDB, workqueue)
	}

	// Wait for the workequeue to stop
	<-stopCh
	bunDB.Close()
}

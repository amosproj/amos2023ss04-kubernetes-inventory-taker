package main

import (
	"time"

	cluster "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/cluster"
	config "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/config"
	db "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/setup"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

func main() {
	externalConfig := config.ReadExternalConfig()

	bunDB := db.DBConnection()
	defer bunDB.Close()

	clientset := cluster.CreateClientSet(externalConfig.KubeconfigPath)
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	workqueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	funcs := cluster.SetupEventHandlerFuncs(workqueue)

	stopCh := make(chan struct{})

	cluster.RegisterEventHandlers(externalConfig.ResourceTypes, informerFactory, funcs)

	defer close(stopCh)
	informerFactory.Start(stopCh)

	for i := 0; i < 1; i++ {
		klog.Info("starting worker ", i)

		go cluster.ProcessWorkqueue(bunDB, workqueue)
	}

	// Wait for the workequeue to stop
	<-stopCh
}

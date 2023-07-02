package cluster

import (
	"reflect"
	"time"

	"github.com/uptrace/bun"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
)

type EventType int

const (
	Add EventType = iota
	Update
	Delete
)

type Event struct {
	Type      EventType
	OldObj    interface{}
	Object    interface{}
	timestamp time.Time
}

func CreateClientSet(kubeconfigPath string) *kubernetes.Clientset {
	// creates the connection
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		klog.Fatal(err)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	return clientset
}

// SetupEventHandlerFuncs setups and returns a new event handler funcs.
func SetupEventHandlerFuncs(workqueue workqueue.RateLimitingInterface) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			workqueue.Add(Event{Type: Add, Object: obj, timestamp: time.Now()})
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			workqueue.Add(Event{Type: Update, OldObj: oldObj, Object: newObj, timestamp: time.Now()})
		},
		DeleteFunc: func(obj interface{}) {
			workqueue.Add(Event{Type: Delete, Object: obj, timestamp: time.Now()})
		},
	}
}

func RegisterEventHandlers(resourceTypes []string, informerFactory informers.SharedInformerFactory,
	funcs cache.ResourceEventHandlerFuncs,
) {
	informerMap := map[string]cache.SharedIndexInformer{
		"persistentvolume":      informerFactory.Core().V1().PersistentVolumes().Informer(),
		"persistentvolumeclaim": informerFactory.Core().V1().PersistentVolumeClaims().Informer(),
		"deployment":            informerFactory.Apps().V1().Deployments().Informer(),
		"namespace":             informerFactory.Core().V1().Namespaces().Informer(),
		"node":                  informerFactory.Core().V1().Nodes().Informer(),
		"pod":                   informerFactory.Core().V1().Pods().Informer(),
		"service":               informerFactory.Core().V1().Services().Informer(),
	}

	for _, resourceType := range resourceTypes {
		if informer, ok := informerMap[resourceType]; ok {
			klog.Info("registering event handler for ", resourceType)

			_, err := informer.AddEventHandler(funcs)
			if err != nil {
				klog.Warning("failed to add ", resourceType, " handler: ", err)
			}
		} else {
			klog.Warning("no handler for ", resourceType, " found")
		}
	}
}

func ProcessWorkqueue(bunDB *bun.DB, workqueue workqueue.RateLimitingInterface) {
	for {
		item, shutdown := workqueue.Get()
		if shutdown {
			return
		}
		//nolint:forcetypeassert
		event := item.(Event)
		klog.Info("processing object of type ", reflect.TypeOf(event.Object))

		switch event.Object.(type) {
		case *corev1.PersistentVolume:
			ProcessPersistentVolume(event, bunDB)
		case *corev1.PersistentVolumeClaim:
			ProcessPersistentVolumeClaim(event, bunDB)
		case *appsv1.Deployment:
		case *corev1.Namespace:
		case *corev1.Node:
			ProcessNode(event, bunDB)

		case *corev1.Pod:
			ProcessPod(event, bunDB)

		case *corev1.Service:
			ProcessService(event, bunDB)
		}

		workqueue.Forget(item)
		workqueue.Done(item)
	}
}

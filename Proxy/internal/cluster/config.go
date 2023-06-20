package cluster

import (
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/uptrace/bun"
	"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"k8s.io/utils/strings/slices"
)

type Config struct {
	KubeconfigPath string   `yaml:"kubeconfigPath"`
	ResourceTypes  []string `yaml:"resourceTypes"`
}

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

func ReadExternalConfig() Config {
	var externalConfig Config

	var proxyConfigPath, kubeConfigPath string

	// parse proxy config file name from cmd flags
	// defaults to same directory
	flag.StringVar(&proxyConfigPath, "config", "config.yaml",
		"(optional) proxy configuration, overwrites kubeconfig flag")

	// parse kubernetes config file location from cmd flags
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeConfigPath, "kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeConfigPath, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	yamlFile, err := os.ReadFile(proxyConfigPath)
	if err == nil {
		klog.Info("Parsing proxy config ", "with file path ", proxyConfigPath)

		err = yaml.Unmarshal(yamlFile, &externalConfig)
		if err != nil {
			klog.Error("Failed umarsheling proxy config file, using empty values ", "config path ", proxyConfigPath)

			externalConfig.KubeconfigPath = ""
			externalConfig.ResourceTypes = []string{}
		}
	} else {
		klog.Error("failed reading proxy config file ", proxyConfigPath, " using empty fields")
	}

	if externalConfig.KubeconfigPath == "" {
		klog.Warning("no Kubeconfig path defined in proxy config, using fallback ", "kubernetes config path ", kubeConfigPath)

		externalConfig.KubeconfigPath = kubeConfigPath
	}

	klog.Info("currently configured values: ", externalConfig)

	return externalConfig
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

//nolint:cyclop,lll
func RegisterEventHandlers(resourceTypes []string, informerFactory informers.SharedInformerFactory, funcs cache.ResourceEventHandlerFuncs) {
	if slices.Contains(resourceTypes, "deployment") {
		klog.Info("found deployment in config")

		deploymentInformer := informerFactory.Apps().V1().Deployments().Informer()

		_, err := deploymentInformer.AddEventHandler(funcs)
		if err != nil {
			klog.Warning("failed to add deployment handler: ", err)
		}
	}

	if slices.Contains(resourceTypes, "namespace") {
		klog.Info("found namespace in config")

		namespaceInformer := informerFactory.Core().V1().Namespaces().Informer()

		_, err := namespaceInformer.AddEventHandler(funcs)
		if err != nil {
			klog.Warning("failed to add namespace handler: ", err)
		}
	}

	if slices.Contains(resourceTypes, "node") {
		klog.Info("found node in config")

		nodeInformer := informerFactory.Core().V1().Nodes().Informer()

		_, err := nodeInformer.AddEventHandler(funcs)
		if err != nil {
			klog.Warning("failed to add node handler: ", err)
		}
	}

	if slices.Contains(resourceTypes, "pod") {
		klog.Info("found pod in config")

		podInformer := informerFactory.Core().V1().Pods().Informer()

		_, err := podInformer.AddEventHandler(funcs)
		if err != nil {
			klog.Warning("failed to add pod handler: ", err)
		}
	}

	if slices.Contains(resourceTypes, "service") {
		klog.Info("found service in config")

		serviceInformer := informerFactory.Core().V1().Services().Informer()

		_, err := serviceInformer.AddEventHandler(funcs)
		if err != nil {
			klog.Warning("failed to add service handler: ", err)
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
		klog.Info("processing object of type", reflect.TypeOf(event.Object))

		switch event.Object.(type) {
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

package cluster

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
	timestamp string
}

// SetupDBConnection setup database connection.
func SetupDBConnection() *pgxpool.Pool {
	// example config string: user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Println("DB_USER environment variable is not set. Trying dbUser = postgres")
		dbUser = "postgres"
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Println("DB_PASSWORD environment variable is not set. Trying dbPassword = example")
		dbPassword = "example"
	}

	configDB, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=localhost port=5432 dbname=postgres pool_max_conns=10", dbUser, dbPassword))
	if err != nil {
		log.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}

func ReadExternalConfig() Config {
	var externalConfig Config

	// parse proxy config file name from cmd flags
	// defaults to same directory
	proxyConfigFile := flag.String("config", "config.yaml",
		"(optional) proxy configuration, overwrites kubeconfig flag")

	yamlFile, err := os.ReadFile(*proxyConfigFile)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, &externalConfig)
		if err != nil {
			externalConfig.KubeconfigPath = ""
			externalConfig.ResourceTypes = []string{}
		}
	}

	if externalConfig.KubeconfigPath != "" {
		return externalConfig
	}

	// parse kubernetes config file location from cmd flags
	if home := homedir.HomeDir(); home != "" {
		externalConfig.KubeconfigPath = *flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		externalConfig.KubeconfigPath = *flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

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
			workqueue.Add(Event{Type: Add, Object: obj, timestamp: time.Now().Format("2006-01-02 15:04:05.000")})
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			workqueue.Add(Event{Type: Update, OldObj: oldObj, Object: newObj, timestamp: time.Now().Format("2006-01-02 15:04:05.000")})
		},
		DeleteFunc: func(obj interface{}) {
			workqueue.Add(Event{Type: Delete, Object: obj, timestamp: time.Now().Format("2006-01-02 15:04:05.000")})
		},
	}
}

func RegisterEventHandlers(resourceTypes []string, informerFactory informers.SharedInformerFactory, funcs cache.ResourceEventHandlerFuncs, deploymentLister, namespaceLister, podLister, serviceLister *cache.Indexer) {
	if slices.Contains(resourceTypes, "deployment") {
		fmt.Println("contains deployment")
		deploymentInformer := informerFactory.Apps().V1().Deployments().Informer()
		*deploymentLister = deploymentInformer.GetIndexer()
		deploymentInformer.AddEventHandler(funcs)
	}
	if slices.Contains(resourceTypes, "namespace") {
		fmt.Println("contains namespace")
		namespaceInformer := informerFactory.Core().V1().Namespaces().Informer()
		*namespaceLister = namespaceInformer.GetIndexer()
		namespaceInformer.AddEventHandler(funcs)
	}
	if slices.Contains(resourceTypes, "node") {
		fmt.Println("contains node")
		nodeInformer := informerFactory.Core().V1().Nodes().Informer()
		nodeInformer.AddEventHandler(funcs)
	}
	if slices.Contains(resourceTypes, "pod") {
		fmt.Println("contains pod")
		podInformer := informerFactory.Core().V1().Pods().Informer()
		*podLister = podInformer.GetIndexer()
		podInformer.AddEventHandler(funcs)
	}
	if slices.Contains(resourceTypes, "service") {
		fmt.Println("contains service")
		serviceInformer := informerFactory.Core().V1().Services().Informer()
		*serviceLister = serviceInformer.GetIndexer()
		serviceInformer.AddEventHandler(funcs)
	}
}

func ProcessWorkqueue(DBpool *pgxpool.Pool, workqueue workqueue.RateLimitingInterface, deploymentLister, namespaceLister, podLister, serviceLister cache.Indexer) {
	for {
		item, shutdown := workqueue.Get()
		if shutdown {
			return
		}
		event := item.(Event)

		switch event.Object.(type) {
		case *appsv1.Deployment:

		case *corev1.Namespace:
		case *corev1.Node:
			ProcessNode(event, DBpool)

		case *corev1.Pod:
			ProcessPod(event, DBpool)

		case *corev1.Service:
		}

		workqueue.Forget(item)
		workqueue.Done(item)
	}
}

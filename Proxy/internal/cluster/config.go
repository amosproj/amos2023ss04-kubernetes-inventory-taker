package cluster

import (
	"database/sql"
	"flag"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

const (
	timeoutSeconds     = 5 // for db
	dialTimeoutSeconds = 5 // for db
)

// SetupDBConnection setup database connection.
func SetupDBConnection() *bun.DB {
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		klog.Warning("DB_USER environment variable is not set. Trying dbUser = postgres")

		dbUser = "postgres"
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		klog.Warning("DB_PASSWORD environment variable is not set. Trying dbPassword = example")

		dbPassword = "example"
	}

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5432"),

		pgdriver.WithUser(dbUser),
		pgdriver.WithPassword(dbPassword),
		pgdriver.WithDatabase("postgres"),
		pgdriver.WithInsecure(true),
		pgdriver.WithTimeout(timeoutSeconds*time.Second),
		pgdriver.WithDialTimeout(dialTimeoutSeconds*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)

	db := bun.NewDB(sqldb, pgdialect.New())
	return db
}

func ReadExternalConfig() Config {
	var externalConfig Config

	// parse proxy config file name from cmd flags
	// defaults to same directory
	proxyConfigFile := flag.String("config", "../../config.yaml",
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

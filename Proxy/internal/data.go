package data

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Clientset is used to load the kubeconfig and setup the clientset to be used for api calls.
func Clientset() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build kubeconfig: %w", err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset from kubeconfig: %w", err)
	}

	return clientset, nil
}

// Pods is a basic wrapper to get a list of Pods on the current cluster.
func Pods(clientset *kubernetes.Clientset) (*v1.PodList, error) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Pods: %w", err)
	}

	return pods, nil
}

func Nodes(clientset *kubernetes.Clientset) (*v1.NodeList, error) {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Nodes: %w", err)
	}

	return nodes, nil
}

func Services(clientset *kubernetes.Clientset) (*v1.ServiceList, error) {
	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Services: %w", err)
	}

	return services, nil
}

func ClusterInfo(clientset *kubernetes.Clientset) (*version.Info, error) {
	clusterInfo, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve server version: %w", err)
	}

	return clusterInfo, nil
}

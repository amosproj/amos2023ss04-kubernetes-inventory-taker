package main

import (
	"log"

	data "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal"
)

func main() {
	log.Printf("Starting Proxy...")

	clientset, err := data.Clientset()
	if err != nil {
		panic(err.Error())
	}

	pods, err := data.Pods(clientset)
	if err != nil {
		panic(err.Error())
	}

	clusterInfo, err := data.ClusterInfo(clientset)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Cluster Version: %s\n", clusterInfo.GitVersion)
	log.Printf("Cluster Major Version: %s\n\n", clusterInfo.Major)

	log.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, pod := range pods.Items {
		log.Printf("\tPod %s: %s\n", pod.GetUID(), pod.GetName())

		for _, container := range pod.Spec.Containers {
			log.Printf("\t|> Container %s\n", container.Name)
		}
	}

	nodes, err := data.Nodes(clientset)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("There are %d nodes in the cluster\n", len(nodes.Items))

	for _, node := range nodes.Items {
		log.Printf("\tNode %s: %s\n", node.GetUID(), node.GetName())
	}

	log.Printf("Exiting proxy ...\n")
}

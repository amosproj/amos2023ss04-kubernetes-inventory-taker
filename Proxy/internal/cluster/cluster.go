package cluster

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	database "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/persistent"
	"gopkg.in/yaml.v2"
)

type KubeConfig struct {
	CurrentContext string `yaml:"current-context"`
}

func WriteCluster(kubeconfigPath string, db *database.Queries) {
	// Read the kubeconfig file
	data, err := os.ReadFile(kubeconfigPath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the kubeconfig file
	var config KubeConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// Print the current context
	fmt.Println("Current context:", config.CurrentContext)
	var clusterParams database.UpdateClusterParams
	clusterParams.ClusterEventID = 0
	clusterParams.ClusterID = 0
	clusterParams.Timestamp.Scan(time.Now())
	clusterParams.Name = config.CurrentContext

	// Insert the Cluster into the database
	err = db.UpdateCluster(context.Background(), clusterParams)
	if err != nil {
		log.Println(err)
	}
}

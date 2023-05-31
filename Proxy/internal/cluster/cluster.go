package cluster

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"gopkg.in/yaml.v2"
)

type KubeConfig struct {
	CurrentContext string `yaml:"current-context"`
}

type Cluster struct {
	bun.BaseModel `bun:"table:Cluster"`
	NodeEventID   int       `bun:"cluster_event_id,type:integer,pk"`
	ClusterID     int       `bun:"cluster_id,type:integer"`
	Timestamp     time.Time `bun:"timestamp,type:timestamp,notnull"`
	name          string    `bun:"name,type:text"`
}

func WriteCluster(kubeconfigPath string, db *bun.DB) {
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

	clusterDB := Cluster{
		BaseModel:   schema.BaseModel{},
		NodeEventID: 0,
		ClusterID:   0,
		Timestamp:   time.Now(),
		name:        config.CurrentContext,
	}

	// Insert the Cluster into the database
	_, err = db.NewInsert().Model(&clusterDB).Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

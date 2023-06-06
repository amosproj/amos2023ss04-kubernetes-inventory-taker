package cluster

import (
	"context"
	"os"
	"time"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"gopkg.in/yaml.v2"
	"k8s.io/klog"
)

type KubeConfig struct {
	//nolint:tagliatelle
	CurrentContext string `yaml:"current-context"`
}

func WriteCluster(kubeconfigPath string, bunDB *bun.DB) {
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

	clusterDB := model.Cluster{
		BaseModel: schema.BaseModel{},

		ClusterID: 0,
		Timestamp: time.Now(),
		Name:      config.CurrentContext,
	}

	// Insert the Cluster into the database
	_, err = bunDB.NewInsert().Model(&clusterDB).Exec(context.Background())
	if err != nil {
		klog.Warning(err)
	}
}

package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessService(event Event, bunDB *bun.DB) {
	//nolint:forcetypeassert
	serviceNew := event.Object.(*corev1.Service)

	//nolint:forcetypeassert
	if event.Type == Update && event.OldObj.(*corev1.Service).ResourceVersion == serviceNew.ResourceVersion {
		return
	}

	// Convert the service's ports to a slice of strings
	ports := []string{}
	for _, port := range serviceNew.Spec.Ports {
		ports = append(ports, fmt.Sprintf("%d/%s", port.Port, port.Protocol))
	}

	// Convert the service's labels to a slice of strings
	labels := []string{}
	for key, value := range serviceNew.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}

	// Get json of node
	jsonData, err := json.Marshal(serviceNew)
	if err != nil {
		klog.Error("Error converting Node to JSON:", err)
		return
	}

	// Create a Service struct from the corev1.Service
	serviceDB := &model.Service{
		BaseModel: schema.BaseModel{},

		Name:              serviceNew.Name,
		Namespace:         serviceNew.Namespace,
		Timestamp:         event.timestamp,
		Labels:            labels,
		CreationTimestamp: serviceNew.CreationTimestamp.Time,
		Ports:             ports,
		ExternalIPs:       serviceNew.Spec.ExternalIPs,
		ClusterIP:         serviceNew.Spec.ClusterIP,
		Data:              string(jsonData),
	}

	// Insert the Service into the database
	_, err = bunDB.NewInsert().Model(serviceDB).Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

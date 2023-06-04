package cluster

import (
	"context"
	"fmt"
	"time"

	database "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/persistent"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

func ProcessService(event Event, db *database.Queries) {
	service := event.Object.(*corev1.Service)

	if event.Type == Update && event.OldObj.(*corev1.Service).ResourceVersion == service.ResourceVersion {
		return
	}

	// Convert the service's ports to a slice of strings
	var ports []string
	for _, port := range service.Spec.Ports {
		ports = append(ports, fmt.Sprintf("%d/%s", port.Port, port.Protocol))
	}

	// Convert the service's labels to a slice of strings
	var labels []string
	for key, value := range service.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", key, value))
	}

	// Create a Service struct from the corev1.Service
	var serviceParams database.UpdateServiceParams

	serviceParams.Name = service.Name
	serviceParams.Namespace = service.Namespace
	serviceParams.Labels = labels
	serviceParams.CreationTimestamp.Scan(service.CreationTimestamp.Time)
	serviceParams.Timestamp.Scan(time.Now())
	serviceParams.Ports = ports
	serviceParams.ExternalIps = service.Spec.ExternalIPs
	serviceParams.ClusterIp = service.Spec.ClusterIP

	// Insert the Service into the database
	err := db.UpdateService(context.Background(), serviceParams)
	if err != nil {
		klog.Fatal(err)
	}
}

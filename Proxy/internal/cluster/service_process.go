package cluster

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
)

type Service struct {
	bun.BaseModel     `bun:"Service"`
	Name              string    `bun:"name,type:text,notnull,pk"`
	Namespace         string    `bun:"namespace,type:text,notnull,pk"`
	Timestamp         time.Time `bun:"timestamp,type:timestamp,notnull"`
	Labels            []string  `bun:"labels,type:varchar[],array,notnull"`
	CreationTimestamp time.Time `bun:"creation_timestamp,type:timestamp,notnull"`
	Ports             []string  `bun:"ports,type:varchar[],array,notnull"`
	ExternalIPs       []string  `bun:"external_ips,type:varchar[],array"`
	ClusterIP         string    `bun:"cluster_ip,type:text,notnull"`
}

func ProcessService(event Event, db *bun.DB) {
	service := event.Object.(*corev1.Service)

	//ignore update event with unchanged resource version
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
	s := &Service{
		Name:              service.Name,
		Namespace:         service.Namespace,
		Labels:            labels,
		CreationTimestamp: service.CreationTimestamp.Time,
		Timestamp:         time.Now(), // replace with the actual value
		Ports:             ports,
		ExternalIPs:       service.Spec.ExternalIPs,
		ClusterIP:         service.Spec.ClusterIP,
	}

	// Insert the Service into the database
	_, err := db.NewInsert().Model(s).Exec(context.Background())
	if err != nil {
		log.Println(err)
	}
}

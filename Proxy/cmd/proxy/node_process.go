package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5/pgxpool"
	corev1 "k8s.io/api/core/v1"
)

func ProcessNode(event Event, DBpool *pgxpool.Pool) {
	node := event.Object.(*corev1.Node)

	var internalIP, externalIP string

	for _, address := range node.Status.Addresses {
		if address.Type == corev1.NodeInternalIP {
			internalIP = address.Address
		} else if address.Type == corev1.NodeExternalIP {
			externalIP = address.Address
		}
	}

	sqlStatementNode := `INSERT INTO "Node" ("node_event_id", "node_id", "timestamp", "cluster_id", "name", "ip_address_internal", "ip_address_external","status") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := DBpool.Exec(context.Background(), sqlStatementNode, rand.Int31(), node.UID, event.timestamp, nil, node.Name, internalIP, externalIP, node.Status)
	if err != nil {
		fmt.Printf("failed to execute node query: %v", err)
	}
}

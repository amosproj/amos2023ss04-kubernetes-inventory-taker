package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	corev1 "k8s.io/api/core/v1"
)

func ProcessPod(event Event, DBpool *pgxpool.Pool) {
	pod := event.Object.(*corev1.Pod)
	// if event.OldObj != nil && event.OldObj.(*corev1.Pod).ObjectMeta.ResourceVersion == event.Object.(*corev1.Pod).ObjectMeta.ResourceVersion {
	// 	fmt.Println("duplicate")
	// }
	fmt.Print(pod.ObjectMeta.ResourceVersion)
	sqlStatementPod := `INSERT INTO "Pod" ("pod_resource_version", "pod_id", "timestamp", "cluster_id", "node_id", "name", "namespace","status", "stamp2") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	sqlStatementContainer := `INSERT INTO "Container" ("container_event_id", "container_id", "timestamp", "pod_id", "name", "image", "status", "ports") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	switch {
	case event.Type == Add || event.Type == Update || event.Type == Delete:

		containers := pod.Spec.Containers
		for i := range containers {
			fmt.Println("found container", containers[i].Name, "with image", containers[i].Image)
			status := "undefined"
			if pod.Status.ContainerStatuses[i].State.Waiting != nil {
				status = "Waiting"
			} else if pod.Status.ContainerStatuses[i].State.Running != nil {
				status = "Running"
			} else if pod.Status.ContainerStatuses[i].State.Terminated != nil {
				status = "Terminated"
			}
			_, err := DBpool.Exec(context.Background(), sqlStatementContainer, rand.Int31(), rand.Int31(), event.timestamp, pod.ObjectMeta.UID, containers[i].Name, containers[i].Image, status, "containers[i].Ports")
			if err != nil {
				fmt.Printf("failed to execute container query: %v", err)
			}
		}
		phase := pod.Status.Phase
		if event.Type == Delete {
			phase = "Deleted"
		}

		var podOld (*corev1.Pod)

		if event.OldObj != nil {
			podOld := event.OldObj.(*corev1.Pod)
			log.Printf("Old Pod Object was present with ID %s\n", podOld.UID)
		}

		if podOld == nil || podOld.ObjectMeta.ResourceVersion != pod.ObjectMeta.ResourceVersion {
			log.Printf("pod with uid %s was updated\n", pod.UID)
			_, err := DBpool.Exec(context.Background(), sqlStatementPod, pod.ObjectMeta.ResourceVersion, pod.ObjectMeta.UID,
				event.timestamp, 0, pod.Spec.NodeName, pod.Name, pod.Namespace, phase, time.Now().Format("2006-01-02 15:04:05.000"))
			if err != nil {
				fmt.Printf("failed to execute query: %v", err)
			}
		} else {
			log.Printf("pod with uid %s was NOT updated\n", pod.UID)
		}

		// podJSON, _ := json.Marshal(pod)
		// fmt.Println(string(podJSON))
	}
}

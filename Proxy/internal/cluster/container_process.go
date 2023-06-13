package cluster

import (
	"context"
	"fmt"
	"strings"
	"time"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/model"
	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessContainer(pod *corev1.Pod, bunDB *bun.DB, timestamp time.Time) {
	containerStatuses := pod.Status.ContainerStatuses

	for idx := range containerStatuses {
		containerSpec := getSpec(*pod, containerStatuses[idx].Name)
		Ports := containerSpec.Ports

		contaierDB := &model.Container{
			Timestamp:   timestamp,
			ContainerID: containerStatuses[idx].ContainerID,
			PodID:       string(pod.UID),
			Name:        containerStatuses[idx].Name,
			Image:       containerStatuses[idx].Image,
			Status:      getContainerState(containerStatuses[idx].State),
			Ports:       formatPorts(Ports),
		}
		// Insert the Container into the database
		_, err := bunDB.NewInsert().Model(contaierDB).Exec(context.Background())
		if err != nil {
			klog.Error(err)
		}
	}
}

func getContainerState(state corev1.ContainerState) string {
	var stateString string

	switch {
	case state.Running != nil:
		stateString = "Running"
	case state.Waiting != nil:
		stateString = "Waiting"
	case state.Terminated != nil:
		stateString = "Terminated"
	}

	return stateString
}

func getSpec(pod corev1.Pod, containerName string) *corev1.Container {
	containers := pod.Spec.Containers

	for idx := range containers {
		if containers[idx].Name == containerName {
			return &containers[idx]
		}
	}

	return nil
}

func formatPorts(ports []corev1.ContainerPort) string {
	portStrings := make([]string, 0, len(ports)) // preallocate the slice with capacity
	for _, port := range ports {
		portStrings = append(portStrings, fmt.Sprintf("%d/%s", port.ContainerPort, port.Protocol))
	}

	portsString := strings.Join(portStrings, ", ")

	return portsString
}

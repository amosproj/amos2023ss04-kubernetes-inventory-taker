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

	for _, cStatus := range containerStatuses {
		containerSpec := getSpec(*pod, cStatus.Name)
		ports := containerSpec.Ports
		state := scanContainerState(&cStatus.State)
		lastState := scanContainerState(&cStatus.LastTerminationState)

		containerDB := &model.Container{
			Timestamp:    timestamp,
			ContainerID:  cStatus.ContainerID,
			PodID:        string(pod.UID),
			Name:         cStatus.Name,
			Image:        cStatus.Image,
			Status:       state.Kind,
			Ports:        formatPorts(ports),
			ImageID:      cStatus.ImageID,
			Ready:        cStatus.Ready,
			RestartCount: int(cStatus.RestartCount),
			Started:      *cStatus.Started,
		}

		if _, err := bunDB.NewInsert().Model(state).Exec(context.Background()); err != nil {
			klog.Error(err)
		}

		containerDB.StateID = state.ID

		if lastState != nil {
			if _, err := bunDB.NewInsert().Model(lastState).Exec(context.Background()); err != nil {
				klog.Error(err)
			}

			containerDB.LastStateID = lastState.ID
		}

		// Insert the Container into the database
		if _, err := bunDB.NewInsert().Model(containerDB).Exec(context.Background()); err != nil {
			klog.Error(err)
		}
	}
}

func scanContainerState(state *corev1.ContainerState) *model.ContainerState {
	if state == nil {
		return nil
	}

	var ret model.ContainerState

	switch {
	case state.Running != nil:
		ret.Kind = "Running"
		ret.StartedAt = state.Running.StartedAt.Time
	case state.Waiting != nil:
		ret.Kind = "Waiting"
		ret.Reason = state.Waiting.Reason
		ret.Message = state.Waiting.Message
	case state.Terminated != nil:
		ret.Kind = "Terminated"
		ret.ContainerID = state.Terminated.ContainerID
		ret.Reason = state.Terminated.Reason
		ret.Message = state.Terminated.Message
		ret.FinishedAt = state.Terminated.FinishedAt.Time
		ret.StartedAt = state.Terminated.StartedAt.Time
		ret.ExitCode = int(state.Terminated.ExitCode)
		ret.Signal = int(state.Terminated.Signal)
	}

	return &ret
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

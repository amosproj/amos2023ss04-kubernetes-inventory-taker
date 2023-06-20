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

// Process all contenders from one Pod.
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

		insertVolumeMounts(bunDB, containerSpec.VolumeMounts, containerDB.ID)
		insertVolumeDevices(bunDB, containerSpec.VolumeDevices, containerDB.ID)
		insertContainerPorts(bunDB, containerSpec.Ports, containerDB.ID)
	}
}

// ContainerState can be one of 3 types, all are mapped to one table.
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

func insertVolumeDevices(bunDB *bun.DB, volDevices []corev1.VolumeDevice, containerID int) {
	for _, volDev := range volDevices {
		model := &model.VolumeDevice{
			ContainerID: containerID,
			DevicePath:  volDev.DevicePath,
			Name:        volDev.Name,
		}

		if _, err := bunDB.NewInsert().Model(model).Exec(context.Background()); err != nil {
			klog.Error(err)
		}
	}
}

func insertVolumeMounts(bunDB *bun.DB, volMounts []corev1.VolumeMount, containerID int) {
	for _, volMount := range volMounts {
		model := &model.VolumeMount{
			ContainerID: containerID,
			MountPath:   volMount.MountPath,
			Name:        volMount.Name,
			ReadOnly:    volMount.ReadOnly,
			SubPath:     volMount.SubPath,
			SubPathExpr: volMount.SubPathExpr,
		}

		if volMount.MountPropagation != nil {
			model.MountPropagation = string(*volMount.MountPropagation)
		}

		if _, err := bunDB.NewInsert().Model(model).Exec(context.Background()); err != nil {
			klog.Error(err)
		}
	}
}

func insertContainerPorts(bunDB *bun.DB, cPorts []corev1.ContainerPort, containerID int) {
	for _, port := range cPorts {
		model := &model.ContainerPort{
			ContainerID:   containerID,
			ContainerPort: int(port.ContainerPort),
			HostIP:        port.HostIP,
			HostPort:      int(port.HostPort),
			Name:          port.Name,
			Protocol:      string(port.Protocol),
		}

		if _, err := bunDB.NewInsert().Model(model).Exec(context.Background()); err != nil {
			klog.Error(err)
		}
	}
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

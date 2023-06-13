package cluster

import (
	"context"
	"time"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/model"
	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessContainer(pod *corev1.Pod, bunDB *bun.DB, timestamp time.Time) {
	statuses := pod.Status.ContainerStatuses

	for idx := range statuses {
		contaierDB := &model.Container{
			Timestamp:   timestamp,
			ContainerID: statuses[idx].ContainerID,
			PodID:       string(pod.UID),
			Name:        statuses[idx].Name,
			Image:       statuses[idx].Image,
			Status:      getContainerState(statuses[idx].State),
			Ports:       "TODO: ports",
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

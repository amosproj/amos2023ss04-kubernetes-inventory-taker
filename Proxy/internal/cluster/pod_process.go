package cluster

import (
	"context"
	"encoding/json"
	"time"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/model"
	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessPod(event Event, bunDB *bun.DB) {
	//nolint:forcetypeassert
	podNew := event.Object.(*corev1.Pod)

	//nolint:forcetypeassert
	if event.Type == Update && event.OldObj.(*corev1.Pod).ResourceVersion == podNew.ResourceVersion {
		return
	}

	insertPod(podNew, bunDB, event.timestamp)

	// This adds the containers inside the pod
	ProcessContainer(podNew, bunDB, event.timestamp)
}

func insertPod(podNew *corev1.Pod, bunDB *bun.DB, eventTimestamp time.Time) {
	jsonData, err := json.Marshal(podNew)
	if err != nil {
		klog.Error(err)
	}

	podStatus := podNew.Status
	podIPs := make([]string, len(podStatus.PodIPs))

	for i, podIP := range podStatus.PodIPs {
		podIPs[i] = podIP.IP
	}

	podStatusDB := &model.PodStatus{
		ID:          0,
		StatusPhase: string(podStatus.Phase),
		HostIP:      podStatus.HostIP,
		PodIP:       podStatus.PodIP,
		PodIPs:      podIPs,
		StartTime:   podStatus.StartTime.Time,
		QOSClass:    string(podStatus.QOSClass),
		// Conditions:  []model.PodStatusCondition{}, //handled by bun
	}

	_, err = bunDB.NewInsert().Model(podStatusDB).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}

	podDB := &model.Pod{
		PodStatusID:        podStatusDB.ID,
		PodResourceVersion: podNew.ResourceVersion,
		PodID:              string(podNew.UID),
		Timestamp:          eventTimestamp,
		NodeName:           podNew.Spec.NodeName,
		Name:               podNew.Name,
		Namespace:          podNew.Namespace,
		StatusPhase:        string(podNew.Status.Phase),
		Data:               string(jsonData),
	}

	_, err = bunDB.NewInsert().Model(podDB).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}

	insertPodStatusConditions(podStatus, podStatusDB.ID, bunDB)
}

func insertPodStatusConditions(podStatus corev1.PodStatus, podStatusID int, bunDB *bun.DB) {
	for _, condition := range podStatus.Conditions {
		podStatusConditionDB := &model.PodStatusCondition{
			PodStatusID:        podStatusID,
			Type:               string(condition.Type),
			Status:             string(condition.Status),
			LastProbeTime:      condition.LastProbeTime.Time,
			LastTransitionTime: condition.LastTransitionTime.Time,
			Reason:             condition.Reason,
			Message:            condition.Message,
		}

		_, err := bunDB.NewInsert().Model(podStatusConditionDB).Exec(context.Background())
		if err != nil {
			klog.Error(err)
		}
	}
}

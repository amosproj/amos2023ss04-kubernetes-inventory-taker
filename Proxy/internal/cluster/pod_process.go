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

	podDB := &model.Pod{
		PodResourceVersion: podNew.ResourceVersion,
		PodID:              string(podNew.UID),
		Timestamp:          eventTimestamp,
		NodeName:           podNew.Spec.NodeName,
		Name:               podNew.Name,
		Namespace:          podNew.Namespace,
		StatusPhase:        string(podNew.Status.Phase),
		Data:               string(jsonData),
		HostIP:             podStatus.HostIP,
		PodIP:              podStatus.PodIP,
		PodIPs:             podIPs,
		StartTime:          packTimestamp(podStatus.StartTime),
		QOSClass:           string(podStatus.QOSClass),
	}

	_, err = bunDB.NewInsert().Model(podDB).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}

	insertPodStatusConditions(podStatus, podDB.ID, bunDB)
	insertPodVolumes(podNew.Spec.Volumes, podDB.ID, bunDB)
}

func insertPodStatusConditions(podStatus corev1.PodStatus, podID int, bunDB *bun.DB) {
	for _, condition := range podStatus.Conditions {
		podStatusConditionDB := &model.PodStatusCondition{
			PodID:              podID,
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

func insertPodVolumes(podVolumes []corev1.Volume, podID int, bunDB *bun.DB) {
	for _, volume := range podVolumes {
		podVolumeDB := &model.PodVolume{
			PodID: podID,
			Name:  volume.Name,
		}

		if volume.VolumeSource.PersistentVolumeClaim == nil {
			podVolumeDB.Type = "other"
		} else {
			podVolumeDB.Type = "pvc"
			podVolumeDB.ClaimName = volume.VolumeSource.PersistentVolumeClaim.ClaimName
			podVolumeDB.ReadOnly = volume.VolumeSource.PersistentVolumeClaim.ReadOnly
		}

		_, err := bunDB.NewInsert().Model(podVolumeDB).Exec(context.Background())
		if err != nil {
			klog.Error(err)
		}
	}
}

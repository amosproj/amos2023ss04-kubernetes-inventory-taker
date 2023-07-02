package cluster

import (
	"context"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/model"
	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessPersistentVolumeClaim(event Event, bunDB *bun.DB) {
	pvcNew, ok := event.Object.(*corev1.PersistentVolumeClaim)
	if !ok {
		klog.Errorf("unexpected type: %T", event.Object)
		return
	}

	// in case of update, check if the resource version is the same
	if event.Type == Update {
		pwOld, ok := event.OldObj.(*corev1.PersistentVolumeClaim)
		if !ok {
			klog.Errorf("unexpected old type: %T", event.OldObj)
			return
		}

		if pwOld.ResourceVersion == pvcNew.ResourceVersion {
			return
		}
	}

	pvcEntry := &model.PersistentVolumeClaim{
		Timestamp:         event.timestamp,
		Labels:            packLabels(pvcNew.Labels),
		Name:              pvcNew.Name,
		Namespace:         pvcNew.Namespace,
		UID:               string(pvcNew.UID),
		CreationTimestamp: pvcNew.CreationTimestamp.Time,
		DeletionTimestamp: packTimestamp(pvcNew.DeletionTimestamp),
		AccessModes:       packAccessModes(pvcNew.Spec.AccessModes),
		StorageClassName:  *pvcNew.Spec.StorageClassName,
		VolumeMode:        string(*pvcNew.Spec.VolumeMode),
		VolumeName:        pvcNew.Spec.VolumeName,
		Capacity:          pvcNew.Status.Capacity.Storage().String(),
		Phase:             string(pvcNew.Status.Phase),
		ResizeStatus:      string(*pvcNew.Status.ResizeStatus),
	}

	// Insert the PersistentVolumeClaim into the database
	_, err := bunDB.NewInsert().Model(pvcEntry).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}

	for _, condition := range pvcNew.Status.Conditions {
		pvcConditionEntry := &model.PersistentVolumeClaimCondition{
			PersistentVolumeClaimID: pvcEntry.ID,
			LastProbeTime:           packTimestamp(&condition.LastProbeTime),
			LastTransitionTime:      packTimestamp(&condition.LastTransitionTime),
			Message:                 condition.Message,
			Reason:                  condition.Reason,
			Status:                  string(condition.Status),
			Type:                    string(condition.Type),
		}

		// Insert the PersistentVolumeClaimCondition into the database
		_, err := bunDB.NewInsert().Model(pvcConditionEntry).Exec(context.Background())
		if err != nil {
			klog.Error(err)
		}
	}
}

package cluster

import (
	"context"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/database/model"
	"github.com/uptrace/bun"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
)

func ProcessPersistentVolume(event Event, bunDB *bun.DB) {
	pvNew, ok := event.Object.(*corev1.PersistentVolume)
	if !ok {
		klog.Errorf("unexpected type: %T", event.Object)
		return
	}

	// in case of update, check if the resource version is the same
	if event.Type == Update {
		pwOld, ok := event.OldObj.(*corev1.PersistentVolume)
		if !ok {
			klog.Errorf("unexpected old type: %T", event.OldObj)
			return
		}

		if pwOld.ResourceVersion == pvNew.ResourceVersion {
			return
		}
	}

	pvEntry := &model.PersistentVolume{
		Timestamp:         event.timestamp,
		Labels:            packLabels(pvNew.Labels),
		Name:              pvNew.Name,
		Namespace:         pvNew.Namespace,
		UID:               string(pvNew.UID),
		CreationTimestamp: pvNew.CreationTimestamp.Time,
		DeletionTimestamp: packTimestamp(pvNew.DeletionTimestamp),
		AccessModes:       packAccessModes(pvNew.Spec.AccessModes),
		Capacity:          pvNew.Spec.Capacity.Storage().String(),
		MountOptions:      pvNew.Spec.MountOptions,
		StorageClassName:  pvNew.Spec.StorageClassName,
		VolumeMode:        string(*pvNew.Spec.VolumeMode),
		Message:           pvNew.Status.Message,
		Phase:             string(pvNew.Status.Phase),
		Reason:            pvNew.Status.Reason,
	}

	// Insert the PersistentVolume into the database
	_, err := bunDB.NewInsert().Model(pvEntry).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}
}

package cluster

import (
	"context"
	"encoding/json"

	model "github.com/amosproj/amos2023ss04-kubernetes-inventory-taker/Proxy/internal/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
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

	jsonData, err := json.Marshal(podNew)
	if err != nil {
		klog.Error("Error converting Node to JSON:", err)
	}

	podDB := &model.Pod{
		BaseModel:          schema.BaseModel{},
		ID:                 0,
		PodResourceVersion: podNew.ResourceVersion,
		PodID:              string(podNew.UID),
		Timestamp:          event.timestamp,
		NodeName:           podNew.Spec.NodeName,
		Name:               podNew.Name,
		Namespace:          podNew.Namespace,
		StatusPhase:        string(podNew.Status.Phase),
		Data:               string(jsonData),
	}

	// Insert the Service into the database
	_, err = bunDB.NewInsert().Model(podDB).Exec(context.Background())
	if err != nil {
		klog.Error(err)
	}
}

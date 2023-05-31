package cluster

import (
	"github.com/uptrace/bun"
	//corev1 "k8s.io/api/core/v1"
)

type Pod struct {
	bun.BaseModel `bun:"table:Cluster"`
	//TODO
}

func ProcessPod(event Event, db *bun.DB) {
	//pod := event.Object.(*corev1.Pod)
	
}
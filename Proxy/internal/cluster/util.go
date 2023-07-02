package cluster

import (
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func packTimestamp(timestamp *v1.Time) time.Time {
	if timestamp == nil {
		return time.Time{}
	}

	return timestamp.Time
}

func packAccessModes(accessModes []corev1.PersistentVolumeAccessMode) []string {
	ret := make([]string, 0, len(accessModes))

	for _, accessMode := range accessModes {
		ret = append(ret, string(accessMode))
	}

	return ret
}

func packLabels(labels map[string]string) []string {
	ret := make([]string, 0, len(labels))

	for k, v := range labels {
		combined := fmt.Sprintf("%s=%s", k, v)
		ret = append(ret, combined)
	}

	return ret
}

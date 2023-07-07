// Package cluster provides abstraction from kubernetes API.
// util.go contains type utility functions for the cluster package.
package cluster

import (
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// packTimestamp handles nil v1.Time pointers as well.
func packTimestamp(timestamp *v1.Time) time.Time {
	if timestamp == nil {
		return time.Time{}
	}

	return timestamp.Time
}

// packAccessModes converts AccessModes to string slice.
func packAccessModes(accessModes []corev1.PersistentVolumeAccessMode) []string {
	ret := make([]string, 0, len(accessModes))

	for _, accessMode := range accessModes {
		ret = append(ret, string(accessMode))
	}

	return ret
}

// packLabels converts labels map to string slice joint with '='.
func packLabels(labels map[string]string) []string {
	ret := make([]string, 0, len(labels))

	for k, v := range labels {
		combined := fmt.Sprintf("%s=%s", k, v)
		ret = append(ret, combined)
	}

	return ret
}

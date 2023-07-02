// Package config handles proxy configuration and command line flag parsing.
package config

import (
	"flag"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

type Config struct {
	KubeconfigPath string   `yaml:"kubeconfigPath"`
	ResourceTypes  []string `yaml:"resourceTypes"`
}

// ReadExternalConfig Parsing commnd line flags and reading external config file with resource types
// to be tracked with informers.
func ReadExternalConfig() Config {
	var externalConfig Config

	var proxyConfigPath, kubeConfigPath string

	// parse proxy config file name from cmd flags
	// defaults to same directory
	flag.StringVar(&proxyConfigPath, "config", "config.yaml",
		"(optional) proxy configuration, overwrites kubeconfig flag")

	// parse kubernetes config file location from cmd flags
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeConfigPath, "kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeConfigPath, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	yamlFile, err := os.ReadFile(proxyConfigPath)
	if err == nil {
		klog.Info("Parsing proxy config ", "with file path ", proxyConfigPath)

		err = yaml.Unmarshal(yamlFile, &externalConfig)
		if err != nil {
			klog.Error("Failed umarsheling proxy config file, using empty values ", "config path ", proxyConfigPath)

			externalConfig.KubeconfigPath = ""
			externalConfig.ResourceTypes = []string{}
		}
	} else {
		klog.Error("failed reading proxy config file ", proxyConfigPath, " using empty fields")
	}

	if externalConfig.KubeconfigPath == "" {
		klog.Warning("no Kubeconfig path defined in proxy config, using fallback ", "kubernetes config path ", kubeConfigPath)

		externalConfig.KubeconfigPath = kubeConfigPath
	}

	klog.Info("currently configured values: ", externalConfig)

	return externalConfig
}

package clusterhelper

import (
	"fmt"
	"path/filepath"

	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type ClusterHelper struct {
	clientset *kubernetes.Clientset
}

func New(kubeconfig string) (*ClusterHelper, error) {
	helper := &ClusterHelper{}
	if kubeconfig == "" {
		// 从本机home目录获取kubeconfig文件
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	// 从 kubeconfig 文件构建配置
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return helper, fmt.Errorf("build config from kubeconfig failed: %s", err)
	}
	// 创建 Kubernetes 客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return helper, fmt.Errorf("create clientset failed: %s", err)
	}
	helper.clientset = clientset
	return helper, nil
}

func (c *ClusterHelper) ClusterServerVersion() (*version.Info, error) {
	// 获取集群信息
	version, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		return version, fmt.Errorf("get cluster server version failed: %s", err)
	}
	return version, nil
}

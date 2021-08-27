package vnode

import (
	"context"
	virtual_kubelet "github.com/laik/vnode/pkg/virtual-kubelet"
	"github.com/laik/vnode/pkg/virtual-kubelet/common"
	"github.com/virtual-kubelet/virtual-kubelet/node"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	// Provider configuration defaults.
	defaultCPUCapacity    = "32"
	defaultMemoryCapacity = "100Gi"
	defaultPodCapacity    = "110"
)

var (
	_ virtual_kubelet.Provider           = (*VirtualNodeProvider)(nil)
	_ virtual_kubelet.PodMetricsProvider = (*VirtualNodeProvider)(nil)
	_ node.PodNotifier                   = (*VirtualNodeProvider)(nil)
)

type VirtualNodeProvider struct {
	*VirtualNodeConfig
}

func NewVirtualNodeProvider(ctx context.Context, initCfg *virtual_kubelet.InitConfig) (*VirtualNodeProvider, error) {
	configures, err := initCfg2VirtualNodeConfig(ctx, initCfg)
	if err != nil {
		return nil, err
	}
	return NewVirtualNodeProviderConfig(configures...)
}

func initCfg2VirtualNodeConfig(ctx context.Context, initCfg *virtual_kubelet.InitConfig) ([]VirtualNodeConfigure, error) {
	configures := []VirtualNodeConfigure{
		OperatingSystem(initCfg.OperatingSystem),
		NodeName(initCfg.NodeName),
		InternalIP(initCfg.InternalIP),
		DaemonEndpointPort(initCfg.DaemonPort),
	}
	var restConfig *rest.Config
	var resetErr error

	if common.InCluster {
		restConfig, resetErr = rest.InClusterConfig()
	} else {
		restConfig, resetErr = getRestConfig(*common.KubeConfig)
	}

	if resetErr != nil {
		return nil, resetErr
	}
	configures = append(configures, ClientSet(kubernetes.NewForConfigOrDie(restConfig)))

	return configures, nil
}

func getRestConfig(path string) (*rest.Config, error) {
	restConfig, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	return restConfig, nil
}

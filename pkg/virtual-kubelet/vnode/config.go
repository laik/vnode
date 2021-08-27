package vnode

import (
	"k8s.io/client-go/kubernetes"
)

type VirtualNodeConfigure func(*VirtualNodeConfig)

type VirtualNodeConfig struct {
	// VirtualNodeConfig contains a node virtual-kubelet's configurable parameters.
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Pods   string `json:"pods,omitempty"`

	nodeName           string
	operatingSystem    string
	internalIP         string
	daemonEndpointPort int32

	clientset *kubernetes.Clientset
}

func NewVirtualNodeProviderConfig(configures ...VirtualNodeConfigure) (*VirtualNodeProvider, error) {
	cfg := &VirtualNodeConfig{
		CPU:    defaultCPUCapacity,
		Memory: defaultMemoryCapacity,
		Pods:   defaultPodCapacity,
	}
	for _, configure := range configures {
		configure(cfg)
	}

	provider := VirtualNodeProvider{
		VirtualNodeConfig: cfg,
	}

	return &provider, nil
}

func NodeName(nodeName string) VirtualNodeConfigure {
	return func(v *VirtualNodeConfig) {
		v.nodeName = nodeName
	}
}

func OperatingSystem(operatingSystem string) VirtualNodeConfigure {
	return func(v *VirtualNodeConfig) {
		v.operatingSystem = operatingSystem
	}
}

func InternalIP(internalIP string) VirtualNodeConfigure {
	return func(v *VirtualNodeConfig) {
		v.internalIP = internalIP
	}
}

func DaemonEndpointPort(daemonEndpointPort int32) VirtualNodeConfigure {
	return func(v *VirtualNodeConfig) {
		v.daemonEndpointPort = daemonEndpointPort
	}
}

func ClientSet(cs *kubernetes.Clientset) VirtualNodeConfigure {
	return func(v *VirtualNodeConfig) {
		v.clientset = cs
	}
}

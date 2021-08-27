package vnode

import (
	"context"
	"github.com/laik/vnode/pkg/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (n VirtualNodeProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	log.G(ctx).Info("create pod %v", pod)
	return nil
}

func (n VirtualNodeProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	log.G(ctx).Info("UpdatePod")
	return nil
}

func (n VirtualNodeProvider) DeletePod(ctx context.Context, pod *v1.Pod) error {
	log.G(ctx).Infof("DeletePod pod namespace %s,name %s", pod.GetNamespace(), pod.GetName())
	return nil
}

func (n VirtualNodeProvider) GetPod(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	log.G(ctx).Info("GetPod")
	return nil, nil
}

func (n VirtualNodeProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	pod, err := n.GetPod(ctx, namespace, name)
	if err != nil {
		return nil, err
	}
	return &pod.Status, nil
}

func (n VirtualNodeProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	var pods []*v1.Pod
	podList, err := n.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range podList.Items {
		_pod := pod
		pods = append(pods, &_pod)
	}
	return pods, nil
}

func (n VirtualNodeProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	log.G(ctx).Info("GetContainerLogs")
	return nil, nil
}

func (n VirtualNodeProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach api.AttachIO) error {
	log.G(ctx).Info("RunInContainer")
	return nil
}

func (n VirtualNodeProvider) ConfigureNode(ctx context.Context, v *v1.Node) {
	v.Status.Conditions = []v1.NodeCondition{
		{
			Type:               "Ready",
			Status:             v1.ConditionTrue,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletReady",
			Message:            "kubelet is ready.",
		},
		{
			Type:               "OutOfDisk",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientDisk",
			Message:            "kubelet has sufficient disk space available",
		},
		{
			Type:               "MemoryPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasSufficientMemory",
			Message:            "kubelet has sufficient memory available",
		},
		{
			Type:               "DiskPressure",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "KubeletHasNoDiskPressure",
			Message:            "kubelet has no disk pressure",
		},
		{
			Type:               "NetworkUnavailable",
			Status:             v1.ConditionFalse,
			LastHeartbeatTime:  metav1.Now(),
			LastTransitionTime: metav1.Now(),
			Reason:             "RouteCreated",
			Message:            "RouteController created a route",
		},
	}
}

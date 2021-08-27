package vnode

import (
	"context"
	"github.com/laik/vnode/pkg/log"
	v1 "k8s.io/api/core/v1"
)

func (n VirtualNodeProvider) NotifyPods(ctx context.Context, f func(*v1.Pod)) {
	log.G(ctx).Info("NotifyPods")

}

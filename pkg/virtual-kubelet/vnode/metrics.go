package vnode

import (
	"context"
	"github.com/virtual-kubelet/virtual-kubelet/node/api/statsv1alpha1"
)

func (n VirtualNodeProvider) GetStatsSummary(ctx context.Context) (*statsv1alpha1.Summary, error) {
	panic("implement me")
}

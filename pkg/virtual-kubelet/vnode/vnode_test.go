package vnode

import (
	virtual_kubelet "github.com/laik/vnode/pkg/virtual-kubelet"
	"github.com/virtual-kubelet/virtual-kubelet/node"
	"gotest.tools/assert"
	"testing"
)

func TestNodeLegacyInterface(t *testing.T) {
	var mlp virtual_kubelet.Provider = &VirtualNodeProvider{}
	_, ok := mlp.(node.PodNotifier)
	assert.Assert(t, ok, true)
}

module github.com/laik/vnode

go 1.16

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/virtual-kubelet/virtual-kubelet v0.0.0-00010101000000-000000000000
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.19.10
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.9.0
)

replace (
	github.com/virtual-kubelet/virtual-kubelet => github.com/liqotech/virtual-kubelet v1.5.1-0.20210726130647-f2333d82a6de
	k8s.io/client-go => k8s.io/client-go v0.22.1
)

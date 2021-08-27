package main

import (
	"context"
	"github.com/laik/vnode/cmd/providers"
	root2 "github.com/laik/vnode/cmd/root"
	"github.com/laik/vnode/cmd/version"
	"github.com/laik/vnode/pkg/log"
	logruslogger "github.com/laik/vnode/pkg/log/logrus"
	provider "github.com/laik/vnode/pkg/virtual-kubelet"
	"github.com/laik/vnode/pkg/virtual-kubelet/constraint2"
	"github.com/laik/vnode/pkg/virtual-kubelet/vnode"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	buildVersion = "N/A"
	buildTime    = "N/A"
	k8sVersion   = "v1.22.0" // This should follow the version of k8s.io/kubernetes we are importing
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	log.L = logruslogger.FromLogrus(logrus.NewEntry(logrus.StandardLogger()))
	var opts root2.Opts
	optsErr := root2.SetDefaultOpts(&opts)
	opts.Version = strings.Join([]string{k8sVersion, "vk", buildVersion}, "-")

	s := provider.NewStore()
	// register provider
	regErr := register(ctx, s)

	rootCmd := root2.NewCommand(ctx, filepath.Base(os.Args[0]), s, opts)
	rootCmd.AddCommand(version.NewCommand(buildVersion, buildTime), providers.NewCommand(s))
	preRun := rootCmd.PreRunE

	var logLevel string
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if optsErr != nil {
			return optsErr
		}
		if regErr != nil {
			return regErr
		}
		if preRun != nil {
			return preRun(cmd, args)
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", `set the log level, e.g. "debug", "info", "warn", "error"`)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if logLevel != "" {
			lvl, err := logrus.ParseLevel(logLevel)
			if err != nil {
				return errors.Wrap(err, "could not parse log level")
			}
			logrus.SetLevel(lvl)
		}
		return nil
	}

	log.G(ctx).Infof("start vnode")
	if err := rootCmd.Execute(); err != nil && errors.Cause(err) != context.Canceled {
		log.G(ctx).Fatal(err)
	}
}

func register(ctx context.Context, s *provider.Store) error {
	return s.Register(
		constraint2.ProviderName,
		func(cfg provider.InitConfig) (provider.Provider, error) { //nolint:errcheck
			return vnode.NewVirtualNodeProvider(ctx, &cfg)
		},
	)
}

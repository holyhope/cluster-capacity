package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/holyhope/cluster-capacity/pkg/logger"
	"github.com/holyhope/cluster-capacity/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

const (
	MemoryCmd = "memory"
	CPUCmd    = "cpu"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "cluster-capacity",
		Short:         "",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewLogger()
			ctx := logger.ToContext(context.TODO(), log)

			if len(args) != 1 {
				return errors.New("expected 1 parameter")
			}

			switch args[0] {
			case MemoryCmd:
				if err := plugin.TotalMemory(ctx, KubernetesConfigFlags); err != nil {
					return errors.Unwrap(err)
				}
			case CPUCmd:
				if err := plugin.TotalCPU(ctx, KubernetesConfigFlags); err != nil {
					return errors.Unwrap(err)
				}
			default:
				return fmt.Errorf("unsupported argument %s, expected %s or %s", args[0], MemoryCmd, CPUCmd)
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}

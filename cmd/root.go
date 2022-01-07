package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	envPrefix = "RETROMEET"

	rootCmd = &cobra.Command{
		Use:   "retromeet",
		Short: "Team retrospective application",
		Long:  "Application server for Web clients",
	}
)

func init() {
	cobra.OnInitialize(initConfiguration)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfiguration() {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
}

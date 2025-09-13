package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "vmopt",
	Short: "VMOpt CLI - Monitor & Optimize VMware recource usage",
	Long:  "VMOpt is a CLI tool written in Go for monitoring and opimizing CPU/RAM usage of VMware wokrstation virtual machines.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	fmt.Println("VMOpt CLI - VMware Optimizer")
}

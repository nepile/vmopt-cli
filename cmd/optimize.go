package cmd

import (
	"fmt"
	"vmopt-cli/logger"
	"vmopt-cli/optimizer"

	"github.com/spf13/cobra"
)

var maxCPU float64
var maxRAM float64
var action string

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize VMware VM usage based on thresholds",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Optimizing VMs...")

		err := optimizer.OptimizeVMs(maxCPU, maxRAM, action)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			logger.LogEvent(fmt.Sprintf("Optimization applied (CPU: %.2f, RAM: %.2f, Action: %s)", maxCPU, maxRAM, action))
		}
	},
}

func init() {
	rootCmd.AddCommand(optimizeCmd)
	optimizeCmd.Flags().Float64Var(&maxCPU, "max-cpu", 80, "Maximum limit of CPU usage")
	optimizeCmd.Flags().Float64Var(&maxRAM, "max-ram", 90, "Maximum limit of RAM usage")
	optimizeCmd.Flags().StringVar(&action, "action", "lower-priority", "Action (lower-priority / kill)")
}

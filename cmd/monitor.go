package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(monitorCmd)
}

var threshold float64

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitors for activity and keeps track",
	Long:  ,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	}
}

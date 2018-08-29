package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.1",
	Long:  `goroot 0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goroot v0.1 ")
	},
}

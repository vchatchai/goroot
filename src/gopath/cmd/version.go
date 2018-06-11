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
	Short: "1.0",
	Long:  `GoPath 1.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoPath v1.0 ")
	},
}

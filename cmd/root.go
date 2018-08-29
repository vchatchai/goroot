package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goroot",
	Short: "goroot is utility for change $GOROOT",
	Long:  `goroot is utility for change $GOROOT`,
	Run: func(cmd *cobra.Command, args []string) {
		ChangePath()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopath",
	Short: "gopath is utility for change $GOPATH",
	Long:  `gopath is utility for change $GOPATH`,
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

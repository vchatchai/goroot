package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vchatchai/gopath/config"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add GOPATH",
	Long:  `add GOPATH`,
	Run: func(cmd *cobra.Command, args []string) {
		AddPath(cmd, args)
	},
}

func AddPath(cmd *cobra.Command, args []string) {

	size := len(args)
	if size == 0 {
		fmt.Println("NO Path!")
		return
	}

	for _, path := range args {
		p, _ := filepath.Abs(path)
		fmt.Println(p)
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
			return
		}
		_, lastpath := filepath.Split(path)
		config.AddNewPath(strings.Trim(path, " "), lastpath)

		fmt.Printf("Add Path %q\n", path)

	}

}

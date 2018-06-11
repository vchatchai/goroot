package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopath/config"

	"github.com/spf13/cobra"
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

	// AddNewPath("/usr/")
	// AddNewPath("/home/")
	// AddNewPath("/var/")

	// config.AddNewPath("")
	fmt.Println(args)

	for _, path := range args {
		path, err := filepath.Abs(filepath.Dir(path))
		fmt.Println(path)
		if err != nil {
			panic(err)
			return
		}
		_, lastpath := filepath.Split(path)
		config.AddNewPath(strings.Trim(path, " "), lastpath)

		fmt.Printf("Path %q\n", path)

	}

}

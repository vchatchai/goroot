package cmd

import (
	"fmt"
	"strings"

	"github.com/vchatchai/gopath/config"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changeCmd)
}

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change $GOPATH",
	Long:  `change $GOPATH`,
	Run: func(cmd *cobra.Command, args []string) {
		ChangePath()
	},
}

func ChangePath() {

	// AddNewPath("/usr/")
	// AddNewPath("/home/")
	// AddNewPath("/var/")

	path, _ := config.GetPath()
	path = append(path, config.Path{Key: config.QUIT})
	// keys := reflect.ValueOf(mapPath).MapKeys()

	// prompt := promptui.Select{
	// 	Label: "Select Path",
	// 	Items: path,
	// }
	prompt := promptui.Select{
		Label: "LIST GOPATH",
		Items: path,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	results := strings.Split(result, ":")

	if results[0] == config.QUIT {
		return
	}

	fmt.Printf("Change to %q\n", strings.Trim(results[1], " "))

}

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

	path, _ := config.GetPath()
	path = append(path, config.Path{Key: config.QUIT})
	value, _ := config.PATH.GetPath()

	prompt := promptui.Select{
		Label: fmt.Sprintf("Current GOPATH: %s", value),
		Items: path,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	results := strings.Split(result, "::")

	if results[0] == config.QUIT {
		return
	}

	gopath := strings.Trim(results[1], " ")
	fmt.Printf("Current GOPATH %q\n", gopath)

	err = config.PATH.ChangePath(gopath)
	if err != nil {
		panic(err)
	}
}

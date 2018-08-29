package cmd

import (
	"fmt"
	"strings"

	"github.com/vchatchai/goroot/config"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changeCmd)
}

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change $GOROOT",
	Long:  `change $GOROOT`,
	Run: func(cmd *cobra.Command, args []string) {
		ChangePath()
	},
}

func ChangePath() {

	path, _ := config.GetPath()
	path = append(path, config.Path{Key: config.QUIT})
	value, _ := config.PATH.GetPath()

	prompt := promptui.Select{
		Label: fmt.Sprintf("Current GOROOT: %s", value),
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

	goroot := strings.Trim(results[1], " ")
	fmt.Printf("Current GOROOT %q\n", goroot)

	err = config.PATH.ChangePath(goroot)
	if err != nil {
		panic(err)
	}
}

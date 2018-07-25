package cmd

import (
	"fmt"
	"strings"

	"github.com/vchatchai/gopath/config"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete path",
	Long:  `delete gopath`,
	Run: func(cmd *cobra.Command, args []string) {
		DeletePath(cmd, args)
	},
}

func DeletePath(cmd *cobra.Command, args []string) {

	path, _ := config.GetPath()
	path = append(path, config.Path{Key: config.QUIT})

	prompt := promptui.Select{
		Label: "LIST GOPATH",
		Items: path,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == config.QUIT {
		return
	}
	results := strings.Split(result, "::")

	result = strings.Trim(results[1], " ")

	err = config.RemovePath(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Remove GOPath %q\n", result)

}

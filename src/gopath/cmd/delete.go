package cmd

import (
	"fmt"
	"strings"

	"gopath/config"

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

	// AddNewPath("/usr/")
	// AddNewPath("/home/")
	// AddNewPath("/var/")

	path, _ := config.GetPath()

	// keys := reflect.ValueOf(mapPath).MapKeys()

	prompt := promptui.Select{
		Label: "Select Project Path",
		Items: path,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Println(result)
	results := strings.Split(result, ":")
	result = strings.Trim(results[1], " ")
	fmt.Println(result)
	err = config.RemovePath(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("RemovePath %q\n", result)

}

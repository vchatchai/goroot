package main

import (
	"fmt"
	"reflect"

	promptui "github.com/manifoldco/promptui"
)

func main() {

	fmt.Println("Test")

	// AddNewPath("/usr/")
	// AddNewPath("/home/")
	// AddNewPath("/var/")

	mapPath, _ := GetConfig()

	keys := reflect.ValueOf(mapPath).MapKeys()

	prompt := promptui.Select{
		Label: "Select Project Path",
		Items: keys,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// RemovePath(result)

	fmt.Printf("You choose %q\n", result)

}

package main

import (
	"fmt"
	"longstraws/go-todo-app/cmd/todo-app/root"
)

func main() {

	rootCmd := root.RootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

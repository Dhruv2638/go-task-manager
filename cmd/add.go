package cmd

import (
	"fmt"
	"os"
	"strings"
	"task-manager/db"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to a task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}

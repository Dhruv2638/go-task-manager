package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task-manager/db"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error : %s\n", task.Key, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", task.Key)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(DoCmd)
}

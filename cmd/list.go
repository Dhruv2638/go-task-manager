package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("first")
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}

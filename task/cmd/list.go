package cmd

import (
	"fmt"
	"log"

	"github.com/Gabriel2233/gophercises/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the tasks inside the list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()

		if err != nil {
			log.Fatal("Error while reading tasks...\n")
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks to complete!")
			return
		}

		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

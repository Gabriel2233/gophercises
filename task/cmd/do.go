package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Gabriel2233/gophercises/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)

			if err != nil {
				fmt.Printf("Failed to parse argument %s\n", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ReadTasks()

		if err != nil {
			log.Fatal("An error ocurred while reading all tasks\n")
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("The provided task doesn't exist")
				continue
			}

			task := tasks[id-1]
			err := db.DoTask(task.Key)

			if err != nil {
				fmt.Printf("Could not do task with id %d\n", id)
			} else {
				fmt.Printf("Task with id %d was marked as done\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

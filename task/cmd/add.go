package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Gabriel2233/gophercises/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task inside your list",
	Run: func(cmd *cobra.Command, args []string) {
		t := strings.Join(args, " ")

		id, err := db.CreateTask(t)

		if err != nil {
			log.Fatal("Error while creating a new task\n")
		}

		fmt.Printf("New Task created, id: %d\n", id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

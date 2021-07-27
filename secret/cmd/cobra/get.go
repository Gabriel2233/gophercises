package cobra

import (
	"fmt"

	secret "github.com/Gabriel2233/gophercises/secret/mem_vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your vault",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(EncodingKey, secretsPath())
		key := args[0]
		val, err := v.Get(key)
		if err != nil {
			fmt.Println("no value set")
			return
		}

		fmt.Printf("%s=%s\n", key, val)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

package cobra

import (
	"fmt"

	secret "github.com/Gabriel2233/gophercises/secret/mem_vault"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your vault",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(EncodingKey, secretsPath())
		key, val := args[0], args[1]
		err := v.Set(key, val)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Set %s in vault\n", key)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

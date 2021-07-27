package cobra

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var EncodingKey string

var rootCmd = &cobra.Command{
	Use:   "secret",
	Short: "A simple CLI to store and manage API keys and other secrets",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&EncodingKey, "key", "k", "", "the encoding key used to encrypt and decrypt secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secret")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "WC",
	Short: "Word, line, character, and byte counter",
	Long:  `A Fast Word counter with some parameters to count lines, words, characters, and bytes.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("test")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

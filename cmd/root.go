package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const version = "1.0.0"

var outputPath string

var rootCmd = &cobra.Command{
	Use:                   "file-crypto COMMAND FILE-PATH [-o --output <file-path>]",
	Short:                 "File-crypto is a tool for encryption and decryption text files.",
	Version:               version,
	DisableFlagsInUseLine: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDescriptions: true,
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
	},
	Run: func(cmd *cobra.Command, args []string) {

		// If just root command run, show help template
		if err := cmd.Help(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {

	// Run cli app
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

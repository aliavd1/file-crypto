package cmd

import (
	"avd.com/file-crypto/cryptoutils"
	"avd.com/file-crypto/fileutils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var encryptCmd = &cobra.Command{
	Use:                   "encrypt FILE-PATH [-o --output <file-path>]",
	Short:                 "Encrypt is a sub-command for encryption text files.",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		// If just encrypt command run, show help template
		if len(args) < 1 {
			if err := cmd.Help(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			plainTextByte := fileutils.ReadFromFile(args[0])
			cipherTextByte := cryptoutils.Encrypt(plainTextByte)
			if outputPath == "" {
				outputPath = args[0] + ".bin"
			}
			fileutils.WriteToFile(cipherTextByte, outputPath)
		}
	},
}

func init() {
	encryptCmd.Flags().StringVarP(
		&outputPath,
		"output",
		"o",
		"",
		"Output path for encrypted file",
	)
	rootCmd.AddCommand(encryptCmd)
}

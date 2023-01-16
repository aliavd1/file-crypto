package cmd

import (
	"avd.com/file-crypto/cryptoutils"
	"avd.com/file-crypto/fileutils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var decryptCmd = &cobra.Command{
	Use:                   "decrypt FILE-PATH [-o --output <file-path>]",
	Short:                 "Decrypt is a sub-command for decryption text files.",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		// If just decrypt command run, show help template
		if len(args) < 1 {
			if err := cmd.Help(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		} else {
			cipherTextByte := fileutils.ReadFromFile(args[0])
			plainTextByte := cryptoutils.Decrypt(cipherTextByte)
			if outputPath == "" {
				outputPath = args[0] + ".txt"
			}
			fileutils.WriteToFile(plainTextByte, outputPath)
		}
	},
}

func init() {
	decryptCmd.Flags().StringVarP(
		&outputPath,
		"output",
		"o",
		"",
		"Output path for decrypted file",
	)
	rootCmd.AddCommand(decryptCmd)
}

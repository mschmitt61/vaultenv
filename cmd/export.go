/*
Copyright © 2023 Mitchell Schmitt mschmitt61@massmutual.com
*/
package cmd

import (
	"fmt"
	helpers "github.com/mschmitt61/vaultenv/helpers"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// applyCmd represents the apply command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports all the variables in the input file to the local environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 || args[0] == "" {
			fmt.Println("`vaultenv generate` command requires exactly two arguments - an input file and an outputfile")
		}
		app := helpers.InitWrapper()

		pwd, err := os.Getwd()
		if err != nil {
			app.Logger.Fatalf("Failed to get working directory, exiting")
		}

		inputFile := args[0]
		pathToInputFile := filepath.Join(pwd, inputFile)

		app.ExportEnvFile(pathToInputFile)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

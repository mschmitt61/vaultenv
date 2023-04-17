/*
Copyright © 2023 Mitchell Schmitt mschmitt61@massmutual.com
*/
package cmd

import (
	"fmt"
	helpers "github.com/massmutual/vaultenv/helpers"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Exports all the variables in the input file to the local environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 2 || args[0] != "" {
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
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

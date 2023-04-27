/*
Copyright © 2023 Mitchell Schmitt mschmitt61@massmutual.com
*/
package cmd

/*
import (
	"fmt"
	helpers "github.com/mschmitt61/vaultenv/helpers"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"errors"
)

// applyCmd represents the apply command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Exports all the variables in the input file to the local environment",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 || args[0] == "" {
			fmt.Println("`vaultenv export` command requires exactly one argument - the file with the env variables you want to export")
		}
		app := helpers.InitWrapper()

		pwd, err := os.Getwd()
		if err != nil {
			app.Logger.Fatalf("Failed to get working directory, exiting")
		}

		inputFile := args[0]
		pathToInputFile := filepath.Join(pwd, inputFile)

		app.ExportEnvFile(pathToInputFile)
		return errors.New("")
	},
}

func init() {
	RootCmd.AddCommand(exportCmd)
}
*/

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

// GenerateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Given an input template file and output file, take the input file and generate the env variables based on the vault values",
	RunE: func(cmd *cobra.Command, args []string) error {
		app, err := helpers.InitApp()
		if err != nil {
			return fmt.Errorf("failed to initialize vaultenv: %v", err)
		}

		if len(args) != 2 {
			return fmt.Errorf("`vaultenv generate` command requires exactly two arguments - an input file and an outputfile")
		}

		pwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %v", err)
		}

		inputFile := args[0]
		pathToInputFile := filepath.Join(pwd, inputFile)

		outputFile := args[1]
		pathToOutputFile := filepath.Join(pwd, outputFile)

		kvs, err := app.ReadEnvFile(pathToInputFile)
		if err != nil {
			return fmt.Errorf("failed to read env file %s: %v", pathToInputFile, err) 
		}

		file, err := os.Create(pathToOutputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file %s: %v", pathToOutputFile, err) 
		}
		defer file.Close()

		for k, v := range kvs {
			stringToWrite := fmt.Sprintf("%s=%s\n", k, v)
			bytesWritten, err := file.WriteString(stringToWrite)
			if err != nil {
				return fmt.Errorf("failed to write to file %s: %v", pathToOutputFile, err)
			}
			app.Logger.Printf("Wrote %d bytes to file %s", bytesWritten, pathToOutputFile)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}

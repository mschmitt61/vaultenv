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

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Given an input template file and output file, take the input file and generate the env variables based on the vault values",
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

		outputFile := args[1]
		pathToOutputFile := filepath.Join(pwd, outputFile)

		kvs, err := app.ReadEnvFile(pathToInputFile)
		if err != nil {
			app.Logger.Fatalf("Failed to read env file %s exiting", pathToInputFile)
		}

		file, err := os.Create(pathToOutputFile)
		if err != nil {
			app.Logger.Fatalf("Failed to create output file %s exiting", pathToOutputFile)
		}
		defer file.Close()

		for k, v := range kvs {
			stringToWrite := fmt.Sprintf("%s=%s\n", k, v)
			bytesWritten, err := file.WriteString(stringToWrite)
			if err != nil {
				app.Logger.Fatalf("Failed to write to file %s exiting", pathToOutputFile)
			}
			app.Logger.Printf("Wrote %d bytes to file %s", bytesWritten, pathToOutputFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
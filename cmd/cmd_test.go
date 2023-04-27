/*
Copyright © 2023 Mitchell Schmitt mschmitt61@massmutual.com
*/
package cmd

import (
	"bytes"
	"fmt"
	"testing"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
	}
	return nil
}

func ExactArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

func TestOneArgError(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:  "root",
		Args: ExactArgs(2),
		Run:  func(_ *cobra.Command, args []string) {},
	}
	aCmd := GenerateCmd
	rootCmd.AddCommand(aCmd)

	expected := "`vaultenv generate` command requires exactly two arguments - an input file and an outputfile"

	_, err := executeCommand(rootCmd, "generate", "one")


	assert.Equal(t, expected, err.Error())
}

func TestMissingFileError(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:  "root",
		Args: ExactArgs(2),
		Run:  func(_ *cobra.Command, args []string) {},
	}
	aCmd := GenerateCmd
	rootCmd.AddCommand(aCmd)

	pwd, _ := os.Getwd()
	pathToInputFile := filepath.Join(pwd, "nofile")

	expected := fmt.Sprintf("failed to read env file %s: failed to read file %s: open %s: no such file or directory", pathToInputFile, pathToInputFile, pathToInputFile)

	_, err := executeCommand(rootCmd, "generate", "nofile", "nofile2")


	assert.Equal(t, expected, err.Error())
}

func TestWorkingPath(t *testing.T) {
	rootCmd := &cobra.Command{
		Use:  "root",
		Args: ExactArgs(2),
		Run:  func(_ *cobra.Command, args []string) {},
	}
	aCmd := GenerateCmd
	rootCmd.AddCommand(aCmd)

	pwd, _ := os.Getwd()
	pathToInputFile := filepath.Join(pwd, ".env.template.test")
	pathToOutputFile := filepath.Join(pwd, ".env.test")
	executeCommand(rootCmd, "generate", ".env.template.test", ".env.test")

	inputFileContents, _ := os.ReadFile(pathToInputFile)
	outputFileContents, _ := os.ReadFile(pathToOutputFile)

	assert.Equal(t, inputFileContents, outputFileContents)

	os.Remove(pathToOutputFile)
}

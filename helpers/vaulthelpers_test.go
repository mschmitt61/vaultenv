package cmd

import (
	"os"
	"strings"
	"testing"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)

func TestAppEnvValuesSet(t *testing.T) {
	_, err := InitApp()

	assert.Equal(t, err, nil)
}

func TestAppVaultAddrNotSet(t *testing.T) {
	vault_addr := os.Getenv("VAULT_ADDR")
	os.Unsetenv("VAULT_ADDR")
	_, err := InitApp()

	expectedErr := "failed to find environment variable VAULT_ADDR, please set and rerun"
	os.Setenv("VAULT_ADDR", vault_addr)

	assert.Equal(t, err.Error(), expectedErr)
}

func TestAppVaultRoleNotSet(t *testing.T) {
	vault_role := os.Getenv("VAULT_ROLE")
	os.Unsetenv("VAULT_ROLE")
	_, err := InitApp()

	expectedErr := "failed to find environment variable VAULT_ROLE, please set and rerun"
	os.Setenv("VAULT_ROLE", vault_role)

	assert.Equal(t, err.Error(), expectedErr)
}

func TestAppVaultSecretNotSet(t *testing.T) {
	vault_secret := os.Getenv("VAULT_SECRET")
	os.Unsetenv("VAULT_SECRET")
	_, err := InitApp()

	expectedErr := "failed to find environment variable VAULT_SECRET, please set and rerun"
	os.Setenv("VAULT_SECRET", vault_secret)

	assert.Equal(t, err.Error(), expectedErr)
}

func TestInitVaultErr(t *testing.T) {
	vault_secret := os.Getenv("VAULT_SECRET")

	os.Setenv("VAULT_SECRET", "notarealsecret")
	_, err := InitApp()

	expectedErr := "failed to initialize vault client: failed to read from vault: Error making API request."
	os.Setenv("VAULT_SECRET", vault_secret)

	res := strings.Contains(err.Error(), expectedErr)

	assert.Equal(t, true, res)
}

func TestReadEnvFile(t *testing.T) {
	app, _ := InitApp()
	pwd, _ := os.Getwd()
	pathToInputFile := filepath.Join(pwd, ".env.template.test")
	output, _ := app.ReadEnvFile(pathToInputFile)


	expectedMap := map[string]string{"key": "value"}


	assert.Equal(t, output, expectedMap)
}

func TestReadEnvFileWithVault(t *testing.T) {
	app, _ := InitApp()
	pwd, _ := os.Getwd()
	pathToInputFile := filepath.Join(pwd, ".env.template.vault")
	output, _ := app.ReadEnvFile(pathToInputFile)

	expectedMap := map[string]string{"vaultsecret": "test"}

	assert.Equal(t, output, expectedMap)
}

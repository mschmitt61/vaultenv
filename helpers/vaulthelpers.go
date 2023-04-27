/*
Copyright © 2023 Mitchell Schmitt mschmitt61@massmutual.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"os"
	"strings"
)

type VaultWrapper struct {
	VaultAddr   string
	VaultRole   string
	VaultSecret string
	Client      *api.Client
	Logger      *log.Logger
}

func InitApp() (*VaultWrapper, error) {
	prefix := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 94, "[vaultenv] ")
	logger := log.New(os.Stdout, prefix, log.LstdFlags)

	vaultAddr, found := os.LookupEnv("VAULT_ADDR")
	if !found {
		return nil, fmt.Errorf("failed to find environment variable VAULT_ADDR, please set and rerun")
	}

	vaultRole, found := os.LookupEnv("VAULT_ROLE")
	if !found {
		return nil, fmt.Errorf("failed to find environment variable VAULT_ROLE, please set and rerun")
	}

	vaultSecret, found := os.LookupEnv("VAULT_SECRET")
	if !found {
		return nil, fmt.Errorf("failed to find environment variable VAULT_SECRET, please set and rerun")
	}

	vault, err := initVault(vaultAddr, vaultRole, vaultSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize vault client: %v", err)
	}

	vaultWrapper := VaultWrapper{
		VaultAddr:   vaultAddr,
		VaultRole:   vaultRole,
		VaultSecret: vaultSecret,
		Client:      vault,
		Logger:      logger,
	}

	return &vaultWrapper, nil
}

func initVault(vaultAddr string, vaultRole string, vaultSecret string) (*api.Client, error) {

	config := &api.Config{
		Address: vaultAddr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return client, fmt.Errorf("failed to initialize client err: %v", err)
	}

	vault_secret := vaultSecret
	vault_role := vaultRole
	params := map[string]interface{}{
		"role_id":   vault_role,
		"secret_id": vault_secret,
	}
	resp, err := client.Logical().Write("auth/approle/login", params)
	if err != nil {
		return client, fmt.Errorf("failed to read from vault: %v", err)
	}

	client.SetToken(resp.Auth.ClientToken)

	return client, nil
}

// Read a value from vault based on the path and value arguments. secret/ is prepended for ease of use
func (vaultWrapper *VaultWrapper) ReadFromVault(path string, value string) (string, error) {
	secret, err := vaultWrapper.Client.Logical().Read(fmt.Sprintf("secret/%s", path))
	if err != nil {
		return "", fmt.Errorf("failed to read from vault: %v", err)
	}

	return fmt.Sprintf("%v", secret.Data[value]), nil
}

// Read the .env file based on the path and return the vault parsed key/value map
func (vaultWrapper *VaultWrapper) ReadEnvFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", path, err)
	}
	defer file.Close()

	retMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Found a line to read from vault
		if strings.Contains(line, "vault::") {
			vaultWrapper.Logger.Printf("Found env variable with vault parsing %s", line)
			split := strings.Split(line, "=")
			envName := split[0]
			vaultSplit := strings.Split(split[1], "::")
			vaultPath := vaultSplit[1]
			vaultKey := vaultSplit[2]
			vaultValue, err := vaultWrapper.ReadFromVault(vaultPath, vaultKey)
			if err != nil {
				return nil, fmt.Errorf("failed to read from vault %s: %v", path, err)
			}
			_, exists := retMap[envName]
			if exists {
				vaultWrapper.Logger.Printf("Found duplicate value in env file %s: %s, skipping", path, envName)
				continue
			}
			retMap[envName] = vaultValue
		}
		// Found a normal environment variable, we still want to export it
		if strings.Contains(line, "=") && !strings.Contains(line, "vault::") {
			vaultWrapper.Logger.Printf("Found env variable %s", line)
			split := strings.Split(line, "=")
			_, exists := retMap[split[0]]
			if exists {
				vaultWrapper.Logger.Printf("Found duplicate value in env file %s: %s, skipping", path, split[0])
				continue
			}
			retMap[split[0]] = split[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed when reading %s: %v", path, err)
	}
	return retMap, nil
}


// Both below functions won't work for exporting to the local terminal environment because they are in a child process
// Export all the variables in the map to the local environment
// func (vaultWrapper *VaultWrapper) ExportEnvs(envMap map[string]string) {
// 	for key, value := range envMap {
// 		os.Setenv(key, value)
// 	}
// }

// Export all the variables in the file to the local environment
// func (vaultWrapper *VaultWrapper) ExportEnvFile(path string) error {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return fmt.Errorf("failed to read file %s: %v", path, err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		split := strings.Split(line, "=")
// 		key := split[0]
// 		value := split[1]
// 		os.Setenv(key, value)
// 	}
// 	return nil
// }

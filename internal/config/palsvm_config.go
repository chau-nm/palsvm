package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	FileConfigLocation     = "~/.palsvm/palsvm_config.json"
	DefaultServerUsername  = "admin"
	DefaultServerPassword  = "admin"
	DefaultServerHostName  = "0.0.0.0"
	DefaultServerPort      = 8080
	FolderConfigPermission = 0700
	FileConfigPermission   = 0600
)

type PalsvmConfig struct {
	ServerUsername string `json:"server_username"`
	ServerPassword string `json:"server_password"`
	ServerHostname string `json:"server_hostname"`
	ServerPort     int    `json:"server_port"`
	ServerSecret   string `json:"server_secret"`
}

// LoadPalsvmConfig read config file and convert to PalsvmConfig
func LoadPalsvmConfig() (*PalsvmConfig, error) {
	configPath := expandPath(FileConfigLocation)

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			config := &PalsvmConfig{
				ServerUsername: DefaultServerUsername,
				ServerPassword: DefaultServerPassword,
				ServerHostname: DefaultServerHostName,
				ServerPort:     DefaultServerPort,
			}
			if err := SavePalsvmConfig(config); err != nil {
				return nil, err
			}
			return config, nil
		}
		return nil, err
	}

	var config PalsvmConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SavePalsvmConfig write PalsvmConfig into config file
func SavePalsvmConfig(config *PalsvmConfig) error {
	configPath := expandPath(FileConfigLocation)

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, FolderConfigPermission); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, FileConfigPermission)
}

// expandPath convert ~/ to root directory
// homeDir = "/home/user" (Linux)
// homeDir = "/Users/user" (Mac)
// homeDir = "C:\Users\user" (Windows)
func expandPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(homeDir, path[1:])
	}
	return path
}

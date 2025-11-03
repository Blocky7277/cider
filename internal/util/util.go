package util

import (
	"os"
	"fmt"
	"path/filepath"
	"gopkg.in/yaml.v3"
)

type ConfigData struct {
	User          string `yaml:"user"`
	InstallDir    string `yaml:"install_dir"`
	CopyManifest  bool   `yaml:"copy_manifest"`
	SteamappsDir  string `yaml:"steamapps_dir"`
}

func DefaultConfig() ConfigData {
	return ConfigData{
		User:         "",
		InstallDir:   "",
		CopyManifest: false,
		SteamappsDir: "",
	}
}

func Exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Can't find OS config directory?")
		fmt.Println("File a bug report")
		return "", err
	}
	appDir := filepath.Join(configDir, "cider")
	if exists, _ := Exists(appDir); !exists {
	    err := os.Mkdir(appDir, 0755)
		if err != nil {
			fmt.Println("Failed to create app directory exiting")
			return "", err
		}
	}
	return filepath.Join(appDir, "config.yaml"), nil
}

func LoadConfig(path string) (ConfigData, error) {
	data := DefaultConfig()
	bytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := SaveConfig(path, data); err != nil {
				fmt.Fprintln(os.Stderr, "Failed to save config:", err)
				os.Exit(4)
			}
			return data, nil
		}
		return data, err
	}
	err = yaml.Unmarshal(bytes, &data)
	return data, err
}

func SaveConfig(path string, conf ConfigData) error {
	bytes, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}

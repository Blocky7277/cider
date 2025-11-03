package cmds

import (
	"os"
	"fmt"
	"cider/internal/util"
)

func Config(args []string) {
	confPath, err := util.GetConfigPath()
	if err != nil {
		os.Exit(1)
	}
	config, err := util.LoadConfig(confPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	if len(args) < 1 {
		fmt.Println("username: " + config.User)
		fmt.Println("install_dir: " + config.InstallDir)
		fmt.Print("manifest: ")
		fmt.Println(config.CopyManifest)
		fmt.Println("steamapps_dir: " + config.SteamappsDir)
	} else if len(args) < 2 {
		//TODO: Print the desired confkey := args[1]
		key := args[0]
		switch key {
		case "username":
			fmt.Println(config.User)
		case "install_dir":
			fmt.Println(config.InstallDir)
		case "manifest":
			fmt.Println(config.CopyManifest)
		case "steamapps_dir":
			fmt.Println(config.SteamappsDir)
		default:
			fmt.Println("Unknown option:", key)
		}
	} else {
		key, value := args[0], args[1]
		// TODO: FIND CASE AND SET
		switch key {
		case "username":
			config.User = value
		case "install_dir":
			config.InstallDir = value
		case "manifest":
			config.CopyManifest = (value == "true")
		case "steamapps_dir":
			config.SteamappsDir = value
		default:
			fmt.Println("Unknown option: ", key)
			return
		}
		if err := util.SaveConfig(confPath, config); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to save config:", err)
			os.Exit(4)
		}
		fmt.Println("Updated", key)
	}
}

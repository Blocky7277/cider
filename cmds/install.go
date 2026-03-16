package cmds

import (
	"os/exec"
	"fmt"
	"os"
	"io"
	"log"
	"bytes"
	"strings"
	"cider/internal/util"
	"path/filepath"
)

func Install() {
	path, err := util.GetConfigPath()
	config, err := util.LoadConfig(path)
	_, err = exec.LookPath("steamcmd")
	if err != nil {
		log.Fatal("SteamCMD missing: ", err)
	}
	var gameId int
	fmt.Print("Enter Steam AppID:\n")
	fmt.Scanln(&gameId)
	var gameName string
	fmt.Print("Enter Steam game name:\n")
	fmt.Scanln(&gameName)
	gameName = strings.ReplaceAll(gameName, "/", "_")
	installDir := filepath.Join(config.InstallDir, gameName) 
	cmd := exec.Command("steamcmd",
		"+@sSteamCmdForcePlatformType", "windows",
		"+force_install_dir", installDir,
		"+login", config.User, "",
		"+app_update", fmt.Sprint(gameId),
		"+quit")

	fmt.Println("Working...")
	cmd.Stdout = os.Stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %v\nStderr: %s", err, stderr.String())
	}
	fmt.Println("Finished...")
	if config.CopyManifest {
		manifest := filepath.Join(installDir, "steamapps", fmt.Sprintf("appmanifest_%d.acf", gameId))	
		target := filepath.Join(config.SteamappsDir, fmt.Sprintf("appmanifest_%d.acf", gameId))
		source, err := os.Open(manifest)
	if err != nil {
		
	}
	defer source.Close()

	dest, err := os.Create(target)
	if err != nil {

	}
	defer dest.Close()

	_, err = io.Copy(dest, source)

	}
}

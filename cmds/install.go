package cmds

import (
	"os/exec"
	"golang.org/x/term"
	"syscall"
	"fmt"
	"os"
	"io"
	"log"
	"cider/internal/util"
	"path/filepath"
)

func Install() {
	path, err := util.GetConfigPath()
	config, err := util.LoadConfig(path)
	_, err = exec.LookPath("steamcmd")
	if err != nil {
		fmt.Println("FATAL: steamcmd not found, exiting")
		os.Exit(0)
	}
	fmt.Print("Enter Steam password:\n")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	pwd := string(bytePassword)
	var gameId int
	fmt.Print("Enter Steam AppID:\n")
	fmt.Scanln(&gameId)
	var gameName string
	fmt.Print("Enter Steam game name:\n")
	fmt.Scanln(&gameName)
	installDir := filepath.Join(config.InstallDir, gameName) 
	cmd := exec.Command("steamcmd",
		"+@sSteamCmdForcePlatformType", "windows",
		"+force_install_dir", installDir,
		"+login", config.User, pwd,
		"+app_update", fmt.Sprint(gameId),
		"+quit")

	fmt.Println("Working...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %v\nStderr: %s", err)
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

package cmds

import (
	"fmt"
)

func Help() {
	fmt.Println("- cider -")
	fmt.Println("A utility to download windows steam games on MacOS")
	fmt.Println()
	fmt.Println("help		displays commands and their purposes")
	fmt.Println("config		prints config")
	fmt.Println("	username	<VAL>		changes steam username")
	fmt.Println("	install_dir	<VAL>		changes install directory")
	fmt.Println("	manifest	<VAL>		true | false copies manifest to steamapps directory")
	fmt.Println("	steamapps_dir	<VAL>		changes steamapps directory")
	fmt.Println("install		installs a steam game to the directory")
	fmt.Println("uninstall	uninstall a steam game to the directory")
}

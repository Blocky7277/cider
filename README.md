# cider

A tool for steamcmd that downloads windows game files to a specified directory.
Also copies the manifest after downloading. <br>

**cider requires you to install steamcmd before you can use it**

## Features

```
- cider -
A utility to download windows steam games on MacOS

help		displays commands and their purposes
config		prints config
	username	<VAL>		changes steam username
	install_dir	<VAL>		changes install directory
	manifest	<VAL>		true | false copies manifest to steamapps directory
	steamapps_dir	<VAL>		changes steamapps directory
install		installs a steam game to the directory
```

## Build

```
> git clone https://github.com/Blocky7277/cider.git
> cd cider
> go build cider
```
To use cider anywhere add it to your PATH

## Example Usage

### Config

```
> cider config username blocky7277
Updated username

> cider config install_dir ~/Games/steamapps/common
Updated install_dir

> cider config manifest true
Updated manifest

> cider config steamapps_dir ~/Games/steamapps
Updated steamapps_dir
```

### Install

Installing the game REPO
```
> cider install
Enter Steam password:
Enter Steam AppID:
3241660
Enter Steam game name:
repo
Working...
[STEAMCMD OUTPUT]
Finished...
```

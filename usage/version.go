package usage

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"github.com/google/go-github/github"
)

var (
	green      = "\033[32m"
	red        = "\033[31m"
	blue       = "\033[34m"
	white      = "\033[37m"
	updateFlag string
)

func CheckVersion() string {
	currVersion := "v1.0.2"
	return currVersion
}

func Version() {
	currVersion := CheckVersion()
	latestVersion := GetLatestVersion()
	fmt.Printf("[ %vINFO%v ] Current Version: %v\n", blue, white, currVersion)
	fmt.Printf("[ %vINFO%v ] Latest Version: %v\n", blue, white, latestVersion)

	if strings.Compare(currVersion, latestVersion) != 0 {
		fmt.Printf("[ %vINFO%v ] A newer version is available: %v\n", green, white, latestVersion)
		fmt.Printf("[ %vQUESTION%v ] Would you like to update to the latest version? (y/n): ", green, white)
		fmt.Scan(&updateFlag)

		if strings.ToLower(updateFlag) == "y" || strings.ToLower(updateFlag) == "Y" {
			fmt.Printf("[ %vINFO%v ] Updating to the latest version. Please wait...\n", blue, white)
			if !UpdateSeek() {
				return
			}
			fmt.Printf("[ %vINFO%v ] Update successfull! You are now using version: %v\n", green, white, latestVersion)
		} else {
			fmt.Printf("[ %vINFO%v ] Okay as you command sire !\n", blue, white)
		}
	} else {
		fmt.Printf("[ %vSUCCESS%v ] You are already using the latest version: %v\n", blue, white, latestVersion)
	}
}

func GetLatestVersion() string {
	owner := "iaakanshff"
	repo := "seek"

	ctx := context.Background()
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		fmt.Println("Error fetching the latest release:", err)
		return ""
	}

	return *release.TagName
}

func Update() {
	currVersion := CheckVersion()
	latestVersion := GetLatestVersion()

	if strings.Compare(currVersion, latestVersion) != 0 {
		fmt.Printf("[ %vINFO%v ] Checking for updates...\n", blue, white)
		fmt.Printf("[ %vINFO%v ] A newer version is available: %v\n", green, white, latestVersion)
		fmt.Printf("[ %vINFO%v ] Updating to the latest version. Please wait...\n", blue, white)
		if !UpdateSeek() {
			return
		}
		fmt.Printf("[ %vSUCCESS%v ] Update successfull! You are now using version: %v\n", green, white, latestVersion)
	} else {
		fmt.Printf("[ %vDONE%v ] You are already using the latest version: %v\n", blue, white, latestVersion)
	}
}

func UpdateSeek() bool{
	latestVersion := GetLatestVersion()

	cmd := exec.Command("go", "install", "github.com/iaakanshff/seek@"+latestVersion)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ %vINFO%v ] Update failed: %v\n", err, red, white)
		return false
	}

	userDir, _ := os.UserHomeDir()
	cmd = exec.Command("sudo", "cp", userDir+"/go/bin/seek", "/bin/")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("[ %vINFO%v ] Update failed: %v\n", err, red, white)
		return false
	}
	
	return true
}

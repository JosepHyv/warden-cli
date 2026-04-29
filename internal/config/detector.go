package config

import (
	_ "embed"
	"errors"
	"os"
	"os/exec"
)


type PackageManager string

type PackageManagerInfo struct {
	PManager PackageManager
	SecurityConfig string
	FileName string
}

//go:embed templates/default/.npmrc
var Npmrc string

//go:embed templates/yarn/.yarnrc.yml
var Yarnrc string

const (
	Npm  PackageManager = "npm"
	Bun  PackageManager = "bun"
	Pnpm PackageManager = "pnpm"
	Yarn PackageManager = "yarn"
)

var PackageManagers = [4]PackageManager{Npm, Bun, Pnpm, Yarn}


var lockfiles = map[PackageManager]string{
	Npm:  "package-lock.json",
	Pnpm: "pnpm-lock.yaml",
	Yarn: "yarn.lock",
	Bun: "bun.lock",
}

func GetSecurityConfig(pm PackageManager) string {
	switch pm {
		case Yarn:
			return Yarnrc
		default:
			return Npmrc
	}
}

func DetectPackageManager() (PackageManagerInfo, error) {
	priority := []PackageManager{Npm, Pnpm, Yarn, Bun}

	for _, pm := range priority {
		filename := lockfiles[pm]
		if _, err := os.Stat(filename); err == nil {
			return PackageManagerInfo{PManager: pm, SecurityConfig: GetSecurityConfig(pm), FileName: filename}, nil
		}
	}

	return PackageManagerInfo{}, errors.New("Undefined Package Manager")
}


func GetSystemPackageManagers() []PackageManager{
	installed := make([]PackageManager, 0, len(PackageManagers))

	// only unix-like systems
	for _, pm := range PackageManagers {
		command := exec.Command("sh", "-c", "command -v " + string(pm))
		_, err := command.Output()
		if err == nil { // i don't know why this is happening
			installed = append(installed, pm)
		}
	}
	// return installed
	return installed
}



func GetConfigFile(pm PackageManager) string {
	switch pm {
		case Yarn:
			return ".yarnrc.yml"
		default:
			return ".npmrc"
	}
}

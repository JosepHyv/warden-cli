package config

import (
	_ "embed"
	"errors"
	"os"
)


type PackageManager string

type PackageManagerInfo struct {
	packageManager PackageManager
	filename string
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
	None PackageManager = "none"
)




var lockfiles = map[PackageManager]string{
	Npm:  "package-lock.json",
	Pnpm: "pnpm-lock.yaml",
	Yarn: "yarn.lock",
	Bun: "bun.lock",
}

func DetectPackageManager() (PackageManagerInfo, error) {
	priority := []PackageManager{Npm, Pnpm, Yarn, Bun}

	for _, pm := range priority {
		filename := lockfiles[pm]
		if _, err := os.Stat(filename); err == nil {
			return PackageManagerInfo{packageManager: pm, filename: filename}, nil
		}
	}

	return PackageManagerInfo{}, errors.New("Undefined Package Manager")
}

func GetConfigFile(pm PackageManager) string {
	switch pm {
	case Yarn:
		return ".yarnrc.yml"
	case Bun:
		return ".bunconf.toml"
	default:
		return ".npmrc"
	}
}

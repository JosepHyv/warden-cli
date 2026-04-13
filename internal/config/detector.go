package config

import (
	"os"
)

type PackageManager string

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

func DetectPackageManager() (PackageManager, string) {
	priority := []PackageManager{Npm, Pnpm, Yarn}

	for _, pm := range priority {
		filename := lockfiles[pm]
		if _, err := os.Stat(filename); err == nil {
			return pm, filename
		}
	}

	return None, ""
}

func GetConfigFile(pm PackageManager) string {
	switch pm {
	case Npm:
		return ".npmrc"
	case Pnpm:
		return ".npmrc"
	case Yarn:
		return ".yarnrc"
	default:
		return ""
	}
}

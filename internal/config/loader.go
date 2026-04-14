package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type SecurityConfig map[string]interface{}

func LoadTemplate(pm PackageManager) (SecurityConfig, string, error) {
	execPath, err := os.Executable()
	if err != nil {
		execPath, _ = os.Getwd()
	}

	templateDir := filepath.Join(filepath.Dir(execPath), "internal", "templates")
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		templateDir = filepath.Join(".", "internal", "template")
	}

	templateFile := map[PackageManager]string{
		Npm:  "npm.yaml",
		Pnpm: "pnpm.yaml",
		Yarn: "yarn.yaml",
	}[pm]

	templatePath := filepath.Join(templateDir, templateFile)

	data, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read template: %w", err)
	}

	var config SecurityConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, "", fmt.Errorf("failed to parse template: %w", err)
	}

	return config, templatePath, nil
}

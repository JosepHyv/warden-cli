
package config
//
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )
//
// func WriteConfig(configFile string, config SecurityConfig, force bool) error {
// 	existing := make(SecurityConfig)
//
// 	if !force {
// 		if data, err := os.ReadFile(configFile); err == nil {
// 			parseNpmrc(string(data), existing)
// 		}
// 	}
//
// 	for key, value := range config {
// 		existing[key] = value
// 	}
//
// 	var sb strings.Builder
// 	for key, value := range existing {
// 		sb.WriteString(fmt.Sprintf("%s=%v\n", key, value))
// 	}
//
// 	if err := os.WriteFile(configFile, []byte(sb.String()), 0644); err != nil {
// 		return fmt.Errorf("failed to write config file: %w", err)
// 	}
//
// 	return nil
// }
//
// func parseNpmrc(content string, config SecurityConfig) {
// 	lines := strings.Split(content, "\n")
// 	for _, line := range lines {
// 		line = strings.TrimSpace(line)
// 		if line == "" || strings.HasPrefix(line, "#") {
// 			continue
// 		}
// 		parts := strings.SplitN(line, "=", 2)
// 		if len(parts) == 2 {
// 			key := strings.TrimSpace(parts[0])
// 			value := strings.TrimSpace(parts[1])
// 			config[key] = value
// 		}
// 	}
// }
//
// func GetAppliedKeys(config SecurityConfig) []string {
// 	keys := make([]string, 0, len(config))
// 	for key := range config {
// 		keys = append(keys, key)
// 	}
// 	return keys
// }

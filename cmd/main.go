package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/JosepHyv/warden-cli/internal/config"
	"github.com/manifoldco/promptui"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "warden",
		Short: "Warden is a security-first package installer for Node.js",
		Long: `Warden protects your Node.js supply chain by auditing CVEs
in real-time and enforcing safe configurations before any package is installed.`,
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize safe project configurations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🛡️ Warden: Checking and creating safe .npmrc configurations...")
			pm, err := config.DetectPackageManager()
			if err != nil {
				fmt.Println("The current project does not have a package manager yet!")
				prompt := promptui.Select{
					Label: "Select one of your installed packagemanagers",
					Items: config.GetSystemPackageManagers(),
				}

				_, result, err := prompt.Run()
				if err != nil {
					fmt.Println("Error")
				}
				fmt.Println(result)

			}

			fmt.Println(pm.SecurityConfig)


		},
	}

	var installCmd = &cobra.Command{
		Use:   "install [packages]",
		Short: "Install packages with real-time CVE auditing",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🔍 Warden: Auditing dependencies before installation...")
			// Aquí irá tu lógica para detectar el lockfile y consultar OSV.dev
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(installCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

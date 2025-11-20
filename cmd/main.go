package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CaioDS/fli/internal/application/commands"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func main() {
	var osContext = context.CreateOSContext()
	var fileService = services.CreateSystemFileService(osContext)
	var versionsService = services.CreateVersionsService(*osContext)

	var rootCommand = &cobra.Command{}
	var version string

	var cmd = &cobra.Command{
		Use: "install",
		Short: "Install flutter framework.",
		Run: func (cmd *cobra.Command, args []string) {
			if version == "" {
				fmt.Println("You must specify flutter version")
				return
			}

			fmt.Println("\nInstalling flutter...")
			commands.HandleInstallCommand(*fileService, *osContext)
		},
	}

	cmd.Flags().StringVarP(&version, "version", "v", "", "Version to be installed")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()
}
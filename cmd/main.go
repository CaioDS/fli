package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CaioDS/fli/internal/application/commands"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func main() {
	// CONTEXTs
	var osContext = context.CreateOSContext()
	var localDBContext = context.CreateLocalDBContext("fli.db")

	// SERVICES 
	var fileService = services.CreateSystemFileService(osContext, localDBContext)
	var versionsService = services.CreateVersionsService(osContext, localDBContext)
	
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
			commands.HandleInstallCommand(*fileService, *versionsService, *osContext, version)

			fmt.Println("\nFlutter SDK v", version, " installed with success!")
			fmt.Println("\nUse 'flutter doctor' on terminal to check the status")
		},
	}

	cmd.Flags().StringVarP(&version, "version", "v", "", "Version to be installed")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()
}
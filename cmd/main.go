package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/CaioDS/fli/internal/application/commands"
	"github.com/CaioDS/fli/internal/infrastructure/config"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/repositories"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func main() {
	var env = config.Get()

	// CONTEXTs
	var osContext = context.NewOSContext()
	var localDBContext = context.NewLocalDBContext("fli.db")
	dbContext, err := context.NewDbContext(env.DynamoEndpoint, env.DynamoRegion)
 	if err != nil {
		panic("Failed to stablish a connection with the database")
	}

	// REPOSITORIES
	var localRepository = repositories.NewLocalRepository(localDBContext)
	var versionsRepository = repositories.NewVersionsRepository(dbContext)

	// SERVICES 
	var fileService = services.NewSystemFileService(*osContext, localRepository)
	var versionsService = services.NewVersionsService(*osContext, localRepository, versionsRepository)
	
	var rootCommand = &cobra.Command{}
	var version string

	var installCmd = &cobra.Command{
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

	installCmd.Flags().StringVarP(&version, "version", "v", "", "Version to be installed")

	var listVersionCmd = &cobra.Command{
		Use: "list",
		Short: "List installed versions",
		Run: func (cmd *cobra.Command, args []string) {
			commands.HandleListVersionsCommand(*versionsService)
			return
		},
	}

	rootCommand.AddCommand(installCmd)
	rootCommand.AddCommand(listVersionCmd)
	rootCommand.Execute()
}
package commands

import (
	"fmt"

	"github.com/CaioDS/fli/internal/application/usecases"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func HandleListVersionsCommand(
	versionsService services.VersionsService,
) {
	output, error := usecases.ListInstalledVersions(versionsService)
	if error != nil {
		panic(error)
	}

	fmt.Println("Installed Flutter Version")
	fmt.Println("Version | Location")
	for _, element := range output {
		fmt.Println(string(element.Key), "|", string(element.Value))
	}
}
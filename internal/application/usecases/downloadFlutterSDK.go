package usecases

import (
	"fmt"

	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func DownloadFlutterSDK(
	service services.VersionsService, 
	context context.OSContext,
	version string,
	destiny string,
) {
	fmt.Println("Fetching flutter version...")
	err := service.DownloadVersion(version, destiny)
	if err != nil {
		fmt.Println("ERROR"+err.Error())
		panic(err)
	}
}
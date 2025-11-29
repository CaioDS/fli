package commands

import (
	"github.com/CaioDS/fli/internal/application/usecases"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)


func HandleInstallCommand(
	systemFileService services.SystemFileService, 
	versionService services.VersionsService,
	context context.OSContext,
	version string,
) {
	defaultPath := usecases.CreateDefaultSDKLocation(systemFileService, context)

	parsedDefaultPath := defaultPath+"/"+version

	// usecases.DownloadFlutterSDK(versionService, context, version, parsedDefaultPath)	
	usecases.CreatePathEnvVariable(context, parsedDefaultPath)
	
}
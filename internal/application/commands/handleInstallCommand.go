package commands

import (
	"github.com/CaioDS/fli/internal/application/usecases"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)


func HandleInstallCommand(service services.SystemFileService, context context.OSContext) {
	_ = usecases.CreateDefaultSDKLocation(service, context)
	
	
}
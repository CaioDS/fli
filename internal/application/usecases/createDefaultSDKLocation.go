package usecases

import (
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func CreateDefaultSDKLocation(service services.SystemFileService) {
	service.CreateDefaultSDKLocation()
}
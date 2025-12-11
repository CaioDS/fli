package usecases

import (
	"github.com/CaioDS/fli/internal/domain/dto"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func ListInstalledVersions(
	service services.VersionsService,
) ([]dto.ListDto, error) {
	return service.GetInstalledVersions()
}
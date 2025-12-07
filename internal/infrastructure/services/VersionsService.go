package services

import (
	"errors"
	"log"

	"github.com/CaioDS/fli/internal/domain/models"
	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/repositories"
	"github.com/hashicorp/go-getter"
)

type VersionsService struct {
	system string
	arch string
	localRepository *repositories.LocalRepository
	versionsRepository *repositories.VersionsRepository
}

type Version struct {
	Version string
	Link string
}

func NewVersionsService(
	osContext context.OSContext, 
	localRepository *repositories.LocalRepository,
	versionsRepository *repositories.VersionsRepository,
) *VersionsService {
	system := osContext.GetOSSystem()
	arch := osContext.GetArchSystem()

	return &VersionsService{
		system: system,
		arch: arch,
		localRepository: localRepository,
		versionsRepository: versionsRepository,
	}
}

func (v *VersionsService) DownloadVersion(version string, destiny string) error {
	data, err := v.versionsRepository.GetVersion(version)
	if err != nil {
		return err
	}

	var link string = v.getLinkByOS(*data)

	log.Println("Downloading version ", version)

	client := &getter.Client{
		Src: link,
		Dst: destiny,
		Mode: getter.ClientModeDir,
	}

	err = client.Get()
	if err != nil {
		return errors.New("failed to download flutter sdk")
	}

	err = v.localRepository.SaveVersionRegistry(version, destiny)

	log.Println("Download finished!")
	log.Println("Saved in: "+destiny)
	return err
}

func (v *VersionsService) getLinkByOS(version models.Version) string {
	switch v.system {
	case "darwin":
		if v.arch == "arm64" {
			return version.Darwin_arm64
		} else {
			return version.Darwin_intel
		}
	case "linux":
		if v.arch == "arm64" {
			return version.Linux_arm64
		} else {
			return version.Linux_x64
		}
	default:
		return version.Windows
	}
}

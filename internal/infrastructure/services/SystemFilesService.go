package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CaioDS/fli/internal/infrastructure/context"
)

type SystemFileService struct {
	osContext *context.OSContext
}

func CreateSystemFileService(osContext *context.OSContext) *SystemFileService {
	return &SystemFileService{
		osContext: osContext,
	}
}

func (s *SystemFileService) CreateDefaultSDKLocation() (string, error) {
	var system = s.osContext.GetOSSystem()
	switch system {
	case "windows":
		return s.CreateDefaultSDKLocationWindows()
	case "darwin":
		return s.CreateDefaultSDKLocationDarwinOrLinux()
	case "linux":
		return s.CreateDefaultSDKLocationDarwinOrLinux()
	default:
		panic("Unsupported system")
	}
}

func (s *SystemFileService) CreateDefaultSDKLocationWindows() (string, error) {
	var base = os.Getenv("LOCALAPPDATA")
	if(base == "") {
		var userProfile = os.Getenv("USERPROFILE")
		if(userProfile == "") {
			return "", errors.New("USERPROFILE environment variable is not set")
		}
		base = filepath.Join(userProfile, "AppData", "Local")
	}
	return filepath.Join(base, "development/test"), nil
}

func (s *SystemFileService) CreateDefaultSDKLocationDarwinOrLinux() (string, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Failed to get home directory: %s\n", err)
		return "", errors.New("failed to get home directory")
	}

	developmentDirectory := filepath.Join(homeDirectory, "development")+"/test"
	var permissionCode = os.FileMode(0755)
	err = os.MkdirAll(developmentDirectory, permissionCode)
	if err != nil {
		fmt.Printf("Failed to create development directory: %s\n", err)
		return "", errors.New("failed to create development directory: "+developmentDirectory)
	}

	return developmentDirectory, nil
}
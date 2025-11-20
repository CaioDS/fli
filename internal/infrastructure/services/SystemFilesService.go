package services

import (
	"errors"
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

func (s *SystemFileService) GetDefaultPath(osContext context.OSContext) (string, error) {
	system := osContext.GetOSSystem()
	switch system {
	case "windows":
		return "C:/src", nil

	case "darwin":
	case "linux":
		var userProfile = os.Getenv("USERPROFILE")
		var dir = filepath.Join(userProfile, "development")
		return dir, nil

	default: 
		return "", errors.New("unsupported system")
	}
	return "", nil
}

func (s *SystemFileService) CreateCustomPath(path string) (string, error) {
	var permissionCode = os.FileMode(0755)

	err := os.MkdirAll(path, permissionCode)
	if err != nil {
		return "", errors.New("failed to create custom directory: "+path)
	}
	
	return path, nil
}
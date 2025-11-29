package services

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/CaioDS/fli/internal/infrastructure/context"
)

type SystemFileService struct {
	osContext *context.OSContext
	dbContext *context.LocalDbContext
}

func CreateSystemFileService(
	osContext *context.OSContext, 
	dbContext *context.LocalDbContext,
) *SystemFileService {
	return &SystemFileService{
		osContext: osContext,
		dbContext: dbContext,
	}
}

func (s *SystemFileService) GetDefaultPath(osContext context.OSContext) (string, error) {
	path, err := s.retrieveStoredDefaultPath()
	if err == nil {
		return path, nil
	}
	
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

	err = s.saveDefaultPath(path)
	return path, err
}

func (s *SystemFileService) retrieveStoredDefaultPath() (string, error) {
	path, err := s.dbContext.Get("config", []byte("defaultPath"))
	if err != nil {
		return "", err
	}

	return string(path), nil
}

func (s *SystemFileService) saveDefaultPath(path string) error {
	err := s.dbContext.CreateBucket("config")
	if err != nil {
		return errors.New("Failed to create config bucket")
	}

	err = s.dbContext.Put("config", []byte("defaultPath"), []byte(path))
	if err != nil {
		return errors.New("failed to save default path on database!")
	}

	return nil
}
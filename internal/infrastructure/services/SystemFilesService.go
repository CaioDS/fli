package services

import (
	"os"

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

func (s *SystemFileService) CreateDefaultSDKLocation() {
	var system = s.osContext.GetOSSystem()
	if system == "windows" {
		os.MkdirAll("C:/src", os.ModePerm)
	} else if system == "mac" || system == "linux" {
		os.MkdirAll("development/flutter", os.ModePerm)
	} else {
		panic("Unsupported system")
	}
}

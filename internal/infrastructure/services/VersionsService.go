package services

import (
	"encoding/json"
	"os"

	"github.com/CaioDS/fli/internal/infrastructure/context"
)

type VersionsService struct {
	Versions []map[string]interface{}
}

func CreateVersionsService(osContext context.OSContext) *VersionsService {
	system := osContext.GetOSSystem()
	
	var versions []byte
	var err error
	
	switch system {
	case "windows":
		versions, err = os.ReadFile("internal\\infrastructure\\datasource\\windowsVersions.json")
	case "darwin":
		versions, err = os.ReadFile("internal\\infrastructure\\datasource\\macOsVersions.json")
	case "linux":
		versions, err = os.ReadFile("internal\\infrastructure\\datasource\\linuxVersions.json")
	}

	if err != nil {
		panic("failed to read system flutter versions")
	}
	
	var data []map[string]interface{}
	err = json.Unmarshal(versions, &data)
		if err != nil {
		panic("failed to parse system flutter versions")
	}

	return &VersionsService{
		Versions: data,
	}
}

func (v *VersionsService) DownloadVersion() {
	
}

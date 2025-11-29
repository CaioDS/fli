package services

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"slices"

	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/hashicorp/go-getter"
)

type VersionsService struct {
	Versions []Version
	dbContext *context.LocalDbContext
}

type Version struct {
	Version string
	Link string
}

func CreateVersionsService(
	osContext *context.OSContext, 
	dbContext *context.LocalDbContext,
) *VersionsService {
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
	
	var data []Version;
	err = json.Unmarshal(versions, &data)
		if err != nil {
		panic("failed to parse system flutter versions")
	}

	return &VersionsService{
		Versions: data,
		dbContext: dbContext,
	}
}

func (v *VersionsService) DownloadVersion(version string, destiny string) error {
	var index int = slices.IndexFunc(v.Versions, func(item Version) bool {
		return item.Version == version
	})
	if index == -1 {
		return errors.New("version not found")
	}

	element := v.Versions[index]

	log.Println("Downloading version from this source: "+element.Link)

	client := &getter.Client{
		Src: element.Link,
		Dst: destiny,
		Mode: getter.ClientModeDir,
	}

	err := client.Get()
	if err != nil {
		return errors.New("failed to download flutter sdk")
	}

	err = v.saveVersionResgistry(version, destiny)

	log.Println("Download was finished!")
	log.Println("Saved in: "+destiny)
	return err
}

func (v *VersionsService) saveVersionResgistry(version string, path string) error {
	err := v.dbContext.CreateBucket("versions")
	if err != nil {
		return errors.New("Failed to create versions bucket")
	}

	err = v.dbContext.Put("versions", []byte(version), []byte(path))
	if err != nil {
		return errors.New("failed to save version occurency on database!")
	}

	return nil
}

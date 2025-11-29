package usecases

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/CaioDS/fli/internal/infrastructure/context"
	"golang.org/x/sys/windows/registry"
)

func CreatePathEnvVariable(context context.OSContext, pathVariable string) error {
	absolutePathVariable := filepath.FromSlash(pathVariable+"/flutter/bin")
	key, _, err := registry.CreateKey(
		registry.LOCAL_MACHINE,
		`SYSTEM\CurrentControlSet\Control\Session Manager\Environment`,
		registry.ALL_ACCESS,
	)
	if err != nil {
		log.Fatal("Failed to add create new path variable on the system: "+err.Error())
		return err
	}

	// GET existing PATH variables
	currentPath, _, err := key.GetStringValue("Path")
	if err != nil && err != registry.ErrNotExist {
		log.Fatal("Failed to get the existing PATH variables: "+err.Error())
		return err
	}

	// AVOID duplicated variables
	paths := strings.Split(currentPath, ";")
	for _, p := range paths {
		if strings.EqualFold(p, absolutePathVariable) {
			return nil
		}
	}

	newValue := currentPath + absolutePathVariable
	key.SetStringValue("Path", newValue)
	
	return nil
}
package usecases

import (
	"github.com/AlecAivazis/survey/v2"

	"github.com/CaioDS/fli/internal/infrastructure/context"
	"github.com/CaioDS/fli/internal/infrastructure/services"
)

func CreateDefaultSDKLocation(service services.SystemFileService, context context.OSContext) (string) {
	defaultPath, err := service.GetDefaultPath(context)
	if err != nil {
		panic(err)
	}

	var pathInput string
	survey.AskOne(
		&survey.Input{
			Message: "Where should flutter be installed?",
			Default: defaultPath,
		},
		&pathInput,
	)

	path, err := service.CreateCustomPath(pathInput)
	if err != nil {
		panic(err)
	}

	return path
}


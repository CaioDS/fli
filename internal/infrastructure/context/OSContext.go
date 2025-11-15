package context

import (
	"runtime"
)

type OSContext struct {}

func CreateOSContext() *OSContext {
	return &OSContext{}
}

func (o *OSContext) GetOSSystem() string {
	return runtime.GOOS;
}
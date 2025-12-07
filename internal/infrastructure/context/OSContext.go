package context

import (
	"runtime"
)

type OSContext struct {}

func NewOSContext() *OSContext {
	return &OSContext{}
}

func (o *OSContext) GetOSSystem() string {
	return runtime.GOOS;
}

func (o *OSContext) GetArchSystem() string {
	return runtime.GOARCH;
}
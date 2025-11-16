package context

import (
	"runtime"
	"fmt"
)

type OSContext struct {}

func CreateOSContext() *OSContext {
	return &OSContext{}
}

func (o *OSContext) GetOSSystem() string {
	fmt.Println("System: %s", runtime.GOOS)
	return runtime.GOOS;
}
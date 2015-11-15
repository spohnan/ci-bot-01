// Package paas contains all the Platform as a Service specific code
package paas

import (
	"log"
	"net/http"
)

type NoOpPaas struct {
	Name string
}

func (noop NoOpPaas) String() string {
	return noop.Name
}

func (noop NoOpPaas) Log(r *http.Request, format string, args ...interface{}) {
	log.Printf(format, args)
}

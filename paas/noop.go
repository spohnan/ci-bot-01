// Package paas contains all the Platform as a Service specific code
package paas

import (
	"log"
	"net/http"
)

// NoOpPaas implements all of our Paas interface methds and does pretty much nothing interesting
type NoOpPaas struct {
	Name string
}

// String returns the name of this Paas implementation
func (noop NoOpPaas) String() string {
	return noop.Name
}

// Log in this case just sends it to the default internal Go logger
func (noop NoOpPaas) Log(r *http.Request, format string, args ...interface{}) {
	log.Printf(format, args)
}

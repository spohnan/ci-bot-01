// Package paas contains all the Platform as a Service specific code
package paas

import (
	"net/http"
	"os"
)

// Paas is our interace we can swap out service provider specific implementations
// This probabably isn't really portable as is but gives at least the ability to
// swap with a no-op logger that does nothing if not running in GAE environment
type Paas interface {
	Log(r *http.Request, format string, args ...interface{})
	String() string
}

// PaasCtx is our handle to the specific service provider API
var PaasCtx Paas

func init() {
	initPaasVendorCtx()
}

func initPaasVendorCtx() {
	switch os.Getenv("PAAS_VENDOR") {
	default:
		PaasCtx = NoOpPaas{"NoOpPaas"}
	case "GoogleAppEngine":
		PaasCtx = GoogleAppEngine{"GoogleAppEngine"}
	}
}

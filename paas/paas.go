// Package paas contains all the Platform as a Service specific code
package paas

import (
	"net/http"
	"os"
)

// This probabably isn't really portable as is but gives at least the ability to
// swap with a no-op logger that does nothing if not running in GAE environment at the moment
type Paas interface {
	Log(r *http.Request, format string, args ...interface{})
	String() string
}

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

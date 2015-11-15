// Package paas contains all the Platform as a Service specific code
package paas

import (
	"net/http"
	"os"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type GoogleAppEngine struct {
	Name string
}

func (gae GoogleAppEngine) String() string {
	return gae.Name
}

// TODO: Pass along a severity code, everything is info for now
func (gae GoogleAppEngine) Log(r *http.Request, format string, args ...interface{}) {
	if isLoggingEnabled() {
		c := appengine.NewContext(r)
		log.Infof(c, format, args)
	}
}

// I don't really want to do GAE integration tests if I'm
// not yet using any of the advanced functionality. For now
// I'm just going to put a guard around logging calls so I
// can continue to use standard go unit tests
func isLoggingEnabled() bool {
	b, err := strconv.ParseBool(os.Getenv("ENABLE_GAE_LOGGING"))
	if err != nil {
		return false
	}
	return b
}

// Package paas contains all the Platform as a Service specific code
package paas

import (
	"testing"
)

func TestNoOpPaasVendorSet(t *testing.T) {
	initPaasVendorCtx()

	if PaasCtx.String() != "NoOpPaas" {
		t.Errorf("NoOpPaas vendor name incorrectly set, returned: %s", PaasCtx.String())
	}

	PaasCtx.Log(nil, "test %s", "test")
}

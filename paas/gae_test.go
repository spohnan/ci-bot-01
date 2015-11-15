package paas

import (
	"os"
	"testing"
)

func TestGAEPaasVendorSet(t *testing.T) {
	defer cleanup()
	os.Setenv("ENABLE_GAE_LOGGING", "false")
	os.Setenv("PAAS_VENDOR", "GoogleAppEngine")
	initPaasVendorCtx()

	if PaasCtx.String() != "GoogleAppEngine" {
		t.Errorf("GoogleAppEngine Paas vendor name incorrectly set, returned: %s", PaasCtx.String())
	}
}

func TestIsLoggingEnabled(t *testing.T) {
	defer cleanup()
	os.Setenv("ENABLE_GAE_LOGGING", "false")
	os.Setenv("PAAS_VENDOR", "GoogleAppEngine")
	initPaasVendorCtx()

	os.Setenv("ENABLE_GAE_LOGGING", "true")
	if isLoggingEnabled() != true {
		t.Error("isLoggingEnabled should have returned true")
	}

	cleanup()

	if isLoggingEnabled() != false {
		t.Error("isLoggingEnabled should have returned true")
	}

	PaasCtx.Log(nil, "test %s", "test")
}

func cleanup() {
	os.Unsetenv("ENABLE_GAE_LOGGING")
	os.Unsetenv("PAAS_VENDOR")
}

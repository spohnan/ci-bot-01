package bot

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const (
	allowedIP = "127.0.0.1"
	deniedIP  = "54.2.2.2"
)

func TestHelthCheck(t *testing.T) {
	defer cleanup()
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", authWrapper(healthHandler))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health", nil)
	request.RemoteAddr = allowedIP
	mux.ServeHTTP(writer, request)
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
	}
	if writer.Body.String() != HealthOK {
		t.Errorf("Response to /health was %s and should be %s", writer.Body.String(), HealthOK)
	}
}

func TestAllowedRemoteIP(t *testing.T) {
	defer cleanup()
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", authWrapper(webHookHandler))
	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"foo","bar"}`)
	request, _ := http.NewRequest("POST", "/webhook", json)
	request.RemoteAddr = allowedIP
	mux.ServeHTTP(writer, request)
	if writer.Code != http.StatusOK {
		t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
	}
}

func TestDeniedRemoteIP(t *testing.T) {
	defer cleanup()
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", authWrapper(webHookHandler))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health", nil)
	request.RemoteAddr = deniedIP
	mux.ServeHTTP(writer, request)
	if writer.Code != http.StatusForbidden {
		t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
	}
}

func TestDeniedNoWhitelist(t *testing.T) {
	defer cleanup()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", authWrapper(webHookHandler))
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/health", nil)
	request.RemoteAddr = deniedIP
	mux.ServeHTTP(writer, request)
	if writer.Code != http.StatusForbidden {
		t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
	}
}

func cleanup() {
	os.Unsetenv("CI_BOT_IP_WHITELIST")
}

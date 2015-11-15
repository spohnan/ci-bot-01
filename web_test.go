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
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	withWebTestContext(func(mux *http.ServeMux, writer *httptest.ResponseRecorder) {
		mux.HandleFunc("/health", authWrapper(healthHandler))
		request, _ := http.NewRequest("GET", "/health", nil)
		request.RemoteAddr = allowedIP
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusOK {
			t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
		}
		if writer.Body.String() != HealthOK {
			t.Errorf("Response to /health was %s and should be %s", writer.Body.String(), HealthOK)
		}
	})
}

func TestAllowedRemoteIP(t *testing.T) {
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	withWebTestContext(func(mux *http.ServeMux, writer *httptest.ResponseRecorder) {
		mux.HandleFunc("/webhook", authWrapper(webHookHandler))
		json := strings.NewReader(`{"foo","bar"}`)
		request, _ := http.NewRequest("POST", "/webhook", json)
		request.RemoteAddr = allowedIP
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusOK {
			t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
		}
	})
}

func TestDeniedRemoteIP(t *testing.T) {
	os.Setenv("CI_BOT_IP_WHITELIST", allowedIP)

	withWebTestContext(func(mux *http.ServeMux, writer *httptest.ResponseRecorder) {
		mux.HandleFunc("/health", authWrapper(webHookHandler))
		request, _ := http.NewRequest("GET", "/health", nil)
		request.RemoteAddr = deniedIP
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
		}
	})
}

func TestDeniedNoWhitelist(t *testing.T) {
	withWebTestContext(func(mux *http.ServeMux, writer *httptest.ResponseRecorder) {
		mux.HandleFunc("/health", authWrapper(webHookHandler))
		request, _ := http.NewRequest("GET", "/health", nil)
		request.RemoteAddr = deniedIP
		mux.ServeHTTP(writer, request)
		if writer.Code != http.StatusForbidden {
			t.Errorf("Response code is %v for address %s", writer.Code, request.RemoteAddr)
		}
	})
}

// A helper function to do the setup and teardown work within a web test
func withWebTestContext(fn func(mux *http.ServeMux, writer *httptest.ResponseRecorder)) {
	defer cleanup()
	mux := http.NewServeMux()
	writer := httptest.NewRecorder()
	fn(mux, writer)
}

// Reset the environment after each test
func cleanup() {
	os.Unsetenv("CI_BOT_IP_WHITELIST")
}

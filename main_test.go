package bot

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMainHandler(t *testing.T) {
	os.Setenv("CI_BOT_IP_WHITELIST", "127.0.0.1")
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", mainHandler)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)
	request.RemoteAddr = "127.0.0.1"
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v for address %v", writer.Code, request.RemoteAddr)
	}
}

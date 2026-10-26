package demo

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPageHeadlineAndVersion(t *testing.T) {
	page, err := RenderPage("8be41d0c0ffee", "2026-10-26T00:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	for _, text := range []string{Headline, "version 8be41d0", "$15.20", "2026-10-26T00:00:00Z"} {
		if !strings.Contains(string(page), text) {
			t.Fatalf("page missing %q", text)
		}
	}
}

func TestHealth(t *testing.T) {
	response := httptest.NewRecorder()
	Handler("v2", "now").ServeHTTP(response, httptest.NewRequest("GET", "/health", nil))
	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if response.Code != 200 || body["status"] != "ok" || body["version"] != "v2" {
		t.Fatalf("unexpected health: %d %v", response.Code, body)
	}
}

func TestHTTPPage(t *testing.T) {
	response := httptest.NewRecorder()
	Handler("v1", "now").ServeHTTP(response, httptest.NewRequest("GET", "/", nil))
	if response.Code != 200 || !strings.Contains(response.Body.String(), "version v1") {
		t.Fatalf("unexpected page: %d %s", response.Code, response.Body.String())
	}
}

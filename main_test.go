package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	router := newRouter()
	for _, tc := range []struct {
		name, method, path string
		status             int
		body, contentType  string
	}{
		{"hello", http.MethodGet, "/", 200, "Hello World!", "text/plain; charset=utf-8"},
		{"health", http.MethodGet, "/healthz", 200, `{"status":"ok"}`, "application/json; charset=utf-8"},
		{"missing", http.MethodGet, "/missing", 404, "404 page not found", "text/plain"},
		{"unsupported method", http.MethodPost, "/", 404, "404 page not found", "text/plain"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			router.ServeHTTP(response, httptest.NewRequest(tc.method, tc.path, nil))
			if response.Code != tc.status || response.Body.String() != tc.body {
				t.Fatalf("got %d %q; want %d %q", response.Code, response.Body.String(), tc.status, tc.body)
			}
			if got := response.Header().Get("Content-Type"); got != tc.contentType {
				t.Errorf("Content-Type = %q; want %q", got, tc.contentType)
			}
		})
	}
}

func TestServerAddress(t *testing.T) {
	for _, tc := range []struct{ port, want string }{{"", ":3000"}, {"8080", ":8080"}} {
		t.Run(tc.want, func(t *testing.T) {
			t.Setenv("PORT", tc.port)
			if got := serverAddress(); got != tc.want {
				t.Fatalf("got %q; want %q", got, tc.want)
			}
		})
	}
}

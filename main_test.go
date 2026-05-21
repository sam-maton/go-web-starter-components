package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPagesRenderWithNavbar(t *testing.T) {
	handler, err := newHandler()
	if err != nil {
		t.Fatalf("newHandler() error = %v", err)
	}

	tests := []struct {
		path    string
		heading string
		title   string
	}{
		{path: "/", heading: "Home", title: "<title>Home</title>"},
		{path: "/about", heading: "About", title: "<title>About</title>"},
	}

	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
			}

			body := rr.Body.String()
			checks := []string{
				"<a href=\"/\">Home</a>",
				"<a href=\"/about\">About</a>",
				"<h1>" + tc.heading + "</h1>",
				tc.title,
			}

			for _, check := range checks {
				if !strings.Contains(body, check) {
					t.Fatalf("response body missing %q", check)
				}
			}
		})
	}
}

func TestStaticFilesAreServed(t *testing.T) {
	handler, err := newHandler()
	if err != nil {
		t.Fatalf("newHandler() error = %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/static/css/styles.css", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}

	if _, err := io.ReadAll(rr.Body); err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
}

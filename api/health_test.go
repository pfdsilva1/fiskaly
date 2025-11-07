package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealth_HappyPath(t *testing.T) {
    t.Parallel()
    s := &Server{}
    req := httptest.NewRequest(http.MethodGet, "/api/v0/health", nil)
    rr := httptest.NewRecorder()

    s.Health(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", rr.Code)
    }
}

func TestHealth_MethodNotAllowed(t *testing.T) {
    t.Parallel()
    s := &Server{}
    req := httptest.NewRequest(http.MethodPost, "/api/v0/health", nil)
    rr := httptest.NewRecorder()

    s.Health(rr, req)

    if rr.Code != http.StatusMethodNotAllowed {
        t.Fatalf("expected status 405, got %d", rr.Code)
    }
}


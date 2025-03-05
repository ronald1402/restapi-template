package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestSetUnhealthy(t *testing.T) {
	h := &HealthCheckHandler{
		mut: new(sync.RWMutex),
	}

	t.Run("Check Initial Healthy", func(t *testing.T) {
		if h.IsUnHealthy {
			t.Fatalf("Expected IsUnHealthy to be false initially, got true")
		}
	})

	t.Run("TestSetUnhealthy", func(t *testing.T) {
		h.SetUnhealthy()

		if !h.IsUnHealthy {
			t.Error("IsUnHealthy should be true, got false")
		}
	})
}

func TestCheck(t *testing.T) {
	h := &HealthCheckHandler{
		mut: new(sync.RWMutex),
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	t.Run("Check Initial Healthy", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := h.Check(c)
		if err != nil {
			t.Fatalf("Check Initial Healthy failed: %v", err)
		}
		if rec.Code != http.StatusOK {
			t.Fatalf("want %d, got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("Check Unhealthy", func(t *testing.T) {
		h.SetUnhealthy()
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := h.Check(c)
		if err != nil {
			t.Fatalf("Check Unhealthy failed: %v", err)
		}
		if rec.Code != http.StatusServiceUnavailable {
			t.Fatalf("want %d, got %d", http.StatusServiceUnavailable, rec.Code)
		}
	})
}

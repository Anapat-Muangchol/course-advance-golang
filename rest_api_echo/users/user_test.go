//go:build integration

package users_test

import (
	"api/users"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessWithGetWithRealServer(t *testing.T) {
	// Setup router
	e := echo.New()
	service := users.NewUserService()
	e.GET("/users", users.GetUserHandler(service))

	// Call API
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// Start server
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, rec.Code, 200)
	assert.Contains(t, rec.Body.String(), "Call get user")
}

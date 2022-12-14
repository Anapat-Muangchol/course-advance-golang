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

func TestSuccessWithGet(t *testing.T) {
	// Start server
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	c := e.NewContext(req, rec)

	repo := users.NewUserRepository()
	service := users.NewUserService(&repo)
	users.GetUserHandler(service)(c)

	// Assert
	assert.Equal(t, rec.Code, 200)
	assert.Contains(t, rec.Body.String(), "TODO next")
}

func TestSuccessWithGetWithRealServer(t *testing.T) {
	// Setup router
	e := echo.New()
	repo := users.NewUserRepository()
	service := users.NewUserService(&repo)
	e.GET("/users", users.GetUserHandler(service))

	// Call API
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// Start server
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "TODO next")
}

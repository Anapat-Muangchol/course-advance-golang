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

func TestGetUserHandlerSuccess(t *testing.T) {
	// Start server
	e := echo.New()
	service := users.NewUserService()
	e.GET("/users", users.GetUserHandler(service))

	rec := httptest.NewRecorder()
	httptest.NewRequest(http.MethodGet, "/users", nil)

	resCodeExpected := 200
	assert.Equal(t, resCodeExpected, rec.Code)
}

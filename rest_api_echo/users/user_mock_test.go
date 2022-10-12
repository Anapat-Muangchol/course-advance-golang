//go:build mock

package users_test

import (
	"api/users"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserRepository struct {
}

func (r MockUserRepository) GetSth() (string, error) {
	return "Mock data", fmt.Errorf("Mock error")
}

func TestSuccessWithGetWithRealServerAndMockRepository(t *testing.T) {
	// Setup router
	e := echo.New()
	repo := MockUserRepository{}
	service := users.NewUserService(repo)
	e.GET("/users", users.GetUserHandler(service))

	// Call API
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// Start server
	e.ServeHTTP(rec, req)

	// Assert
	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Mock data")
}

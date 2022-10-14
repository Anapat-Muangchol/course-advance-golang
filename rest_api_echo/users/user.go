package users

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
)

type UserResponse struct {
	Message string `json:"message"`
}

var tracer = otel.Tracer("demo-api")

func GetUserHandler(service *UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, span := tracer.Start(c.Request().Context(), "GetUserHandler", oteltrace.WithAttributes(attribute.String("layer", "handler")))
		defer span.End()
		res := UserResponse{
			Message: service.GetAll(),
		}
		return c.JSON(http.StatusOK, res)
	}
}

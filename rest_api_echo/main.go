package main

import (
	"api/users"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Command start server -> $go run main.go -port 8080
var port = flag.String("port", "1323", "Server Port")

func main() {
	fmt.Println(os.Getenv("GOPATH"))

	flag.Parse()

	// Echo server
	e := echo.New()

	// Add middlewares ...

	// Routers
	e.GET("/", homeHandler)
	e.GET("/users", users.UserHandler)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

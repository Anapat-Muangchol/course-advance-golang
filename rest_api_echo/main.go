package main

import (
	"api/internal"
	"api/users"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Command start server -> $go run main.go -port 8080
var port = flag.String("port", "1323", "Server Port")
var url = flag.String("url", "mongodb://root:password@localhost:27017", "Host Name MongoDb")

func main() {
	fmt.Println(os.Getenv("GOPATH"))

	flag.Parse()

	// Echo server
	e := echo.New()

	// Add global middlewares ...
	e.Use(middleware.Recover())

	// Routers
	e.GET("/", homeHandler)

	// User
	client := internal.NewMongoClient(*url)
	userRepository := users.NewUserRepository(client)
	userService := users.NewUserService(&userRepository)
	e.GET("/users", users.GetUserHandler(userService))

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

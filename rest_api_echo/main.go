package main

import (
	"api/internal"
	"api/users"
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"net/http"
	"os"
)

// Command start server -> $go run main.go -port 8080
var port = flag.String("port", "1323", "Server Port")
var url = flag.String("url", "mongodb://root:password@localhost:27017", "Host Name MongoDb")

func main() {

	// Init tracer
	tp, err := internal.InitTracer("http://zipkin:9411/api/v2/spans", "demo-api")
	if err != nil {
		log.Fatal(err)
	}
	defer func(tp *sdktrace.TracerProvider, ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(tp, context.Background())

	fmt.Println(os.Getenv("GOPATH"))
	flag.Parse()

	// Echo server
	e := echo.New()

	// Add global middlewares ...
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(middleware.Recover())
	e.Use(otelecho.Middleware("demo-api"))

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

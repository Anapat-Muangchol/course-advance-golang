# course-advance-golang-day-1

----

Run & Build command

$ go run cmd/main.go

$ go build -o demo cmd/main.go

---

Testing

$ go test ./...

$ go test demo

$ go test demo -v

$ go test demo -v -cover

https://pkg.go.dev/testing

---

Ref: lib & framework

https://github.com/go-chi/chi

https://github.com/gin-gonic/gin

https://echo.labstack.com/

https://docs.gofiber.io/

https://github.com/uber-go/zap

https://github.com/sirupsen/logrus

https://github.com/stretchr/testify

---

Field Alignment จัดเรียง Field ใน Struct เพื่อให้ใช้ memory น้อยลง

go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest

(GOPAHT)/fieldalignment -fix (file.go)

example.

fieldalignment -fix cmd/struct_reduce_resorce.go

---
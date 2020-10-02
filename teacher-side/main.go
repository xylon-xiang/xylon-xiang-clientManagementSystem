package main

import (
	"github.com/labstack/echo"
	"net/http"
)



func main() {
	e := echo.New()
	e.GET("/ws", helloWorld)

	e.Logger.Fatal(e.Start(":1234"))



}

func helloWorld(context echo.Context) error {

	return context.String(http.StatusOK, "Hello world")
}


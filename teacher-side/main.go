package main

import (
	"clientManagementSystem/config"
	"clientManagementSystem/teacher-side/constant"
	"clientManagementSystem/teacher-side/log"
	"clientManagementSystem/teacher-side/module"
	"clientManagementSystem/teacher-side/util"
	"github.com/labstack/echo"
	"net/http"
)

var Hub *util.Hub

func init() {
	Hub = util.NewHub()
	go Hub.Run()
}

func main() {
	e := echo.New()

	e.POST(config.Config.APIConfig.StudentLogAPI.URL, LoginController)

	e.Logger.Fatal(e.Start(":1234"))

}

func LoginController(context echo.Context) error {

	studentInfo := new(module.StudentInfo)
	if err := context.Bind(studentInfo); err != nil {
		return err
	}

	studentInfo.StudentId = context.Param("studentId")

	isMatch, err := log.CheckPassword(studentInfo.StudentId, studentInfo.StudentPassword)
	if err != nil {
		return err
	}

	if isMatch {
		host := context.Request().Host

		err = log.NewWebsocket(Hub, host)
		if err != nil {
			return err
		}

		return context.String(http.StatusOK, constant.ACCEPT)

	} else {
		return context.String(http.StatusUnauthorized, constant.FAILURE)
	}

}

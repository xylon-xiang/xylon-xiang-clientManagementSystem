package log

import (
	"bytes"
	"clientManagementSystem/config"
	"clientManagementSystem/module"
	"clientManagementSystem/student-side/util"
	"clientManagementSystem/teacher-side/constant"
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendLoginHttp(studentId string, studentLogPost module.StudentLogPost) (bool, error) {

	hostUrl := config.Config.APIConfig.TeacherHost +
		config.Config.APIConfig.StudentLogAPI.Path

	urlStr := util.GetRealUrl(hostUrl, studentId)

	logUrl, err := url.Parse(string(urlStr))
	if err != nil {
		return false, err
	}

	postBody := studentLogPost

	requestBody, err := json.Marshal(postBody)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(logUrl.String(), "application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// start an websocket connection
	if string(body) == constant.ACCEPT {
		return true, nil
	}

	return false, nil
}

func StartWebsocket(request *http.Request, responseWriter http.ResponseWriter) (*websocket.Conn, error) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	websocketConn, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		return nil, err
	}

	return websocketConn, nil

}

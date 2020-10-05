package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigObj struct {
	DatabaseConfig DataBaseConfig `json:"database_config"`
	APIConfig      APIConfig      `json:"api_config"`
}

//  Database config   //

type DataBaseConfig struct {
	MongoConfig MongoConfig `json:"mongo_config"`
}

type MongoConfig struct {
	DBAddress         string `json:"db_address"`
	DBName            string `json:"db_name"`
	TimeOut           int    `json:"time_out"`
	ClassCollection   string `json:"class_collection"`
	StudentCollection string `json:"student_collection"`
}

//   API interface    //

type APIConfig struct {
	TeacherHost            string              `json:"teacher_host"`
	WebsocketPort          string              `json:"websocket_port"`
	WebsocketCloseDuration int64               `json:"websocket_close_duration"`
	StudentLogAPI          StudentLogAPI       `json:"student_log_api"`
	StudentHandUpAPI StudentHandUpAPI `json:"student_hand_up_api"`
}

// student log in api for check password

type StudentLogAPI struct {
	Method             string             `json:"method"`
	Path               string             `json:"path"`
}


type StudentHandUpAPI struct {
	Method string	`json:"method"`
	Path string `json:"path"`
}




var Config ConfigObj

func init() {

	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	byteStream, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteStream, &Config)
	if err != nil {
		log.Fatal(err)
	}

}

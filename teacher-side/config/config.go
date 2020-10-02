package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigObj struct {
	DatabaseConfig DataBaseConfig `json:"database_config"`
}

type DataBaseConfig struct {
	MongoConfig MongoConfig `json:"mongo_config"`
}

type MongoConfig struct {
	DBAddress           string `json:"db_address"`
	DBName              string `json:"db_name"`
	TimeOut             int    `json:"time_out"`
	MusicInfoCollection string `json:"music_info_collection"`
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
	if err != nil{
		log.Fatal(err)
	}

}

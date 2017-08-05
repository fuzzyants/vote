package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Address    string `json:"address"`
	Dbuser     string `json:"dbuser"`
	Dbpassword string `json:"dbpassword"`
}

var config Configuration

func init() {
	parseConfig()
}

func parseConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("ERROR: Could not read config file.")
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
}

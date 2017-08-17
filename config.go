package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Address string `json:"address"`
}

var Config Configuration

func init() {
	parseConfig()
}

func parseConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("ERROR: Could not read config file.")
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
}

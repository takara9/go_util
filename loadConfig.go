package go_util

import (
        "encoding/json"
        "log"
        "os"
)

var Config Configuration
var logger *log.Logger

type Configuration struct {
	IpAddress    string
	TcpPort      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func init() {
	loadConfig()
	file, err := os.OpenFile("webapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
        logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
        decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

func addInt(a int, b int) int {
	return (a+b)
}




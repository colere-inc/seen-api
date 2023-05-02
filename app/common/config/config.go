package config

import (
	"log"
	"os"
)

var Port string

func Init() {
	log.Println("Port")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		log.Printf("defaulting to Port %s", Port)
	}
	log.Println(Port)
}

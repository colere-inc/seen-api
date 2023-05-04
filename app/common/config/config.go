package config

import (
	"fmt"
	"log"
	"os"
)

var Port string
var FreeeCompanyId string
var FreeeAccessToken string

func Init() {
	log.Println("Port")
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
		log.Printf("defaulting to Port %s", Port)
	}
	log.Println(Port)

	log.Println("FreeeCompanyId")
	FreeeCompanyId = getenv("FREEE_COMPANY_ID")
	log.Println(FreeeCompanyId)

	log.Println("FreeeAccessToken")
	FreeeAccessToken = os.Getenv("FREEE_ACCESS_TOKEN")
	if FreeeAccessToken == "" {
		panic("failed to get environment variable `FREEE_ACCESS_TOKEN")
	}
	log.Println("Success")
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("failed to get environment variable `%s", key))
	}
	return value
}

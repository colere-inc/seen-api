package config

import (
	"fmt"
	"os"
)

var Port string
var ProjectID string // GCP Project ID

func Init() {
	fmt.Println("Port")
	Port = getenv("PORT")
	fmt.Println(Port)

	fmt.Println("GCP Project ID")
	ProjectID = getenv("GCP_PROJECT_ID")
	fmt.Println(ProjectID)
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		errorString := fmt.Sprintf("failed to get environment variable %s", key)
		panic(any(errorString))
	}
	return value
}

package config

import (
	"fmt"
	"os"
)

var Port string

func Init() {
	fmt.Println("Port")
	Port = getenv("PORT")
	fmt.Println(Port)
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		errorString := fmt.Sprintf("failed to get environment variable %s", key)
		panic(any(errorString))
	}
	return value
}

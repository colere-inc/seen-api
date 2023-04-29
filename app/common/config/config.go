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
		panic(interface{}(fmt.Sprintf("failed to get environment variable `%s", key)))
	}
	return value
}

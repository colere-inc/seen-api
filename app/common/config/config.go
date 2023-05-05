package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// public constants
const (
	FreeeCompaniesCollectionId   = "freeeCompanies"
	FreeePartnersSubCollectionId = "partners"
)

// private constants
const freeeApiTokenPath = "/secrets/freee-api-token"

var Port string
var ProjectID string
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

	log.Println("ProjectID")
	ProjectID = getenv("GCP_PROJECT_ID")
	log.Println(ProjectID)

	log.Println("FreeeCompanyId")
	FreeeCompanyId = getenv("FREEE_COMPANY_ID")
	log.Println(FreeeCompanyId)

	log.Println("FreeeAccessToken")
	initFreeeAccessToken()
	log.Println("Success")
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("failed to get environment variable `%s", key))
	}
	return value
}

func initFreeeAccessToken() {
	// 環境変数に設定されている場合
	FreeeAccessToken = os.Getenv("FREEE_ACCESS_TOKEN")
	if FreeeAccessToken != "" {
		log.Println("loaded FreeeAccessToken from Environmental Variable (FREEE_ACCESS_TOKEN)")
		return
	}

	// ファイルが mount されている場合
	secretBytes, err := os.ReadFile(freeeApiTokenPath)
	if err != nil {
		panic(fmt.Sprintf("failed to read secret: %s", err))
	}
	var token freeeApiToken
	err = json.Unmarshal(secretBytes, &token)
	if err != nil {
		panic(err)
	}
	FreeeAccessToken = token.AccessToken
	msg := fmt.Sprintf("to load FreeeAccessToken from mounted secret file %s", freeeApiTokenPath)
	if FreeeAccessToken == "" {
		panic(fmt.Sprintf("failed %s: %s", msg, err))
	}
	log.Printf("success %s", msg)
}

type freeeApiToken struct {
	AccessToken string `json:"access_token"`
}

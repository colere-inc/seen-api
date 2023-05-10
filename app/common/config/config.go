package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

const FreeePartnersCollectionId = "partners"

const freeeApiTokenSecretName = "freee-api-token"

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
	log.Printf("%s***", FreeeAccessToken[:3])
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

	// Cloud Secret Manager から取得する場合
	secretBytes := getSecret()
	var token freeeApiToken
	err := json.Unmarshal(secretBytes, &token)
	if err != nil {
		panic(err)
	}
	FreeeAccessToken = token.AccessToken
	msg := fmt.Sprintf("to load FreeeAccessToken from Cloud Secret Manager (secret name: %s)", freeeApiTokenSecretName)
	if FreeeAccessToken == "" {
		panic(fmt.Sprintf("failed %s: %s", msg, err))
	}
	log.Printf("success %s", msg)
}

func getSecret() []byte {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)

	if err != nil {
		panic(fmt.Sprintf("failed to create secretmanager client: %v", err))
	}
	defer func(client *secretmanager.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	// Build the request.
	secretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", ProjectID, freeeApiTokenSecretName)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	return result.Payload.Data
}

type freeeApiToken struct {
	AccessToken string `json:"access_token"`
}

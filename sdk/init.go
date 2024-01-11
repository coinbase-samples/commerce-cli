package sdk

import (
	"log"
	"net/http"

	"github.com/coinbase-samples/commerce-sdk-go"
)

var Client *commerce.Client

const (
	COMMERCE_API_KEY = "COMMERCE_API_KEY"
)

func InitClient() {
	creds, err := commerce.ReadEnvCredentials(COMMERCE_API_KEY)
	if err != nil {
		log.Fatalf("error reading environmental variable: %s", err)
	}

	Client = commerce.NewClient(creds, http.Client{})
}

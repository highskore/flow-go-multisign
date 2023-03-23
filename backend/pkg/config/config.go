package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AdminAddress is the address of the TFC admin account.
var AdminAddress string

// AdminKeyIndex is the key index of the TFC admin account.
var AdminKeyIndex int

// GCPKmsResourceName is the resource name of the gcp kms account.
var GCPKmsResourceName string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AdminAddress = os.Getenv("ADMIN_ADDRESS")
	AdminKeyIndex, err = strconv.Atoi(os.Getenv("ADMIN_KEY_INDEX"))
	GCPKmsResourceName = os.Getenv("GCP_KMS_RESOURCE_NAME")
	if err != nil {
		log.Fatal("Error parsing ADMIN_KEY_INDEX")
	}
}

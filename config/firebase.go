package config

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func SetupFirebase() *auth.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	firebaseServiceAccountKeyFilePath := os.Getenv("FIREBASE_SERVICE_ACCOUNT_KEY_FILE_PATH")

	serviceAccountKeyFilePath, err := filepath.Abs(firebaseServiceAccountKeyFilePath)
	
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	
	if err != nil {
		panic("Firebase load error")
	}
	
	//Firebase Auth
	auth, err := app.Auth(context.Background())
	
	if err != nil {
		panic("Firebase load error")
	}
	
	return auth
}
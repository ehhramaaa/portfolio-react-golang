package main

import (
	"context"
	"log"
	"os"
	"portfolio/server/routes"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var app *fiber.App
var firebaseApp *firebase.App

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase.json")
	firebaseApp, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
}

func main() {
	PORT := os.Getenv("PORT")
	homepageEndpoint := os.Getenv("HOMEPAGE_ENDPOINT")
	homepagesCollectionName := os.Getenv("FIREBASE_HOMEPAGES_COLLECTION_NAME")
	homepagesDocumentID := os.Getenv("FIREBASE_HOMEPAGES_DOCUMENT_ID")

	servicepageEndpoint := os.Getenv("SERVICE_ENDPOINT")
	servicesCollectionName := os.Getenv("FIREBASE_SERVICES_COLLECTION_NAME")
	servicesDocumentID := os.Getenv("FIREBASE_SERVICES_DOCUMENT_ID")

	app = fiber.New()
	routes.Home(app, firebaseApp, &homepagesCollectionName, &homepagesDocumentID, &homepageEndpoint)
	routes.Service(app, firebaseApp, &servicesCollectionName, &servicesDocumentID, &servicepageEndpoint)
	log.Fatal(app.Listen(":" + PORT))
}

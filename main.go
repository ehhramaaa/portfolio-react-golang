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
	homePageEndpoint := os.Getenv("HOMEPAGE_ENDPOINT")
	homePagesCollectionName := os.Getenv("FIREBASE_HOMEPAGES_COLLECTION_NAME")
	homePagesDocumentID := os.Getenv("FIREBASE_HOMEPAGES_DOCUMENT_ID")

	servicePageEndpoint := os.Getenv("SERVICE_ENDPOINT")
	servicesCollectionName := os.Getenv("FIREBASE_SERVICES_COLLECTION_NAME")
	servicesDocumentID := os.Getenv("FIREBASE_SERVICES_DOCUMENT_ID")

	skillPageEndpoint := os.Getenv("SKILL_ENDPOINT")
	skillsCollectionName := os.Getenv("FIREBASE_SKILLS_COLLECTION_NAME")
	skillsDocumentID := os.Getenv("FIREBASE_SKILLS_DOCUMENT_ID")

	contactPageEndpoint := os.Getenv("CONTACT_ENDPOINT")
	contactsCollectionName := os.Getenv("FIREBASE_CONTACTS_COLLECTION_NAME")
	contactsDocumentID := os.Getenv("FIREBASE_CONTACTS_DOCUMENT_ID")

	projectPageEndpoint := os.Getenv("PROJECT_ENDPOINT")
	projectsCollectionName := os.Getenv("FIREBASE_PROJECTS_COLLECTION_NAME")

	app = fiber.New()
	routes.Home(app, firebaseApp, &homePagesCollectionName, &homePagesDocumentID, &homePageEndpoint)
	routes.Service(app, firebaseApp, &servicesCollectionName, &servicesDocumentID, &servicePageEndpoint)
	routes.Skill(app, firebaseApp, &skillsCollectionName, &skillsDocumentID, &skillPageEndpoint)
	routes.Contact(app, firebaseApp, &contactsCollectionName, &contactsDocumentID, &contactPageEndpoint)
	routes.Project(app, firebaseApp, &projectsCollectionName, &projectPageEndpoint)
	log.Fatal(app.Listen(":" + PORT))
}

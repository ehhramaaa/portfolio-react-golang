package controllers

import (
	"context"
	"log"
	"portfolio/server/models"
	"portfolio/server/utils"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx, firebaseApp *firebase.App, homepagesCollectionName *string, homepagesDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var home models.Home

	doc, err := client.Collection(*homepagesCollectionName).Doc(*homepagesDocumentID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := doc.DataTo(&home); err != nil {
		log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(home)
}

func PostHome(c *fiber.Ctx, firebaseApp *firebase.App, homepagesCollectionName *string, homepagesDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var home models.Home

	if err := c.BodyParser(&home); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	if home.Desc == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Home Desc Is Require"})
	}

	if home.Experience == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Home Experience Is Require"})
	}

	_, err = client.Collection(*homepagesCollectionName).Doc(*homepagesDocumentID).Set(ctx, home)
	if err != nil {
		log.Fatalf("Failed To Save Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(home)
}

func PatchHome(c *fiber.Ctx, firebaseApp *firebase.App, homepagesCollectionName *string, homepagesDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var updateData map[string]interface{}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	_, err = client.Collection(*homepagesCollectionName).Doc(*homepagesDocumentID).Update(ctx, utils.FirestoreUpdate(updateData))

	if err != nil {
		log.Fatalf("Failed To Update Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).SendString("Document Update Successfully")
}

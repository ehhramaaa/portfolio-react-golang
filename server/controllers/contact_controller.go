package controllers

import (
	"context"
	"log"
	"portfolio/server/models"
	"portfolio/server/utils"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx, firebaseApp *firebase.App, contactsCollectionName *string, contactsDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Contact models.Contact

	doc, err := client.Collection(*contactsCollectionName).Doc(*contactsDocumentID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := doc.DataTo(&Contact); err != nil {
		log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(Contact)
}

func PostContact(c *fiber.Ctx, firebaseApp *firebase.App, contactsCollectionName *string, contactsDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Contact models.Contact

	if err := c.BodyParser(&Contact); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	if Contact.Address == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Contact Address Is Require"})
	}

	if Contact.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Contact Email Is Require"})
	}

	if Contact.Phone == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Contact Phone Is Require"})
	}

	_, err = client.Collection(*contactsCollectionName).Doc(*contactsDocumentID).Set(ctx, Contact)
	if err != nil {
		log.Fatalf("Failed To Save Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(Contact)
}

func PatchContact(c *fiber.Ctx, firebaseApp *firebase.App, contactsCollectionName *string, contactsDocumentID *string) error {
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

	_, err = client.Collection(*contactsCollectionName).Doc(*contactsDocumentID).Update(ctx, utils.FirestoreUpdate(updateData))

	if err != nil {
		log.Fatalf("Failed To Update Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).SendString("Document Update Successfully")
}

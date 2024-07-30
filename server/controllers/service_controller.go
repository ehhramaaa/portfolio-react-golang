package controllers

import (
	"context"
	"log"
	"portfolio/server/models"
	"portfolio/server/utils"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func GetService(c *fiber.Ctx, firebaseApp *firebase.App, servicesCollectionName *string, servicesDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Service models.Service

	doc, err := client.Collection(*servicesCollectionName).Doc(*servicesDocumentID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := doc.DataTo(&Service); err != nil {
		log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(Service)
}

func PostService(c *fiber.Ctx, firebaseApp *firebase.App, servicesCollectionName *string, servicesDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Service models.Service

	if err := c.BodyParser(&Service); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	if Service.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Service Title Is Require"})
	}

	if Service.Desc == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Service Desc Is Require"})
	}

	if Service.Image == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Service Image Is Require"})
	}

	_, err = client.Collection(*servicesCollectionName).Doc(*servicesDocumentID).Set(ctx, Service)
	if err != nil {
		log.Fatalf("Failed To Save Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(Service)
}

func PatchService(c *fiber.Ctx, firebaseApp *firebase.App, servicesCollectionName *string, servicesDocumentID *string) error {
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

	_, err = client.Collection(*servicesCollectionName).Doc(*servicesDocumentID).Update(ctx, utils.FirestoreUpdate(updateData))

	if err != nil {
		log.Fatalf("Failed To Update Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).SendString("Document Update Successfully")
}

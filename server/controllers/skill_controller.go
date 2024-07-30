package controllers

import (
	"context"
	"log"
	"portfolio/server/models"
	"portfolio/server/utils"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func GetSkill(c *fiber.Ctx, firebaseApp *firebase.App, skillsCollectionName *string, skillsDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Skill models.Skill

	doc, err := client.Collection(*skillsCollectionName).Doc(*skillsDocumentID).Get(ctx)
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := doc.DataTo(&Skill); err != nil {
		log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(Skill)
}

func PostSkill(c *fiber.Ctx, firebaseApp *firebase.App, skillsCollectionName *string, skillsDocumentID *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var Skill models.Skill

	if err := c.BodyParser(&Skill); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	if Skill.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Skill Title Is Require"})
	}

	if Skill.Desc == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Skill Desc Is Require"})
	}

	if Skill.Cv == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Skill CV Is Require"})
	}

	_, err = client.Collection(*skillsCollectionName).Doc(*skillsDocumentID).Set(ctx, Skill)
	if err != nil {
		log.Fatalf("Failed To Save Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(Skill)
}

func PatchSkill(c *fiber.Ctx, firebaseApp *firebase.App, skillsCollectionName *string, skillsDocumentID *string) error {
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

	_, err = client.Collection(*skillsCollectionName).Doc(*skillsDocumentID).Update(ctx, utils.FirestoreUpdate(updateData))

	if err != nil {
		log.Fatalf("Failed To Update Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).SendString("Document Update Successfully")
}

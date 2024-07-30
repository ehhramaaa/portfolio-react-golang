package controllers

import (
	"context"
	"log"
	"portfolio/server/models"
	"portfolio/server/utils"

	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func GetProject(c *fiber.Ctx, firebaseApp *firebase.App, projectsCollectionName *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var projects []models.Project

	docs, err := client.Collection(*projectsCollectionName).Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	for _, doc := range docs {
		var project models.Project
		if err := doc.DataTo(&project); err != nil {
			log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		projects = append(projects, project)
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}

func GetProjectById(c *fiber.Ctx, firebaseApp *firebase.App, projectsCollectionName *string, projectDocumentId string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var project models.Project

	doc, err := client.Collection(*projectsCollectionName).Doc(projectDocumentId).Get(ctx)
	if err != nil {
		log.Fatalf("Failed To Get Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if err := doc.DataTo(&project); err != nil {
		log.Fatalf("Failed To Map Firestore Data To Struct : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

func PostProject(c *fiber.Ctx, firebaseApp *firebase.App, projectsCollectionName *string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	var project models.Project

	if err := c.BodyParser(&project); err != nil {
		log.Fatalf("Failed To Parse Request Body : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid Request Payload")
	}

	if project.Title == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Project Title Is Require"})
	}

	if project.Desc == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Project Desc Is Require"})
	}

	docRef, _, err := client.Collection(*projectsCollectionName).Add(ctx, project)
	if err != nil {
		log.Fatalf("Failed To Save Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	project.Id = docRef.ID
	project.Detail.ProjectId = docRef.ID

	_, err = docRef.Set(ctx, project)

	if err != nil {
		log.Fatalf("Failed To Add Project ID To ID And Detail Project : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}

func PatchProject(c *fiber.Ctx, firebaseApp *firebase.App, projectsCollectionName *string, projectDocumentId string) error {
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

	_, err = client.Collection(*projectsCollectionName).Doc(projectDocumentId).Update(ctx, utils.FirestoreUpdate(updateData))

	if err != nil {
		log.Fatalf("Failed To Update Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).SendString("Document Update Successfully")
}

func DeleteProject(c *fiber.Ctx, firebaseApp *firebase.App, projectCollectionName *string, projectDocumentId string) error {
	ctx := context.Background()
	client, err := firebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed To Create Firestore Client : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	defer client.Close()

	_, err = client.Collection(*projectCollectionName).Doc(projectDocumentId).Delete(ctx)

	if err != nil {
		log.Fatalf("Failed To Delete Document : %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Document Not Found")
	}

	return c.Status(fiber.StatusOK).SendString("Document Delete Successfully")
}

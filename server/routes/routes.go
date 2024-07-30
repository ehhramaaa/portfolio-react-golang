package routes

import (
	"portfolio/server/controllers"

	"firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
)

func Home(app *fiber.App, firebaseApp *firebase.App, homepagesCollectionName *string, homepagesDocumentID *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetHome(c, firebaseApp, homepagesCollectionName, homepagesDocumentID)
	})

	app.Post(*endpoint, func(c *fiber.Ctx) error {
		return controllers.PostHome(c, firebaseApp, homepagesCollectionName, homepagesDocumentID)
	})

	app.Patch(*endpoint, func(c *fiber.Ctx) error {
		return controllers.PatchHome(c, firebaseApp, homepagesCollectionName, homepagesDocumentID)
	})
}

func Service(app *fiber.App, firebaseApp *firebase.App, servicesCollectionName *string, servicesDocumentID *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})

	app.Post(*endpoint, func(c *fiber.Ctx) error {
		return controllers.PostService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})

	app.Patch(*endpoint, func(c *fiber.Ctx) error {
		return controllers.PatchService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})
}

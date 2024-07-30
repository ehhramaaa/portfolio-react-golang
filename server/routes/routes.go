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

	app.Post(*endpoint+"/add", func(c *fiber.Ctx) error {
		return controllers.PostHome(c, firebaseApp, homepagesCollectionName, homepagesDocumentID)
	})

	app.Patch(*endpoint+"/update", func(c *fiber.Ctx) error {
		return controllers.PatchHome(c, firebaseApp, homepagesCollectionName, homepagesDocumentID)
	})
}

func Service(app *fiber.App, firebaseApp *firebase.App, servicesCollectionName *string, servicesDocumentID *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})

	app.Post(*endpoint+"/add", func(c *fiber.Ctx) error {
		return controllers.PostService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})

	app.Patch(*endpoint+"/update", func(c *fiber.Ctx) error {
		return controllers.PatchService(c, firebaseApp, servicesCollectionName, servicesDocumentID)
	})
}

func Skill(app *fiber.App, firebaseApp *firebase.App, skillsCollectionName *string, skillsDocumentID *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetSkill(c, firebaseApp, skillsCollectionName, skillsDocumentID)
	})

	app.Post(*endpoint+"/add", func(c *fiber.Ctx) error {
		return controllers.PostSkill(c, firebaseApp, skillsCollectionName, skillsDocumentID)
	})

	app.Patch(*endpoint+"/update", func(c *fiber.Ctx) error {
		return controllers.PatchSkill(c, firebaseApp, skillsCollectionName, skillsDocumentID)
	})
}

func Contact(app *fiber.App, firebaseApp *firebase.App, contactsCollectionName *string, contactsDocumentID *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetContact(c, firebaseApp, contactsCollectionName, contactsDocumentID)
	})

	app.Post(*endpoint+"/add", func(c *fiber.Ctx) error {
		return controllers.PostContact(c, firebaseApp, contactsCollectionName, contactsDocumentID)
	})

	app.Patch(*endpoint+"/update", func(c *fiber.Ctx) error {
		return controllers.PatchContact(c, firebaseApp, contactsCollectionName, contactsDocumentID)
	})
}

func Project(app *fiber.App, firebaseApp *firebase.App, projectsCollectionName *string, endpoint *string) {
	app.Get(*endpoint, func(c *fiber.Ctx) error {
		return controllers.GetProject(c, firebaseApp, projectsCollectionName)
	})

	app.Get(*endpoint+"/:id", func(c *fiber.Ctx) error {
		return controllers.GetProjectById(c, firebaseApp, projectsCollectionName, c.Params("id"))
	})

	app.Post(*endpoint+"/add", func(c *fiber.Ctx) error {
		return controllers.PostProject(c, firebaseApp, projectsCollectionName)
	})

	app.Patch(*endpoint+"/update/:id", func(c *fiber.Ctx) error {
		return controllers.PatchProject(c, firebaseApp, projectsCollectionName, c.Params("id"))
	})

	app.Delete(*endpoint+"/delete/:id", func(c *fiber.Ctx) error {
		return controllers.DeleteProject(c, firebaseApp, projectsCollectionName, c.Params("id"))
	})
}

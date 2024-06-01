package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "gitlab.com/rodrigoodhin/gocure/api/docs"
	"gitlab.com/rodrigoodhin/gocure/api/routes"
)

// @title Gocure API
// @version 1.0
// @description Swagger for Gocure API
// @contact.name Gocure Support
// @contact.email rodrigo@odhin.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7087
// @BasePath /
func main() {
	// Setup a fiber service
	service := Setup()

	if err := service.Listen(":7087"); err != nil {
		log.Fatalln(err.Error())
	}
}

// Setup - setup a fiber service with all of its routes
func Setup() *fiber.App {
	// Init API Service
	app := fiber.New()

	// Default middleware config
	app.Use(requestid.New())

	// Add swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	//app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	//	URL:          "/swagger/doc.json",
	//	DeepLinking:  false,
	//	DocExpansion: "none",
	//}))

	// Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, X-Requested-With, Authorization, sentry-trace",
		AllowCredentials: true,
	}))

	// Set security headers
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")
		// Go to next middleware:
		return c.Next()
	})

	// Set index route
	app.Static("/", "api/public/index.html")

	// Create embed route group
	embed := app.Group("/embed")
	embed.Post("/toFeature", routes.EmbedToFeature)
	embed.Post("/toScenario", routes.EmbedToScenario)
	embed.Post("/toStep", routes.EmbedToStep)
	html := app.Group("/html")
	html.Post("/generate", routes.HTMLGenerate)

	// Return the configured service
	return app
}

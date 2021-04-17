package main

import (
	"github.com/avaldevilap/books_api/v1/pkg/configs"
	"github.com/avaldevilap/books_api/v1/pkg/middleware"
	"github.com/avaldevilap/books_api/v1/pkg/routes"
	"github.com/avaldevilap/books_api/v1/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/avaldevilap/books_api/v1/docs" // load API Docs files (Swagger)
	_ "github.com/joho/godotenv/autoload"        // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}

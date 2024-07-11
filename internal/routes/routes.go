package routes

import (
	"Doku/internal/app"
	"Doku/internal/handlers"
)

func InitializeRoutes(app *app.App) {
	api := app.Router.Group("/api")
	store := api.Group("/store")
	{
		store.POST("/create", handlers.CreateStore(app.DB))
		// store.DELETE()
	}

	admin := api.Group("/admin")
	{
		admin.POST("/create", handlers.CreateAdmin(app.DB))
		// admin.DELETE()
	}
}
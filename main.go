package main

import (
	"os"

	"github.com/iamtakdir/url-shortner-go/db"
	"github.com/iamtakdir/url-shortner-go/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	db.Connect()
	// Echo instance
	e := echo.New()

	// Routes
	routes.Routes(e)

	// Start server
	e.Start(":" + os.Getenv("PORT"))
	
}

// Handler

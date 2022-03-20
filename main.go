package main

import (
	"log"
	"net/http"
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
	if err := e.Start(":" + os.Getenv("PORT")); err != http.ErrServerClosed {
		log.Fatal(err)
	  }
	
}

// Handler

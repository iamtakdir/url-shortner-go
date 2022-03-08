package routes

import (
	"github.com/iamtakdir/url-shortner-go/controllers"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	// Routes
	e.POST("create", controllers.ShortUrl)
	e.GET("/:id", controllers.GetUrl)

}

package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/iamtakdir/url-shortner-go/db"
	"github.com/iamtakdir/url-shortner-go/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const link = "https://lfix-url.herokuapp.com"

func ShortUrl(c echo.Context) error {

	uuid := uuid.New().String()
	main_uuid := uuid[:5]

	var newReceiver models.ReceiverModel

	if err := c.Bind(&newReceiver); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "unprocessable"})
	}
	if newReceiver.Url == "" {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"error": "cant be empty"})
	}

	db.DB.Create(&models.ShortUrlModel{Url: newReceiver.Url, Uuid: main_uuid})

	url := fmt.Sprintf("%s/%s", link, main_uuid)

	return c.JSON(http.StatusOK, url)
}

func GetUrl(c echo.Context) error {
	//Shorten Url
	prams := c.Param("id")

	if prams == "" {
		c.JSON(http.StatusBadRequest, echo.Map{"error": "wrong url"})
	}

	var ShortenUrl models.ShortUrlModel

	result := db.DB.First(&ShortenUrl, "uuid = ?", prams)

	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		c.JSON(http.StatusBadRequest, echo.Map{"error": "Record not found"})
	} else {
		return c.JSON(http.StatusOK, echo.Map{"url": ShortenUrl.Url})
	}

	return nil

}

package router

import (
	"fmt"
	"net/http"

	"github.com/iresharma/REST1/internal/pkg/constants"
	"github.com/iresharma/REST1/internal/pkg/database"
	"github.com/iresharma/REST1/internal/pkg/models"
	"github.com/labstack/echo"
)

// CreateShortLink - creates the shorten link
func CreateShortLink(context echo.Context) error {
	link := models.CreateLinkModel{}
	if err := context.Bind(&link); err != nil {
		fmt.Println(err.Error())
		return context.JSON(http.StatusBadRequest, err)
	}
	ret, err := database.CreateLink(link)
	if err {
		return context.String(http.StatusOK, ret)
	}
	erRes := models.ResponseErrorModel{Error: "kuch toh ho gaya"}
	return context.JSON(http.StatusInternalServerError, erRes)
}

// GetShortenLink - route to redirect
func GetShortenLink(context echo.Context) error {
	finder := context.Param("route")
	seacrhQuery := models.SearchDb{Route: constants.Host + "/" + finder}
	get, err := database.GetLink(seacrhQuery)
	if !err {
		return context.String(http.StatusNotFound, "Not Found")
	}
	fmt.Printf("%s ----------", get)
	return context.Redirect(http.StatusMovedPermanently, get)
}

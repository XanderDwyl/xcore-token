package controllers

// GET: /

import (
	"log"

	"github.com/a-fis/xcore-token/app/helpers"
	"github.com/a-fis/xcore-token/app/models"
	"github.com/labstack/echo"
)

// WelcomeHandler ...
func WelcomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("WelcomeHandler")

		list, err := models.GetComments()
		dataFromEth := helpers.GetDataFromEthNetwork()

		if err != nil {
			log.Print(err)
		}
		data := map[string]interface{}{}
		data["list"] = list
		data["dataFromEth"] = dataFromEth

		log.Println("DATA FROM ETHER!!!")
		log.Println(dataFromEth)

		tmpl := "welcome/index"

		return RenderTemplate(c, tmpl, data)
	}
}

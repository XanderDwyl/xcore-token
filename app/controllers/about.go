package controllers

// GET: /about

import (
	"log"

	"github.com/labstack/echo"
)

// AboutHandler ...
func AboutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("AboutHandler")

		data := map[string]interface{}{}

		tmpl := "about/index"
		//if c.QueryParams().Get("amp") == "1" {
		//	tmpl = "about/index_amp"
		//}
		return RenderTemplate(c, tmpl, data)
	}
}

package controllers

// GET: /about

import (
	"log"
	"strconv"

	"github.com/a-fis/xcore-token/app/models"
	"github.com/labstack/echo"
)

// DataInBlockHandler ...
func DataInBlockHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("DataInBlockHandler")
		idstr := c.Param("idstr")
		log.Printf("idstr:%s", idstr)
		id, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			log.Print(err)
			return c.String(403, err.Error())
		}

		comment, err := models.GetCommentByID(id)
		if err != nil {
			log.Print(err)
			return c.String(403, err.Error())
		}
		data := map[string]interface{}{}
		data["hash"] = comment.Hash
		data["timestamp"] = comment.Utime
		return c.JSON(200, data)
	}
}

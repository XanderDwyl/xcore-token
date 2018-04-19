package controllers

// GET: /about

import (
	"log"
	"strconv"

	"github.com/a-fis/xcore-token/app/models"
	"github.com/labstack/echo"
)

// S3JSONHandler ...
func S3JSONHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("S3JSONHandler")
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
		j, err := comment.JSONString()
		if err != nil {
			log.Print(err)
			return c.String(403, err.Error())
		}
		return c.String(200, j)
	}
}

package controllers

// GET: /about

import (
	"log"
	"strconv"

	"github.com/a-fis/xcore-token/app/models"
	"github.com/labstack/echo"
)

// HashHandler ...
func HashHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("HashHandler")
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
		h := comment.HashString()
		return c.String(200, h)
	}
}

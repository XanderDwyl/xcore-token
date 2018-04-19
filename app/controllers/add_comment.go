package controllers

// GET: /about

import (
	"fmt"
	"log"

	"github.com/a-fis/xcore-token/app/models"
	"github.com/labstack/echo"
)

// AddCommentHandler ...
func AddCommentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("AddCommentHandler")

		// Save to database process
		title := c.Request().PostFormValue("title")
		body := c.Request().PostFormValue("body")
		log.Printf("body:%s", body)
		comment, err := models.CreateComment(title, body)

		if err != nil {
			log.Print(err)
			SetFlashError(c, err.Error())
		} else {
			SetFlashError(c, fmt.Sprintf("comment created, id=%d", comment.ID))
		}

		return Redirect(c, "/")
	}
}

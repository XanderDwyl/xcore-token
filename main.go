package main

import (
	"log"
	"net/http"

	//
	"github.com/a-fis/xcore-token/app/controllers"
	_ "github.com/a-fis/xcore-token/app/models"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// init ..
func init() {
	log.SetFlags(log.Lshortfile)

}
func main() {
	// Make an engine
	engine := echo.New()

	engine.Renderer = NewTemplate("app/views/")

	// Set up echo debug level
	engine.Debug = true

	// use actual files
	engine.Static("/assets", "app/data/assets")
	//engine.Use(middleware.Recover())

	//engine.Use(middlewares.SetupDisplayCurrency())
	//engine.Use(middlewares.I18n())
	//engine.Use(middlewares.ServerHeader())
	//engine.Use(middlewares.ServerHeader())
	engine.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf_token",
		CookiePath:  "/",
	}))
	//gzip
	engine.Use(middleware.Gzip())
	//session
	store := session.NewCookieStore([]byte("secret-key-kaskdfkjhkedknjjkelaasdfjkajkeasd"))
	store.MaxAge(86400)
	engine.Use(session.Sessions("GSESSION", store))

	engine.GET("/favicon.ico", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/assets/favicon.ico")
	})

	engine.GET("/", controllers.WelcomeHandler())
	engine.GET("/s3_json/:idstr", controllers.S3JSONHandler())
	engine.GET("/hash/:idstr", controllers.HashHandler())
	engine.GET("/data_in_block/:idstr", controllers.DataInBlockHandler())
	engine.POST("/add_comment", controllers.AddCommentHandler())
	engine.GET("/about", controllers.AboutHandler())
	// cool custom logging
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
	}))

	engine.Start(":3000")
	return

}

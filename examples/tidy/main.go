package main

import(
	"net/http"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/dilmorja/ovni/utils"
)

func main() {
	app := echo.New()

	app.GET("/", home)

	log.Fatal(app.Start(":9090"))
}

func home(c echo.Context) error {
	return c.HTML(http.StatusOK, utils.GenHTML(&utils.HTMLOptions{
		Lang: "en",
		Charset: "UTF-8",
		Title: "Tidy",
		CSS: "h1 { color: green; }",
		WASMP: "",
		GOII: "ovni_goii.js",
	}))
}
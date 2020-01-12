package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"net/http"
)


func setupRouter() *gin.Engine {
	// Bundle to use for application lifetime
	bundle := i18n.NewBundle(language.English)

	// Load translations into bundle during initialization
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.es.toml")
	bundle.MustLoadMessageFile("active.en.toml")

	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/index", func(c *gin.Context) {
		lang := c.Request.URL.Query().Get("lang")
		if lang == "" {// if exists is false
			lang = "en" // default language
		}
		localizer := i18n.NewLocalizer(bundle, lang)
		greetingDefault := &i18n.Message{
			ID:          "greeting",
			Description: "Greeting message",
		}
		greetingTranslated := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: greetingDefault,
		})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Greeting": greetingTranslated,
			"lang": lang,
		})
	})

	return r
}

func main() {


	r := setupRouter()
	r.Run(":8080")
}

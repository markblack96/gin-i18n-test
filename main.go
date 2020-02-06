package main

import (
	"github.com/gin-gonic/gin"
	"i18n-l10n"

	"net/http"
)


func setupRouter() *gin.Engine {
	// Bundle to use for application lifetime
	var t i18n_l10n.Translator
	_, err := t.LoadStrings([]string{"en", "es"}); if err != nil {
		println(err.Error())
	}
	println(t.Strings)
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html", "templates/sample-page.html")
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		c.Request.Context()
	})
	r.GET("/index", func(c *gin.Context) {
		lang := c.Request.URL.Query().Get("lang")
		if lang == "" {// if exists is false
			lang = "en" // default language
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			// "strings": (t.Strings)[lang].(map[string]interface{})["index"], // this works but is really ugly
			"strings": t.GetStringsForPage("index", lang),
			"lang": lang,
		})
	})
	r.GET("/sample-page", func(c *gin.Context) {
		lang := c.Request.URL.Query().Get("lang")
		if lang == "" {// if exists is false
			lang = "en" // default language
		}
		c.HTML(http.StatusOK, "sample-page.html", gin.H{
			"strings": t.GetStringsForPage("sample-page", lang), // (t.Strings)[lang].(map[string]interface{})["sample-page"],
			"lang": lang,
		})
	})
	return r
}

func main() {


	r := setupRouter()
	r.Run(":8080")
}

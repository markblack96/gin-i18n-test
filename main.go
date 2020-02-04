package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"net/http"
)

/*
I'm gonna put some stuff in here to automatically get strings from the template
strings := getStrings('index.html')
*/
func getStrings(template string, lang string, localizer *i18n.Localizer) (map[string]interface{}, error) {
	// use template and lang to grab the right translations
	// id := template + "." + lang
	var activeStrings map[string]interface{}
	_, err := toml.DecodeFile("active."+lang+".toml", &activeStrings)
	// handle potential errors gracefully
	if err != nil {
		println(err.Error())
		return activeStrings, err
	}
	for k, v := range activeStrings[template].(map[string]interface{}) {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
	return activeStrings, nil
	// we could just return the strings themselves
}


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
			ID:          "index.greeting",
			// Description: "Greeting message",
		}
		greetingTranslated := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: greetingDefault,
		})
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Greeting": greetingTranslated,
			"lang": lang,
		})
		_, _ = getStrings("index", lang, localizer)
	})
	r.GET("/sample-page", func(c *gin.Context) {

	})

	return r
}

func main() {


	r := setupRouter()
	r.Run(":8080")
}

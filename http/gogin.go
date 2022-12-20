package gogin

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	config "github.com/Minfenglin/owrt-ezsetup/config"
	local_pkg "github.com/Minfenglin/owrt-ezsetup/pkg"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS

func header(c *gin.Context) {
	c.HTML(http.StatusOK, "header.tmpl", gin.H{
		"Title": "owrt-ezsetup",
	})
}

func status(c *gin.Context) {
	c.HTML(http.StatusOK, "status.tmpl", gin.H{
		// Status
		"Status": local_pkg.Parser_Status(),
	})
}

// func wifi(c *gin.Context) {
// 	c.HTML(http.StatusOK, "wifi.tmpl", gin.H{
// 		// Status
// 		"Wifi": local_pkg.Parser_WiFi(),
// 	})
// }

func footer(c *gin.Context) {
	c.HTML(http.StatusOK, "footer.tmpl", gin.H{
		// Status
		"Name": "xxx",
	})
}

func Server_run() {
	server := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	server.SetHTMLTemplate(templ)

	// example: ./http/assets/css/org.css
	server.StaticFS("/http", http.FS(f))

	conf := config.Initial()
	server.GET("/", header, status, footer)
	log.Fatal(server.Run(":" + conf.Port))
}

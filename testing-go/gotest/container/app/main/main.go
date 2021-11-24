package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"showtable"
	"github.com/gin-gonic/gin"

)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("templates/*")
	

	basicAuth := gin.BasicAuth(gin.Accounts{
        "devopenspace": "tecracer#2021",
    })

	authorized := router.Group("/", basicAuth)

	authorized.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", map[string]interface{}{})
	})
	authorized.POST("/query", showtable.Query)

	
	router.Static("/js", "./assets/js")
	router.Static("/css", "./assets/css")
	
	authorized.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })


	router.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Article struct{
	Title string
	Content string

}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("/go/src/app/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"首页",
		})
	})
	news := &Article{
		Title: "intership get started!",
		Content: "easy html",
	}
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK,"news.html",gin.H{
			"title":"新闻页面",
			"news":news, 
		})
	})
	r.Run(":50052")
}
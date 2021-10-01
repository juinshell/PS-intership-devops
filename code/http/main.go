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
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
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
	r.Run("localhost:50052")
}
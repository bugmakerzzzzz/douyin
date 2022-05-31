package main


import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
)


func main() {
	databsae.D= databsae.Connect()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
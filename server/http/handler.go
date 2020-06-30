package http

import "github.com/gin-gonic/gin"

func tagAdd(c *gin.Context) {
	hello := "hello world"
	c.Writer.Write([]byte(hello))
}

func tagSearch(c *gin.Context) {
	hello := "hello world"
	c.Writer.Write([]byte(hello))
}

func entityLink(c *gin.Context) {

}

func entitySearch(c *gin.Context) {

}

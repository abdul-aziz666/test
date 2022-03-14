package main

import (
	"test/test1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/binary", handler.ChangeDecimalToBinary)
	route.GET("/decimal", handler.ChangeBinaryToDecimal)

	route.Run(":8085")
}

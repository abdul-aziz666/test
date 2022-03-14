package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ChangeDecimalToBinary ...
func ChangeDecimalToBinary(c *gin.Context) {

	query := c.Query("query")
	params, err := strconv.Atoi(query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "please check your params query",
		})
		return
	}

	result := convertDecimalToBinary(params)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    result,
	})
}

func convertDecimalToBinary(number int) int {
	var result, mod int
	count := 1
	for number != 0 {
		mod = number % 2
		number = number / 2
		result += mod * count
		count *= 10
	}
	return result
}

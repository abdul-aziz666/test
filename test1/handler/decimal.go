package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ChangeBinaryToDecimal ...
func ChangeBinaryToDecimal(c *gin.Context) {
	query := c.Query("query")
	params, err := strconv.Atoi(query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "please check your params query",
		})
		return
	}

	result := convertBinaryToDecimal(params)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    result,
	})
}
func convertBinaryToDecimal(number int) int {
	var result, mod int
	count := float64(0)
	for number != 0 {
		mod = number % 10
		result += mod * int(math.Pow(2, count))
		number = number / 10
		count++
	}
	return result
}

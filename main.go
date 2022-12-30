package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
			return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html", "templates/fib.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",nil)
	})

	router.GET("/fib", func(c *gin.Context) {
		idx := c.DefaultQuery("fib-idx", "0")
		i, err := strconv.Atoi(idx)
		if err != nil {
			c.JSON(http.StatusBadRequest, "")
		}
		value := FibonacciRecursion(i)
		c.HTML(http.StatusOK, "fib.html", gin.H{
			"idx": i,
			"value": value,
		})
	})
	router.Run(":8080")
}
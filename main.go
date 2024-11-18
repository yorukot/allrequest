package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/*path", func(c *gin.Context) {
		fmt.Printf("Method: %s, Path: %s\n", c.Request.Method, c.Request.URL.Path)

		fmt.Println("Headers:")
		for key, values := range c.Request.Header {
			fmt.Printf("%s: %v\n", key, values)
		}

		fmt.Println("Query Parameters:")
		for key, values := range c.Request.URL.Query() {
			fmt.Printf("%s: %v\n", key, values)
		}

		c.Request.ParseForm()
		fmt.Println("Form Data:")
		for key, values := range c.Request.Form {
			fmt.Printf("%s: %v\n", key, values)
		}

		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("Body: %s\n", string(body))

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/*path", func(c *gin.Context) {
		fmt.Printf("Method: %s, Path: %s\n", c.Request.Method, c.Request.URL.Path)

		fmt.Println("Headers:")
		for key, values := range c.Request.Header {
			fmt.Printf("%s: %v\n", key, values)
		}

		fmt.Println("Query Parameters:")
		for key, values := range c.Request.URL.Query() {
			fmt.Printf("%s: %v\n", key, values)
		}

		c.Request.ParseForm()
		fmt.Println("Form Data:")
		for key, values := range c.Request.Form {
			fmt.Printf("%s: %v\n", key, values)
		}

		body, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("Body: %s\n", string(body))

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

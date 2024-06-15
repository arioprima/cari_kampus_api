package main

import (
	"github.com/gin-gonic/gin"
)

type DataUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Skills []string `json:"skills,omitempty"`
	Data DataUser `json:"data"`
}

func main() {
	r := gin.Default();
	r.GET("/user", func(c *gin.Context) {
		response := Response{
			Status: "success",
			Message: "Succsessfully Get Data",
			// Skills: []string{"Golang", "Python", "Java"},
			Data: DataUser{
				Name: "John Doe",
				Email: "john@gmail.com",
			},
		}
		c.JSON(200, response)
	})
	r.Run(":8080")
}
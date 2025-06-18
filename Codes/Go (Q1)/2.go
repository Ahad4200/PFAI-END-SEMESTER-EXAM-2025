package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a Student struct
type Student struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
	GPA  float64 `json:"gpa"`
}

// Main function
func main() {
	// Initialize Gin router
	router := gin.Default()

	// Create GET /students endpoint
	router.GET("/students", func(c *gin.Context) {
		students := []Student{
			{Name: "Ahad", Age: 20, GPA: 4.0},
			{Name: "Sara", Age: 34, GPA: 3.8},
			{Name: "Basit", Age: 45, GPA: 3.9},
		}

		// Return JSON with status 200
		c.JSON(http.StatusOK, students)
	})

	// Start the server on port 8080
	router.Run(":8080")
}

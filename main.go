package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("hello world")
	app := gin.Default()

	todos := []Todo{}

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "hello world"})
	})

	fmt.Println("Before the post")

	// Create a Todo
	app.POST("/api/todos", func(c *gin.Context) {
		todo := &Todo{}

		if err := c.ShouldBindJSON(todo); err != nil {
			fmt.Println("JSON binding error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error() + "ldw"})
			return
		}

		if todo.Body == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Todo body is required"})
			return
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		c.JSON(http.StatusCreated, todo)

	})

	app.PATCH("/api/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				c.JSON(http.StatusAccepted, todos[i])
				return
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"Error": "No id has been found"})

	})

	app.Run(":4000")

}

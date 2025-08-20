package main

import (
	"fmt"
	"net/http"
	"slices"

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

	app.GET("/api/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
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

	app.DELETE("/api/todos/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = slices.Delete(todos, i, i+1)
				c.JSON(http.StatusAccepted, gin.H{"Success": fmt.Sprintf("Successfully deleted Id number: %d", i+1)})
				return

			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not Delete"})
	})

	app.Run(":4000")

}

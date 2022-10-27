package main

import (
	"jidarwish/todoApp/controllers"
	"jidarwish/todoApp/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDb()
}

func main() {

  r := gin.Default()


  // Endpoints for todolist
  r.GET("/list", controllers.GetAllTodoLists)
  r.GET("/list/:id", controllers.GetTodoList)
  r.POST("/list", controllers.PostNewTodoList)
  r.POST("/list/:id", controllers.ModiftExistingTodoList)
  r.DELETE("/list/:id", controllers.DeleteTodoList)
  
  // Endpoints for todo item
  r.Run()
}


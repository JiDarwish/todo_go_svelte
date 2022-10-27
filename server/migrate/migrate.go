package main

import (
 "jidarwish/todoApp/initializers"
 "jidarwish/todoApp/models"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDb()
}

func main() {
  initializers.DB.AutoMigrate(&models.Todo{})
  initializers.DB.AutoMigrate(&models.TodoList{})
}

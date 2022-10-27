package controllers

import (
	"fmt"
	"jidarwish/todoApp/initializers"
	"jidarwish/todoApp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllTodoLists(c *gin.Context) {

  var todoLists []models.TodoList
  result := initializers.DB.Find(&todoLists)

  if result.Error != nil {
    log.Fatal("Could not load all todo lists")
    c.Status(http.StatusBadRequest)
    return
  }


  c.JSON(http.StatusOK, gin.H {
    "data": todoLists,
  })
}

func GetTodoList(c *gin.Context) {
  id := c.Param("id")

  var todoList models.TodoList
  res := initializers.DB.Find(&todoList, id) 
  fmt.Println("Here")

  // Handle error
  if res.Error != nil {
    log.Fatal("Problem fetching todoList")
    c.JSON(http.StatusInternalServerError, "")
    return
  }
  // Handle record not found
  if res.RowsAffected < 1 {
    log.Println(fmt.Sprintf("Cannot find todolist with id %s", id))
    c.Status(http.StatusNotFound)
    return
  }
  
  c.JSON(http.StatusOK, gin.H{
    "data": todoList,
  })
}


func PostNewTodoList(c *gin.Context) {
  // Get data from the req body
  var body struct{
    Title         string
    Description   string
  }

  c.Bind(&body)


  // Create list in the database
  list := models.TodoList{Title: body.Title, Description: body.Description}
  result := initializers.DB.Create(&list)

  // There's an error!  
  if result.Error != nil {
    log.Fatal("Error creating the todolist")
    c.Status(http.StatusBadRequest)
    return
  }
  
  c.JSON(http.StatusOK, gin.H{
    "todoList": list, 
  })
}

func ModiftExistingTodoList(c *gin.Context) {
  // Body
  ID := c.Param("id")
  var body struct {
    Title         string
    Description   string
  }

  c.Bind(&body)

  var todoList models.TodoList
  res := initializers.DB.Model(&todoList).Where("ID = ?", ID).Updates(models.TodoList{Title: body.Title, Description: body.Description})

  if res.Error != nil {
    log.Fatal("Error updating model")
    c.Status(http.StatusBadRequest)
    return
  } 

  if res.RowsAffected < 1 {
    log.Println("Could not find todolist to edit")
    c.Status(http.StatusBadRequest)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "data": todoList,
  })
}


func DeleteTodoList(c *gin.Context) {
  ID := c.Param("id")

  var todoList models.TodoList
  res := initializers.DB.Where("Id = ?", ID).Delete(&todoList)

  if res.Error != nil {
    log.Fatal(fmt.Sprintf("Error Deleting the todolist with id %s", ID))
    c.Status(http.StatusBadRequest)
    return
  }
  
  
  if res.RowsAffected < 1 {
    log.Println(fmt.Sprintf("Could not delete todo list with id %s because it was not found", ID))
    c.Status(http.StatusBadRequest)
    return
  }

  c.JSON(http.StatusAccepted, gin.H{
    "data": ID,
  })
}

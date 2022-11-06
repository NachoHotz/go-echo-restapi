// Package routes ...
package routes

import (
  "github.com/labstack/echo"
  "github.com/nachohotz/go-echo-restapi/controllers"
)

// TaskRouter ...
func TaskRouter(e *echo.Echo) {
  tasksRouter := e.Group("/tasks")

  tasksRouter.GET("", controllers.GetTasks)
  tasksRouter.GET("/:id", controllers.GetTask)
  tasksRouter.POST("", controllers.CreateTask)
  tasksRouter.PUT("/:id", controllers.UpdateTask)
  tasksRouter.DELETE("/:id", controllers.DeleteTask)
}

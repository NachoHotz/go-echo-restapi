package routes

import (
	"github.com/labstack/echo"
	"github.com/nachohotz/go-echo-restapi/controllers"
	"github.com/nachohotz/go-echo-restapi/middlewares"
)

func TaskRouter(e *echo.Echo) {
  tasksRouter := e.Group("/tasks", middlewares.BasicMiddleware)

  tasksRouter.GET("", controllers.GetTasks)
  tasksRouter.GET("/:id", controllers.GetTask)
  tasksRouter.POST("", controllers.CreateTask)
  tasksRouter.PUT("/:id", controllers.UpdateTask)
  tasksRouter.DELETE("/:id", controllers.DeleteTask)
}

package routes

import (
  "github.com/labstack/echo"
  "github.com/nachohotz/go-echo-restapi/controllers"
)

func IndexRouter(e *echo.Echo) {
  e.GET("/", controllers.IndexHello)
}

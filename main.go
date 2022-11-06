// Package main
package main

import (
  "github.com/labstack/echo"
  "github.com/nachohotz/go-echo-restapi/routes"
)

func main() {
  e := echo.New()

  routes.IndexRouter(e)
  routes.TaskRouter(e)

  e.Start(":3001")
}

// Package controllers ...
package controllers

import "github.com/labstack/echo"

// IndexHello ...
func IndexHello(c echo.Context) error {
  return c.String(200, "Hello, World! from IndexHello")
}

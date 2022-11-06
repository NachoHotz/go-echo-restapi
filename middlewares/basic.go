package middlewares

import (
  "github.com/labstack/echo"
)

func BasicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    authToken := c.Request().Header.Get("Authorization")

    if authToken != "Bearer 123456" {
      return c.String(401, "No access token provided")
    }

    return next(c)
  }
}

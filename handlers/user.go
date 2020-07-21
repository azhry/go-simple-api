package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// JSON define type that represent JSON-like data type
type JSON map[string]string

// GetUser handler for getting specific user data by id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, JSON{
		"data": id,
	})
}

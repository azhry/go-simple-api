package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetUser handler for getting specific user data by id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"data": id,
	})
}

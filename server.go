package main

import (
	"go-simple-api/handlers"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const jwtSecretKey = "secret-dll"

func main() {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		token := jwt.New(jwt.SigningMethodHS256)

		// set claims
		claims := token.Claims.(jwt.MapClaims)
		// add any key value fields to the token
		claims["email"] = "arliansyah_azhary@yahoo.com"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// generate encoded token and send it as response.
		t, err := token.SignedString([]byte(jwtSecretKey))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	})

	user := e.Group("/users")
	{
		user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(jwtSecretKey),
		}))
		user.GET("/:id", handlers.GetUser)
	}

	donation := e.Group("/donations")
	{
		donation.GET("/list", handlers.GetDonations)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

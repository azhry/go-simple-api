package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"go-simple-api/models"

	"github.com/labstack/echo/v4"
)

// JSON define type that represent JSON-like data type
type JSON map[string][]models.Donation

// GetDonations handler for getting list of donations
func GetDonations(c echo.Context) error {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_donation")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("SELECT * FROM donasi")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	var donations []models.Donation
	for rows.Next() {
		var id, idYayasan, amount int
		var phoneNumber, category, createdAt, updatedAt string
		var deletedAt sql.NullString

		if err := rows.Scan(&id, &idYayasan, &amount, &phoneNumber, &category, &createdAt, &updatedAt, &deletedAt); err != nil {
			fmt.Println(err.Error())
		}

		var deletedAtValue string

		if deletedAt.Valid {
			deletedAtValue = deletedAt.String
		} else {
			deletedAtValue = "null"
		}

		donations = append(donations, models.Donation{
			ID:          id,
			IDYayasan:   idYayasan,
			Amount:      amount,
			PhoneNumber: phoneNumber,
			Category:    category,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   deletedAtValue,
		})
	}

	return c.JSON(http.StatusOK, JSON{
		"data": donations,
	})
}

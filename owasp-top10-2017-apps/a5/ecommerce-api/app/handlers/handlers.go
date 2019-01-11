package handlers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globocom/secDevLabs/owasp-top10-2017-apps/a5/ecommerce-api/app/db"
	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING\n")
}

func ErrorMessage(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Error finding this UserID."})
}

// GetTicket returns the userID ticket.
func GetTicket(c echo.Context) error {

	cookie, err := ReadCookie(c)

	if err != nil {
		return ErrorMessage(c)
	}

	type UserClaims struct {
		Name string `json:"name"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(cookie, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	claims, _ := token.Claims.(*UserClaims)

	if !token.Valid {
		return ErrorMessage(c)
	}

	id := c.Param("id")
	userDataQuery := map[string]interface{}{"userID": id, "username": claims.Name}
	userDataResult, err := db.GetUserData(userDataQuery)
	if err != nil {
		// could not find this user in MongoDB (or MongoDB err connection)
		return ErrorMessage(c)
	}
	msgTicket := fmt.Sprintf("Hey, %s! This is your ticket: %s\n", userDataResult.Username, userDataResult.Ticket)
	return c.String(http.StatusOK, msgTicket)
}

package api

import (
	"fmt"
	"golang.org/x/crypto/bcrypt" //"hasher"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	db "github.com/globocom/secDevLabs/owasp-top10-2017-apps/a3/snake-pro/app/db/mongo"
	//"github.com/globocom/secDevLabs/owasp-top10-2017-apps/a3/snake-pro/app/pass"
	"github.com/globocom/secDevLabs/owasp-top10-2017-apps/a3/snake-pro/app/types"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// HealthCheck is the heath check function.
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING!\n")
}

func Root(c echo.Context) error {
	return c.Redirect(302, "/login")
}

//PassHash hash password
func PassHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

//ChkPassHash Check if hashed password matches original
func ChkPassHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// WriteCookie writes a cookie into echo Context
func WriteCookie(c echo.Context, jwt string) error {
	cookie := new(http.Cookie)
	cookie.Name = "sessionIDsnake"
	cookie.Value = jwt
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "")
}

// ReadCookie reads a cookie from echo Context.
func ReadCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie("sessionIDsnake")
	if err != nil {
		return "", err
	}
	return cookie.Value, err
}

// Register registers a new user into MongoDB.
func Register(c echo.Context) error {

	userData := types.UserData{}
	err := c.Bind(&userData)
	if err != nil {
		// error binding JSON
		return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Invalid Input."})
	}

	if userData.Password != userData.RepeatPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Passwords do not match."})
	}

	newGUID1 := uuid.Must(uuid.NewRandom())
	userData.UserID = newGUID1.String()
	userData.HighestScore = 0
	userData.Password, _ = PassHash(userData.Password)

	err = db.RegisterUser(userData)
	if err != nil {
		// could not register this user into MongoDB (or MongoDB err connection)
		return c.JSON(http.StatusInternalServerError, map[string]string{"result": "error", "details": "Error user data2."})
	}

	msgUser := fmt.Sprintf("User %s created!", userData.Username)
	return c.String(http.StatusOK, msgUser)
}

// Login checks MongoDB if this user exists and then returns a JWT session cookie.
func Login(c echo.Context) error {

	loginAttempt := types.LoginAttempt{}
	err := c.Bind(&loginAttempt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Error login."})
	}
	// input validation missing! do it later!

	userDataQuery := map[string]interface{}{"username": loginAttempt.Username}
	userDataResult, err := db.GetUserData(userDataQuery)
	if err != nil {
		// could not find this user in MongoDB (or MongoDB err connection)
		return c.JSON(http.StatusForbidden, map[string]string{"result": "error", "details": "Error login."})
	}

	validPass := ChkPassHash(loginAttempt.Password, userDataResult.Password)
	if !validPass {
		// wrong password
		return c.JSON(http.StatusForbidden, map[string]string{"result": "error", "details": "Error login."})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userDataResult.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return err
	}

	err = WriteCookie(c, t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"result": "error", "details": "Error login5."})
	}
	c.Response().Header().Set("Content-type", "text/html")
	messageLogon := fmt.Sprintf("Hello, %s! Welcome to SnakePro", userDataResult.Username)
	// err = c.Redirect(http.StatusFound, "http://www.localhost:10003/game/ranking")
	return c.String(http.StatusOK, messageLogon)
}

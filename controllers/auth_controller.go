package controllers

import (
	"authentication/config"
	"authentication/models"
	"authentication/utils"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	_, err = config.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, hashed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Email already used"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered"})
}

func Login(c echo.Context) error {
	var input models.User
	var dbUser models.User

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	row := config.DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", input.Email)
	err := row.Scan(&dbUser.ID, &dbUser.Email, &dbUser.Password)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	if !utils.CheckPasswordHash(input.Password, dbUser.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	claims := jwt.MapClaims{}
	claims["user_id"] = dbUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Token generation failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenStr})
}

func Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID := claims["user_id"]

	return c.JSON(http.StatusOK, echo.Map{
		"user_id": userID,
	})
}

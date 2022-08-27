package handlers

import (
	"net/http"

	"brewing_support/app/repositories"
	"brewing_support/app/usecases"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

/*
Login return staff and jwt token
*/
func Login(c echo.Context) error {
	m := echo.Map{}
	c.Bind(&m)

	repo := repositories.NewStaffRepo()

	staff, token, err1 := usecases.Login(repo, m["id"].(string), m["password"].(string))

	if err1 != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	refreshToken, err2 := usecases.GenerateRefreshToken(staff.ID)
	if err2 != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.JSON(http.StatusOK, echo.Map{"staff": staff, "jwtToken": token, "refreshToken": refreshToken})
}

/*
Login return staff and jwt token
*/
func RefreshToken(c echo.Context) error {
	m := echo.Map{}
	c.Bind(&m)
	id := m["id"].(string)

	token, err1 := usecases.GenerateToken(id)

	if err1 != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	refreshToken, err2 := usecases.GenerateRefreshToken(id)
	if err2 != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.JSON(http.StatusOK, echo.Map{"jwtToken": token, "refreshToken": refreshToken})
}

func getAuthStaffID(c echo.Context) string {
	if c.Get("user") == nil {
		return ""
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	staffID := claims["staffID"].(string)
	return staffID
}

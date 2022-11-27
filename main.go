package main

import (
	"encoding/json"

	_ "github.com/dilaragorum/swagger-playground/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

var Users []User

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
func main() {
	e := echo.New()

	v1 := e.Group("/")
	{
		v1.GET("users", GetUsers)
		v1.POST("users", CreateUser)
		v1.DELETE("users", DeleteUsers)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

// GetUsers
// @Summary      Show users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  []User
// @Router       /users [get]
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, Users)
}

// CreateUser Create User
// @Summary      Create new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      201  {object}  string
// @Param		_			body	User	true	"A user which will be created"
// @Router       /users [post]
func CreateUser(c echo.Context) error {
	var u User
	err := json.NewDecoder(c.Request().Body).Decode(&u)
	if err != nil {
		return err
	}

	Users = append(Users, u)

	return c.String(http.StatusCreated, "User is created")
}

// DeleteUsers Erase All Users
// @Summary      Delete users
// @Tags         users
// @Produce      json
// @Success      204  {object}  string
// @Router       /users [delete]
func DeleteUsers(c echo.Context) error {
	Users = nil
	return c.String(http.StatusNoContent, "All users are successfully deleted")
}

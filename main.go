package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   int
	Age  int
	Name string
}

var users = []User{}

func getUsers(c echo.Context) error {

	return c.JSON(http.StatusOK, users)
}

func saveUser(c echo.Context) error {
	u := new(User)   // u := &User{}
	err := c.Bind(u) //body
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "json unmarshal error"})
	}
	// usecase, repository, db
	u.Id = len(users) + 1
	users = append(users, *u)
	//
	return c.JSON(http.StatusCreated, map[string]string{"message": "success"})

}

func main() {
	// log
	// graceful shutdown
	// middleware
	// struct project
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", getUsers)
	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":1323"))
}

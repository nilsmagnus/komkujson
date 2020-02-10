package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/typedresponse", typed)
	e.GET("/", fallback)

	e.Logger.Fatal(e.Start(":8080"))
}

func fallback(c echo.Context) error {
	response := map[string]string{"hei": "sann", "foo": "bar"}
	return c.JSON(200, response)
}

func typed(c echo.Context) error {

	type a struct {
		A string `json:"a"`
	}

	return c.JSON(200, a{"foo"})
}

func withJsonBody(c echo.Context) error {
	type User struct {
		Name string `json:"name"`
	}
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
}

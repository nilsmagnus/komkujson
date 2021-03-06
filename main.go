package main

import (
	"net/http"
	"sort"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.POST("/jsonbodyrequest", withJsonBody)
	e.GET("/typedresponse", typedresponse)
	e.GET("/", fallback)

	e.Logger.Fatal(e.Start(":8080"))

}

func fallback(c echo.Context) error {
	response := map[string]string{"hei": "sann", "foo": "bar"}
	return c.JSON(200, response)
}

func typedresponse(c echo.Context) error {

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
}

// sorting is not hard in go, use the closure
func sortingfoo() {
	toSort := "heiabdc"
	sort.Slice(toSort, func(i, j int) bool {
		return toSort[i] < toSort[j]
	})
}

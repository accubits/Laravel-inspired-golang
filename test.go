// package main

// import (
// 	"sync"
// 	"fmt"
// )

// type person struct {
// 	name string
// 	age int
// }

// func NewPerson(name string) *person{

// 	p := person{name:name}
// 	p.age = 42
// 	return &p
// }

// type singleton struct {
// 	name string
// }

// var instance *singleton
// var once sync.Once

// func GetInstance() *singleton {
//     once.Do(func() {
//          instance = &singleton{name: "shsin"}
//     })
//     return instance
// }
// func main() {

// 	// fmt.Println(person{"Bob", 20})
//     // fmt.Println(person{name: "Fred"})
//     // fmt.Println(&person{name: "Ann", age: 40})

// 	my:=NewPerson("shehin")
// 	fmt.Println(my)

// 	fmt.Println(&my)

// 	fmt.Println(GetInstance())
// }

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

type handler struct{}

// Most of the code is taken from the echo guide
// https://echo.labstack.com/cookbook/jwt
func (h *handler) login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println("test")
	// Check in your db if the user exists or not
	if username == "jon" && password == "password" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		// This is the information which frontend can use
		// The backend can also decode the token and get admin etc.
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Shehin"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		// The signing string should be secret (a generated UUID          works too)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	h := &handler{}
    e.POST("/login", h.login)
	e.Logger.Fatal(e.Start(":1323"))
}

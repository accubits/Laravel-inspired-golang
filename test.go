package main


import (
	"sync"
	"fmt"
)

type person struct {
	name string
	age int
}

func NewPerson(name string) *person{

	p := person{name:name}
	p.age = 42
	return &p
}


type singleton struct {
	name string
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
         instance = &singleton{name: "shsin"}
    })
    return instance
}
func main() {

	// fmt.Println(person{"Bob", 20})
    // fmt.Println(person{name: "Fred"})
    // fmt.Println(&person{name: "Ann", age: 40})

	my:=NewPerson("shehin")
	fmt.Println(my)


	fmt.Println(&my)

	fmt.Println(GetInstance())
}



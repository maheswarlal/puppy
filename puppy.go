package puppy

import (
	"fmt"

	"github.com/maheswarlal/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}
func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func Fromlal1() {
	fmt.Println("I am from version 1.1.0")
}

func Fromlal2() {
	fmt.Println("I am from version 1.2.0")
}

func Taglal4() {
	fmt.Println("tagging V1.4.0")
}
func Taglal5() {
	fmt.Println("tagging V1.5.0")
}
func Taglal8() {
	fmt.Println("tagging V1.8.0")
}

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

func taglal2() {
	fmt.Println("tagging V1.2.0")
}
func taglal30() {
	fmt.Println("tagging V1.3.0")
}

package main

import (
	"fmt"
	sr "goexcrs/stringrotation"
)

func main() {
	isRotation := sr.IsRotation("waterbottle", "bottlewater")
	fmt.Println(isRotation)
}

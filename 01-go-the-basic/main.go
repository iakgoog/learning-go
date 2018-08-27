// shortcut "pkgm"
// shortcut "fp" of "ff"
// shortcut "forr"
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World")
	log.Println("Hello from log")

	// variable declaration : 1
	var gopher string
	gopher = "iakgoog"
	fmt.Printf("Hello, %s.\n", gopher)

	// variable declaration : 2
	var name = "Sutthinart Khunvadhana"
	fmt.Printf("My name is %s.\n", name)

	// variable declaration : 3
	nickName := "iakgoog"
	fmt.Printf("My nick name is %s.\n", nickName)

	// flow
	fmt.Print("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" { // if len(fruit) == 0
		fmt.Println("meh! ")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "lemon":
		fmt.Println("🍋")
	default:
		fmt.Println("👅")
	}

	// flow 2
	fmt.Print("Input score (0-100): ")
	var score int
	fmt.Scanln(&score)

	if score < 50 {
		fmt.Println("F")
	} else if score < 60 {
		fmt.Println("D")
	} else if score < 70 {
		fmt.Println("C")
	} else if score < 80 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}

	// array
	var a [5]int // a := [3]int{1, 2, 3}
	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(len(a))

	for i := 0; i < len(a); i++ { // for i := range a
		fmt.Println(a[i])
	}
	// or
	for _, v := range a {
		fmt.Println(v)
	}

	var b [2][3]int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = j
		}
	}
	fmt.Println(b)
}

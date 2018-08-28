// shortcut "pkgm"
// shortcut "fp" of "ff"
// shortcut "forr"
package main

import (
	"fmt"
	"log"
)

func main() {
	// basicExample()
	// flowExample()
	// arrayExample()
	// mapExample()
	pointerExample()
}

func basicExample() {
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
}

func flowExample() {
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
}

func arrayExample() {
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
		for j := 0; j < len(b[i]); j++ { // ++j is prohibited in Golang
			b[i][j] = j
		}
	}
	fmt.Println(b)

	// array 2
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(c)
	fmt.Println(c[4:])

	// array 3
	d := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	e := d[:2]
	e[0] = 20
	fmt.Println(d)
	fmt.Println(e)

	// array
	f := make([]int, 5)
	f[0] = 10
	f[2] = 30
	f[3] = 40
	fmt.Println(f) // [10 0 30 40 0]
	f = append(f, 90)
	fmt.Println(f) // [10 0 30 40 0 90]
}

func mapExample() {
	// map
	g := make(map[string]string)
	g["hello"] = "gopher"
	g["name"] = "iakgoog"

	fmt.Println(g)
	fmt.Println(g["hello"])
	fmt.Println(g["x"]) // print nothing

	g["x"] = "sugar daddy" // comment this line to see different result
	// delete(g, "x")
	// check if g contain "x"
	h, ok := g["x"]
	fmt.Println(ok)
	fmt.Println(h)

	// x would be valid in if scope
	if x, ok := g["x"]; ok {
		fmt.Println(x)
	}

	// loop through map
	for key, value := range g {
		fmt.Println(key, ":", value)
	}

	// init map #2
	i := map[string]string{
		"hello": "gopher",
		"name":  "iakgoog",
	}

	_ = i
}

func pointerExample() {

}

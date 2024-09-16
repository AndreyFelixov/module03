package main

import (
	"01task/citd"
	"fmt"

	"github.com/huandu/xstrings"
)

func main() {
	fmt.Println(citd.Digit())
	fmt.Println(citd.City())
	fmt.Println(xstrings.Shuffle(citd.City()))
}

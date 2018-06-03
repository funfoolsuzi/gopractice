package main

import (
	"fmt"

	"github.com/funfoolsuzi/gopractice/place"
)

func main() {
	p1 := place.Place{Name: "Tacoma", Lat: 47.246840, Long: -122.441727}
	p2 := place.Place{Name: "Zhengzhou", Lat: 34.743419, Long: 113.679614}
	fmt.Println(p1.Distance(&p2))
}

package main

import (
	"fmt"
	"time"

	. "github.com/funfoolsuzi/gopractice/place"
)

func main() {
	p1 := Place{Name: "Tacoma", Coordinate: Coordinate{Lat: 47.246840, Long: -122.441727}}
	p2 := Place{Name: "Zhengzhou", Coordinate: Coordinate{Lat: 34.743419, Long: 113.679614}}
	fmt.Println(p1.DistanceTo(p2.Coordinate))

	ac1 := Aircraft{
		ID:    9,
		Model: "OEING",
		Speed: 300,
	}
	ap1 := Airport{
		Place: Place{
			Name:       "SeaTac",
			Coordinate: Coordinate{Lat: 47.438164914, Long: -122.28916551},
		},
		Aircrafts: []Aircraft{ac1},
	}
	ap2 := Airport{
		Place: Place{
			Name:       "LAX",
			Coordinate: Coordinate{Lat: 33.9416, Long: -118.4085},
		},
		Aircrafts: []Aircraft{},
	}
	err := ap1.ScheduleFlight(&(ap1.Aircrafts[0]), ap2, 0)
	fmt.Println(err)
	time.Sleep(time.Second * 8)
}

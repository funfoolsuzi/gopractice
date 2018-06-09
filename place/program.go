package place

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartProgram will start 'plcae' program
func StartProgram() {
	airports := []Airport{
		Airport{
			Place: Place{
				Name:       "SeaTac",
				Coordinate: Coordinate{Lat: 47.438164914, Long: -122.28916551},
			},
			Aircrafts: []Aircraft{
				Aircraft{
					ID:    9,
					Model: "OEING",
					Speed: 300,
				},
			},
		},
		Airport{
			Place: Place{
				Name:       "LAX",
				Coordinate: Coordinate{Lat: 33.9416, Long: -118.4085},
			},
			Aircrafts: []Aircraft{},
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	options := []struct {
		key     byte
		word    string
		handler func() bool
	}{
		{'q', "[q]uit", func() bool { return false }},
		{'l', "[l]ist", func() bool {
			fmt.Println(airports)
			return true
		}},
	}

	for {
		fmt.Print("Options: ")
		for _, v := range options {
			fmt.Print(v.word)
			fmt.Print(" ")
		}
		fmt.Println()

		if scanner.Scan() {
			input := scanner.Text()
			if len(input) < 1 {
				fmt.Println("Received empty string. Try again.")
				continue
			}
			char := strings.ToLower(input)[0]
			ifContinue := true
			for _, v := range options {
				if v.key == char {
					ifContinue = v.handler()
					break
				}
			}
			if !ifContinue {
				break
			}
		} else {
			if e := scanner.Err(); e != nil {
				fmt.Println(e.Error())
			}
		}
	}
}

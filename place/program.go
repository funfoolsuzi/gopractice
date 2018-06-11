package place

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			for i, ap := range airports {
				fmt.Printf("  %d - %s\n", i+1, ap.String())
				for j, ac := range ap.Aircrafts {
					fmt.Printf("     %d - %s\n", j+1, ac.String())
				}
			}
			return true
		}},
		{'f', "[f]ly", func() bool {
			var origin, dest *Airport
			var acID int
			for {
				fmt.Print("Origin airport name >>>")
				if scanner.Scan() {
					originName := scanner.Text()
					for i, ac := range airports {
						if ac.Name == originName {
							origin = &airports[i]
							break
						}
					}
					if origin != nil {
						break
					}
					fmt.Println("Invalid. Try again.")
				}
			}
			for {
				fmt.Print("Aircraft ID >>>")
				if scanner.Scan() {
					idstr := scanner.Text()
					if id, err := strconv.Atoi(idstr); err == nil {
						acID = id
						break
					}
					fmt.Println("Invalid ID. Try again.")
				}
			}
			for {
				fmt.Print("Destination airport name >>>")
				if scanner.Scan() {
					destName := scanner.Text()
					for i, ac := range airports {
						if ac.Name == destName {
							dest = &airports[i]
							break
						}
					}
					if dest != nil {
						break
					}
					fmt.Println("Invalid. Try again.")
				}
			}
			go origin.FlyAircraft(acID, dest)
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

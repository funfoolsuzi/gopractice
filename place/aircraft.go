package place

import (
	"fmt"
)

// AircraftStatus describes the status of an airplane
type AircraftStatus int

// All aircraft statuses
const (
	Landed      AircraftStatus = 0
	Flying      AircraftStatus = 1
	Maintenance AircraftStatus = 2
)

// Aircraft is just a type of transporation
type Aircraft struct {
	Coordinate
	ID     int
	Model  string
	Speed  int // km/hour
	Status AircraftStatus
}

// StatusStr will return the status in string format
func (ac *Aircraft) StatusStr() string {
	switch ac.Status {
	case Landed:
		return "Landed"
	case Flying:
		return "Flying"
	case Maintenance:
		return "Maintenance"
	default:
		return ""
	}
}

func (ac *Aircraft) String() string {
	return fmt.Sprintf("Aircraft %d(ID) - {Coords: %v, Model: %s, Speed: %d km/hour, Status: %s}", ac.ID, ac.Coordinate, ac.Model, ac.Speed, ac.StatusStr())
}

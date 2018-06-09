package place

import (
	"fmt"
	"time"
)

// Airport is a substruct of Place, which can hold airplanes.
type Airport struct {
	Place
	Aircrafts []Aircraft
}

// RemoveAircraftByID removes an aircraft
func (a *Airport) RemoveAircraftByID(id int) (ac Aircraft, err error) {
	index := -1

	for idx, ac := range a.Aircrafts {
		if ac.ID == id {
			index = idx
		}
	}

	if index == -1 {
		err = &NonExistAircraftError{AirportName: a.Name, AircraftID: id}
		return
	}

	ac = a.Aircrafts[index]

	a.Aircrafts = append(a.Aircrafts[:index], a.Aircrafts[index+1:]...)

	return
}

// AddAircraft will add an aircraft to the airport
func (a *Airport) AddAircraft(ac *Aircraft) error {
	// check duplicates
	for _, v := range a.Aircrafts {
		if v.ID == ac.ID {
			return &DuplicateAircraftError{AirportName: a.Name, AircraftID: v.ID}
		}
	}
	a.Aircrafts = append(a.Aircrafts, *ac)
	return nil
}

// FlyAircraft will schedule an certain aircraft parked in the airport and fly it after certain hours.
// Return whether the scheduling succeed.
func (a *Airport) FlyAircraft(acID int, dest *Airport) error {
	// remove the aircraft
	acft, err := a.RemoveAircraftByID(acID)
	if err != nil {
		return err
	}
	fmt.Printf("Aircraft %d has taken off to fly from %s to %s\n", acft.ID, a.Name, dest.Name)

	// set up variables
	dis := a.DistanceTo(dest.Coordinate)
	flightDur := dis / float64(acft.Speed)
	rateLat := (dest.Lat - a.Lat) / flightDur
	rateLong := (dest.Long - a.Long) / flightDur

	// reset acft status
	acft.Coordinate = a.Coordinate
	acft.Status = Flying

	for ; flightDur > 1; flightDur-- {
		time.Sleep(time.Second)
		acft.Lat += rateLat
		acft.Long += rateLong
		fmt.Printf("Aircraft %d has just moved to %f, %f.\n", acft.ID, acft.Lat, acft.Long)
	}
	time.Sleep(time.Duration(flightDur) * time.Second)
	fmt.Printf("Aircraft %d has just arrived %s(%v) from %s.\n", acft.ID, dest.Name, dest.Coordinate, a.Name)
	acft.Coordinate = dest.Coordinate
	acft.Status = Landed

	if err := dest.AddAircraft(&acft); err != nil {
		return err
	}

	return nil
}

// DuplicateAircraftError is an error emited by AddAircraft
type DuplicateAircraftError struct {
	AirportName string
	AircraftID  int
}

func (err *DuplicateAircraftError) Error() string {
	return fmt.Sprintf("Airport(%s) has a duplicate aircraft(ID:%d)", err.AirportName, err.AircraftID)
}

// NonExistAircraftError is an error emited by RemoveAircraftByID
type NonExistAircraftError struct {
	AirportName string
	AircraftID  int
}

func (err *NonExistAircraftError) Error() string {
	return fmt.Sprintf("Airport(%s) does not have aircraft(ID:%d)", err.AirportName, err.AircraftID)
}

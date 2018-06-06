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

// RemoveAircraftByIndex removes an aircraft
func (a *Airport) RemoveAircraftByIndex(index int) (ac Aircraft, err error) {
	if index > len(a.Aircrafts)-1 {
		err = fmt.Errorf("The index of the aircraft being removed is out of range") // add more error info
		return
	}

	ac = a.Aircrafts[index]

	a.Aircrafts = append(a.Aircrafts[:index], a.Aircrafts[index+1:]...)

	return
}

// ScheduleFlight will schedule an certain aircraft parked in the airport and fly it after certain hours.
// Return whether the scheduling succeed.
func (a *Airport) ScheduleFlight(ac *Aircraft, dest Airport, sec float64) error {
	// check if aircraft is parked in this airport
	idx := -1
	fmt.Printf("%p\n", ac)
	for i, v := range a.Aircrafts {
		fmt.Printf("%p\n", &v)
		if ac == &v {
			idx = i
		}
	}

	if idx == -1 {
		return fmt.Errorf("The aircraft being scheduled to fly is not in this airport(%s)", a.Name)
	}

	dis := a.DistanceTo(dest.Coordinate)

	flightDur := dis / float64(ac.Speed)
	rateLat := (dest.Lat - a.Lat) / flightDur
	rateLong := (dest.Long - a.Long) / flightDur

	fmt.Printf("It will take %f seconds for aircraft %d to fly from %s to %s", flightDur, ac.ID, a.Name, dest.Name)

	go func() error {
		time.Sleep(time.Duration(sec) * time.Second)
		acft, err := a.RemoveAircraftByIndex(idx)

		acft.Coordinate = a.Coordinate
		acft.Status = Flying

		if err != nil {
			return err
		}

		for ; flightDur > 1; flightDur-- {
			time.Sleep(time.Second)
			acft.Lat += rateLat
			acft.Long += rateLong
			fmt.Printf("Aircraft %d has just moved to %f, %f", acft.ID, acft.Lat, acft.Long)
		}

		time.Sleep(time.Duration(flightDur) * time.Second)
		acft.Coordinate = dest.Coordinate

		return nil
	}()

	return nil
}

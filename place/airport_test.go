package place_test

import (
	"testing"

	. "github.com/funfoolsuzi/gopractice/place"
	"github.com/stretchr/testify/assert"
)

func TestRemoveAircraftByID(t *testing.T) {
	ac1 := Aircraft{ID: 9}
	ap1 := Airport{
		Place: Place{
			Name:       "SeaTac",
			Coordinate: Coordinate{Lat: 47.438164914, Long: -122.28916551},
		},
		Aircrafts: []Aircraft{ac1},
	}

	cases := []struct {
		in     int
		expID  int
		expErr error
	}{
		{in: 8, expErr: &NonExistAircraftError{AirportName: ap1.Name, AircraftID: 8}},
		{in: 9, expID: 9},
	}
	for _, c := range cases {
		acRemoved, err := ap1.RemoveAircraftByID(c.in)
		assert.Equal(t, c.expID, acRemoved.ID)
		if c.expErr == nil {
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, c.expErr.Error())
		}
	}
}

func TestAddAircraft(t *testing.T) {
	ap1 := Airport{
		Place: Place{
			Name:       "LAX",
			Coordinate: Coordinate{Lat: 33.9416, Long: -118.4085},
		},
		Aircrafts: []Aircraft{},
	}
	ac1 := Aircraft{ID: 9}

	cases := []struct {
		in       Aircraft
		expCount int
		expErr   error
	}{
		{in: ac1, expCount: 1},
		{in: ac1, expCount: 1, expErr: &DuplicateAircraftError{AirportName: ap1.Name, AircraftID: ac1.ID}},
	}

	for _, c := range cases {
		err := ap1.AddAircraft(&ac1)
		assert.Equal(t, c.expCount, len(ap1.Aircrafts))
		if c.expErr == nil {
			assert.Nil(t, err)
		} else {
			assert.EqualError(t, err, c.expErr.Error())
		}
	}
}

func TestFlyAircraft(t *testing.T) {
	seatac := Airport{
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
	}
	lax := Airport{
		Place: Place{
			Name:       "LAX",
			Coordinate: Coordinate{Lat: 33.9416, Long: -118.4085},
		},
		Aircrafts: []Aircraft{},
	}
	cases := []struct {
		origAirport  *Airport
		aircraftID   int
		destAirport  *Airport
		expOrigCount int
		expDestCount int
		expErr       error
	}{
		{origAirport: &seatac, aircraftID: seatac.Aircrafts[0].ID, destAirport: &lax, expOrigCount: 0, expDestCount: 1},
	}

	for _, c := range cases {
		if err := c.origAirport.FlyAircraft(c.aircraftID, &lax); err == nil {
			// no error
			assert.Equal(t, c.expOrigCount, len(c.origAirport.Aircrafts))
			assert.Equal(t, c.expDestCount, len(c.destAirport.Aircrafts))
		} else {
			// an error
		}
	}
}

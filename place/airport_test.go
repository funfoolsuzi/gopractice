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

func TestAirportString(t *testing.T) {
	cases := []struct {
		airport Airport
		expect  string
	}{
		{airport: Airport{Place: Place{Coordinate: Coordinate{Lat: 34.5253, Long: 113.8462}, Name: "Xinzheng"}, Aircrafts: []Aircraft{
			Aircraft{ID: 3},
			Aircraft{ID: 4},
			Aircraft{ID: 5},
		}}, expect: "Airport Xinzheng at (34.5253, 113.8462) has 3 aircrafts"},
		{airport: Airport{Place: Place{Coordinate: Coordinate{Lat: 37.4602, Long: 126.4407}, Name: "Incheon"}, Aircrafts: []Aircraft{
			Aircraft{ID: 6},
			Aircraft{ID: 7},
			Aircraft{ID: 8},
			Aircraft{ID: 9},
			Aircraft{ID: 10},
		}}, expect: "Airport Incheon at (37.4602, 126.4407) has 5 aircrafts"},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, c.airport.String())
	}
}

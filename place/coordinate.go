package place

import (
	"math"
)

// Coordinate is the Latitude and Longitude of a certain location
type Coordinate struct {
	Lat  float64
	Long float64
}

// EarthRadius is the radius of the Earth
const EarthRadius int = 6371 // kilo meter

// DistanceTo calculates the distance between two coordinates.
// The result is in km.
func (p *Coordinate) DistanceTo(p2 Coordinate) float64 {
	ra := math.Pi / 180
	a := 0.5 - math.Cos((p2.Lat-p.Lat)*ra)/2 + math.Cos(p.Lat*ra)*math.Cos(p2.Lat*ra)*(1-math.Cos((p2.Long-p.Long)*ra))/2
	return float64(2*EarthRadius) * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km
}

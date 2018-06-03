package place

import (
	"fmt"
	"math"
)

// Place is a struct that describes a physical place, like city, site or building.
type Place struct {
	Name string
	Lat  float64
	Long float64
}

// String implements Stringer
func (p *Place) String() string {
	return fmt.Sprintf("A Place(%v) at %f, %f", p.Name, p.Lat, p.Long)
}

// DistanceTo calculates the distance between two places.
// The result is in km.
func (p *Place) DistanceTo(p2 Place) float64 {
	ra := math.Pi / 180
	a := 0.5 - math.Cos((p2.Lat-p.Lat)*ra)/2 + math.Cos(p.Lat*ra)*math.Cos(p2.Lat*ra)*(1-math.Cos((p2.Long-p.Long)*ra))/2
	return 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km
}

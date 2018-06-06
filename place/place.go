package place

import (
	"fmt"
)

// Place is a struct that describes a physical place, like city, site or building.
type Place struct {
	Coordinate
	Name string
}

// String implements Stringer
func (p *Place) String() string {
	return fmt.Sprintf("A Place(%v) at %f, %f", p.Name, p.Lat, p.Long)
}

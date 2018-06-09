package place_test

import (
	txng "testing"

	. "github.com/funfoolsuzi/gopractice/place"
	"github.com/stretchr/testify/assert"
)

var tacoma Place = Place{Name: "Tacoma", Coordinate: Coordinate{Lat: 47.2529, Long: -122.4443}}
var seattle Place = Place{Name: "Seattle", Coordinate: Coordinate{Lat: 47.6062, Long: -122.3321}}

func TestString(t *txng.T) {
	cases := []struct {
		in   Place
		want string
	}{
		{tacoma, "A Place(Tacoma) at 47.252900, -122.444300"},
		{seattle, "A Place(Seattle) at 47.606200, -122.332100"},
	}

	for _, c := range cases {
		got := c.in.String()
		assert.Equal(t, c.want, got)
	}
}

func TestDistanceTo(t *txng.T) {
	got := tacoma.DistanceTo(seattle.Coordinate)
	want := 40.1815530370966
	assert.Equal(t, want, got)
}

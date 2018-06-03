package place_test

import (
	txng "testing"

	. "github.com/funfoolsuzi/gopractice/place"
)

var tacoma Place = Place{Name: "Tacoma", Lat: 47.2529, Long: -122.4443}
var seattle Place = Place{Name: "Seattle", Lat: 47.6062, Long: -122.3321}

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
		if got != c.want {
			t.Errorf("Place.String()\n result: %v \n want: %v", got, c.want)
		}
	}
}

func TestDistanceTo(t *txng.T) {
	got := tacoma.DistanceTo(seattle)
	want := 40.1815530370966
	if got != want {
		t.Errorf("Place.DistanceTo(pl *Place)\n result:%v \n want: %v", got, want)
	}
}

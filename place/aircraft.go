package place

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

package contracts

type SatelliteDetail struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"Message"`
}

type Coordinate struct {
	X        float64
	Y        float64
	Distance float32
}

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var MapCoordinates = map[string]Coordinate{
	"kenobi":    {X: -500.0, Y: -200.0, Distance: 0.0},
	"skywalker": {X: 100.0, Y: -100.0, Distance: 0.0},
	"sato":      {X: 500.0, Y: 100.0, Distance: 0.0},
}

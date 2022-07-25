package contracts

type SatelliteDetail struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"Message"`
}

type Coordinate struct {
	X        float32
	Y        float32
	Distance float32
}

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

var MapCoordinates = map[string]Coordinate{
	"kenobi":    {X: -500.0, Y: -200.0, Distance: 0.0},
	"skywalker": {X: 100.0, Y: -100.0, Distance: 0.0},
	"sato":      {X: 500.0, Y: 100.0, Distance: 0.0},
}

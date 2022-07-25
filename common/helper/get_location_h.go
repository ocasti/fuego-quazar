package helper

import (
	"github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/savaki/trilateration"
)

func GetLocation(distances ...float32) (x, y float32, err error) {
	kenobiCoordinate := trilateration.Point{
		X: float64(contracts.MapCoordinates["kenobi"].X),
		Y: float64(contracts.MapCoordinates["kenobi"].Y),
		R: float64(distances[0]),
	}

	skywalkerCoordinate := trilateration.Point{
		X: float64(contracts.MapCoordinates["skywalker"].X),
		Y: float64(contracts.MapCoordinates["skywalker"].Y),
		R: float64(distances[1]),
	}

	satoCoordinate := trilateration.Point{
		X: float64(contracts.MapCoordinates["sato"].X),
		Y: float64(contracts.MapCoordinates["sato"].Y),
		R: float64(distances[2]),
	}

	position, err := trilateration.Solve(kenobiCoordinate, skywalkerCoordinate, satoCoordinate)

	if err != nil {
		return 0, 0, err
	}

	return float32(position[0].X), float32(position[0].Y), nil
}

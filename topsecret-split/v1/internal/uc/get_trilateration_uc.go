package uc

import (
	"encoding/json"
	"errors"
	cc "github.com/ocasti/fuego-quazar/common/contracts"
	commonUC "github.com/ocasti/fuego-quazar/common/helper"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
)

type GetTrilaterationUC struct {
	SatelliteRepository contracts.SatelliteRepository
}

func (g *GetTrilaterationUC) Handler() (float32, float32, string, error) {
	satellites, err := g.SatelliteRepository.Get()
	if err != nil {
		return 0, 0, "", nil
	}

	if len(satellites) < 3 {
		return 0, 0, "", errors.New("information is not enough")
	}

	spaceShip := map[string]cc.Coordinate{}
	var spaceShipMessages [][]string

	for _, satellite := range satellites {
		var message []string
		err := json.Unmarshal([]byte(satellite.Message), &message)
		if err != nil {
			return 0, 0, "", err
		}

		spaceShipMessages = append(spaceShipMessages, message)

		spaceShip[satellite.SatelliteName] = cc.Coordinate{
			Distance: satellite.Distance,
		}
	}

	x, y, err := commonUC.GetLocation(spaceShip["kenobi"].Distance, spaceShip["skywalker"].Distance, spaceShip["sato"].Distance)
	if err != nil {
		return 0, 0, "", err
	}

	message := commonUC.GetMessage(spaceShipMessages...)

	return x, y, message, nil
}

func NewGetTrilaterationUC(satelliteRepository contracts.SatelliteRepository) *GetTrilaterationUC {
	return &GetTrilaterationUC{
		SatelliteRepository: satelliteRepository,
	}
}

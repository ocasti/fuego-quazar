package contracts

import "github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"

type SatelliteRepository interface {
	Get() ([]model.Satellite, error)
	Save(satellite model.Satellite) error
}

type GetTrilaterationUC interface {
	Handler() (float32, float32, string, error)
}

type PostLocationSatelliteUC interface {
	Handler(body RequestBody) error
}

type RequestBody struct {
	SatelliteName string
	Distance      float32  `json:"distance"`
	Message       []string `json:"Message"`
}

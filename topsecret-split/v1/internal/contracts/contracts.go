package contracts

import (
	"github.com/ocasti/fuego-quazar/common/contracts"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
)

type SatelliteRepository interface {
	Get() ([]model.Satellite, error)
	Save(satellite model.Satellite) error
}

type GetTrilaterationUC interface {
	Handler() (model.Satellite, error)
}

type PostLocationSatelliteUC interface {
	Handler(body contracts.SatelliteDetail) error
}

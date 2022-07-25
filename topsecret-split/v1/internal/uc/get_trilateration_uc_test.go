package uc

import (
	"errors"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/model"
	"github.com/stretchr/testify/mock"
	"testing"
)

type satelliteRepositoryMock struct {
	mock.Mock
}

func (s *satelliteRepositoryMock) Get() ([]model.Satellite, error) {
	args := s.Called()
	return args.Get(0).([]model.Satellite), args.Error(1)
}

func (s *satelliteRepositoryMock) Save(satellite model.Satellite) error {
	return s.Called(satellite).Error(0)
}

func TestGetTrilaterationUC_Handler(t *testing.T) {
	type fields struct {
		SatelliteRepository *satelliteRepositoryMock
	}
	tests := []struct {
		name    string
		fields  fields
		mocker  func(f fields)
		want    float32
		want1   float32
		want2   string
		wantErr bool
	}{
		{
			name: "shouldReturnSuccess",
			fields: fields{
				SatelliteRepository: &satelliteRepositoryMock{},
			},
			mocker: func(f fields) {
				f.SatelliteRepository.On("Get").Return([]model.Satellite{
					{
						SatelliteName: "skywalker",
						Distance:      250,
						Message:       `["este","","","mensaje",""]`,
					}, {
						SatelliteName: "sato",
						Distance:      600,
						Message:       `["","","un","","secreto"]`,
					}, {
						SatelliteName: "kenobi",
						Distance:      600,
						Message:       `["","es","","mensaje",""]`,
					},
				}, nil).Once()
			},
			want:    70.3125,
			want1:   -284.375,
			want2:   "este es un mensaje secreto",
			wantErr: false,
		}, {
			name: "shouldReturnError",
			fields: fields{
				SatelliteRepository: &satelliteRepositoryMock{},
			},
			mocker: func(f fields) {
				f.SatelliteRepository.On("Get").Return([]model.Satellite{
					{
						SatelliteName: "skywalker",
						Distance:      1,
						Message:       `["este","","","mensaje",""]`,
					}, {
						SatelliteName: "sato",
						Distance:      250,
						Message:       `["","","un","","secreto"]`,
					}, {
						SatelliteName: "kenobi",
						Distance:      600,
						Message:       `["","es","","mensaje",""]`,
					},
				}, nil).Once()
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		}, {
			name: "shouldReturnErrorIncompleteSatellites",
			fields: fields{
				SatelliteRepository: &satelliteRepositoryMock{},
			},
			mocker: func(f fields) {
				f.SatelliteRepository.On("Get").Return([]model.Satellite{
					{
						SatelliteName: "skywalker",
						Distance:      600,
						Message:       `["este","","","mensaje",""]`,
					}, {
						SatelliteName: "sato",
						Distance:      250,
						Message:       `["","","un","","secreto"]`,
					},
				}, nil).Once()
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		}, {
			name: "shouldReturnErrorOnGet",
			fields: fields{
				SatelliteRepository: &satelliteRepositoryMock{},
			},
			mocker: func(f fields) {
				f.SatelliteRepository.On("Get").Return([]model.Satellite{}, errors.New("error")).Once()
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.fields)
			g := NewGetTrilaterationUC(tt.fields.SatelliteRepository)
			got, got1, msg, err := g.Handler()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Handler() got1 = %v, want %v", got1, tt.want1)
			}
			if msg != tt.want2 {
				t.Errorf("Handler() got1 = %v, want %v", got1, tt.want1)
			}
			tt.fields.SatelliteRepository.AssertExpectations(t)
		})
	}
}

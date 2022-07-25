package uc

import (
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"testing"
)

func TestPostLocationSatelliteUC_Handler(t *testing.T) {
	type fields struct {
		SatelliteRepository *satelliteRepositoryMock
	}
	type args struct {
		request contracts.RequestBody
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				SatelliteRepository: &satelliteRepositoryMock{},
			},
			args: args{
				request: contracts.RequestBody{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPostLocationSatelliteUC(tt.fields.SatelliteRepository)
			if err := p.Handler(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package helper

import "testing"

func TestGetLocation(t *testing.T) {
	type args struct {
		distances []float32
	}
	tests := []struct {
		name       string
		args       args
		wantX      float32
		wantY      float32
		wantErr    bool
		errMessage string
	}{
		{
			name: "shouldBeSuccessOnSendDistances",
			args: args{
				distances: []float32{
					600.0,
					250.0,
					600.0,
				},
			},
			wantX:   70.3125,
			wantY:   -284.375,
			wantErr: false,
		},
		{
			name: "shouldBeLocationNotFound",
			args: args{
				distances: []float32{
					100.0,
					100.0,
					100.0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := GetLocation(tt.args.distances...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("GetLocation() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetLocation() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

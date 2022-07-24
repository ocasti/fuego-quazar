package ctx

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_Handler(t *testing.T) {
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "shouldBeReturnSuccessData",
			args: args{
				req: events.APIGatewayProxyRequest{
					Body: `{"satellites": [{"name": "kenobi","distance": 600.0,"message": ["este", "", "", "mensaje", ""]},{"name": "skywalker","distance": 250.0,"message": ["", "es", "", "", "secreto"]},{"name": "sato","distance": 600.0,"message": ["este", "", "un", "", ""]}]}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       `{"position":{"x":70.3125,"y":-284.375},"message":"este es un mensaje secreto"}`,
			},
			wantErr: false,
		},
		{
			name: "shouldBeReturnError",
			args: args{
				req: events.APIGatewayProxyRequest{
					Body: `{"satellites": [{"name": "kenobi","distance": 100.0,"message": ["este", "", "", "mensaje", ""]},{"name": "skywalker","distance": 100.0,"message": ["", "es", "", "", "secreto"]},{"name": "sato","distance": 100.0,"message": ["este", "", "un", "", ""]}]}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler()
			got, err := h.Handler(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

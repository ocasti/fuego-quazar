package ctx

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/contracts"
	"github.com/stretchr/testify/mock"
)

type getTrilaterationUCMock struct {
	mock.Mock
}

func (g *getTrilaterationUCMock) Handler() (float32, float32, string, error) {
	args := g.Called()
	return args.Get(0).(float32), args.Get(1).(float32), args.Get(2).(string), args.Error(3)
}

type postLocationSatelliteUCMock struct {
	mock.Mock
}

func (p *postLocationSatelliteUCMock) Handler(body contracts.RequestBody) error {
	return p.Called(body).Error(0)
}

func TestHandler_Handler(t *testing.T) {
	type fields struct {
		GetTrilaterationUC      *getTrilaterationUCMock
		PostLocationSatelliteUC *postLocationSatelliteUCMock
	}
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		mocker  func(a args, f fields)
		wantErr bool
	}{
		{
			name: "shouldCreateSatelliteSuccessfully",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodPost,
					PathParameters: map[string]string{
						"satellite_name": "skywalker",
					},
					Body: `{"distance":120.5,"message":["","es","","","secreto"]}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
			},
			mocker: func(a args, f fields) {
				f.PostLocationSatelliteUC.On("Handler", mock.IsType(contracts.RequestBody{})).Return(nil).Once()
			},
			wantErr: false,
		}, {
			name: "shouldReturnErrorOnPost",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodPost,
					PathParameters: map[string]string{
						"satellite_name": "skywalker",
					},
					Body: `{"distance":120.5,"message":["","es","","","secreto"]}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"message":error}`,
			},
			mocker: func(a args, f fields) {
				f.PostLocationSatelliteUC.On("Handler", mock.IsType(contracts.RequestBody{})).Return(errors.New("error")).Once()
			},
			wantErr: false,
		}, {
			name: "shouldReturnErrorOnUnmarshall",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodPost,
					PathParameters: map[string]string{
						"satellite_name": "skywalker",
					},
					Body: `{"distance":120.5,"message":["","es","","","secreto",]}`,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       `{"message":invalid character ']' looking for beginning of value}`,
			},
			mocker: func(a args, f fields) {
			},
			wantErr: false,
		},
		{
			name: "shouldReturnDataSuccess",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodGet,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       `{"position":{"x":0,"y":0},"message":""}`,
			},
			mocker: func(a args, f fields) {
				f.GetTrilaterationUC.On("Handler").Return(float32(0.0), float32(0.0), "", nil).Once()
			},
			wantErr: false,
		}, {
			name: "shouldReturnErrorOnGet",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodGet,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       `{"message":error}`,
			},
			mocker: func(a args, f fields) {
				f.GetTrilaterationUC.On("Handler").Return(float32(0.0), float32(0.0), "", errors.New("error")).Once()
			},
			wantErr: false,
		}, {
			name: "shouldReturnErrorMethodNotAllowed",
			fields: fields{
				GetTrilaterationUC:      &getTrilaterationUCMock{},
				PostLocationSatelliteUC: &postLocationSatelliteUCMock{},
			},
			args: args{
				req: events.APIGatewayProxyRequest{
					HTTPMethod: http.MethodPut,
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: http.StatusMethodNotAllowed,
			},
			mocker: func(a args, f fields) {
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocker(tt.args, tt.fields)
			h := NewHandler(tt.fields.GetTrilaterationUC, tt.fields.PostLocationSatelliteUC)
			got, err := h.Handler(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() got = %v, want %v", got, tt.want)
			}
			tt.fields.PostLocationSatelliteUC.AssertExpectations(t)
			tt.fields.GetTrilaterationUC.AssertExpectations(t)
		})
	}
}

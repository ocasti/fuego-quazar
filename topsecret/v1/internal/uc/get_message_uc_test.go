package uc

import "testing"

func TestGetMessage(t *testing.T) {
	type args struct {
		messages [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantMsg string
	}{
		{
			name: "shouldSuccessOnSendMessages",
			args: args{
				messages: [][]string{
					{"", "este", "es", "un", "mensaje"},
					{"", "este", "es", "un", "mensaje"},
					{"", "", "es", "", "mensaje"},
				},
			},
			wantMsg: "este es un mensaje",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMsg := GetMessage(tt.args.messages...); gotMsg != tt.wantMsg {
				t.Errorf("GetMessage() = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}

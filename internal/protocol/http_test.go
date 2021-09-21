package protocol

import (
	"context"
	"testing"
)

func TestPingHTTP(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "baidu.com",
			args: args{
				addr: "https://baidu.com",
			},
			wantErr: false,
		},
		{
			name: "not.exist.domain.com",
			args: args{
				addr: "https://not.exist.domain.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PingHTTP(context.Background(), tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("PingHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package proto

import "testing"

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
			name: "http://baidu.com",
			args: args{
				addr: "http://baidu.com",
			},
			wantErr: false,
		},
		{
			name: "https://baidu.com",
			args: args{
				addr: "https://baidu.com",
			},
			wantErr: false,
		},
		{
			name: "https://not.exist.domain.com",
			args: args{
				addr: "https://not.exist.domain.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PingHTTP(tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("PingHTTP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

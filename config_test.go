package watchdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Interval uint
		Service  []Service
	}
	tests := []struct {
		name   string
		fields fields
		err    string
	}{
		{
			name: "missing global interval config",
			fields: fields{
				Service: make([]Service, 0),
			},
			err: "Config.Interval is required!",
		},
		{
			name: "missing service",
			fields: fields{
				Interval: 10,
			},
			err: "Config.Service is required!",
		},
		{
			name: "service property",
			fields: fields{
				Interval: 10,
				Service: []Service{
					{
						Name: "",
					},
				},
			},
			err: "Config.Service[0].Addr is required!\nConfig.Service[0].Interval must be greater than 0\nConfig.Service[0].Name is required!\nConfig.Service[0].Protocol is required!",
		},
		{
			name: "none error",
			fields: fields{
				Interval: 10,
				Service: []Service{
					{
						Name:     "test",
						Addr:     "localhost:22",
						Protocol: "ssh",
						Interval: 10,
					},
				},
			},
			err: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Interval: tt.fields.Interval,
				Service:  tt.fields.Service,
			}
			if tt.err == "" {
				assert.Nil(t, c.Validate(), tt.name)

			} else {
				assert.EqualError(t, c.Validate(), tt.err, tt.name)
			}
		})
	}
}

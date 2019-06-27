package color

import (
	"testing"
)

func TestColorMultiple(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    interface{}
		wantErr bool
	}{
		{"Blue Color", "blue", "#0000FF", false},
		{"Green Color", "green", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, err := ColorNew(tt.args)
			if (err != nil) != tt.wantErr || args != tt.want {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				//t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	return
}

package builtins_test

import (
	"bytes"
	"testing"

	"github.com/masonbesmer/CSCE4600/Project2/builtins"
)

func Test_Pwd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "pwd",
			args:    args{args: []string{}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// testing
			if err := builtins.Pwd(&out, tt.args.args...); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Pwd() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

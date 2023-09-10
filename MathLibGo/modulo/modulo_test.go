package modulo

import (
	"testing"
)

func TestModularAdder(t *testing.T) {
	type args struct {
		N       int
		addends []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "ModularAdder modulo 5",
			args:    args{5, []int{9, 7}},
			want:    1,
			wantErr: false,
		},
		{
			name:    "ModularAdder modulo 11",
			args:    args{11, []int{81, 38, 72, 64}},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0, nil},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adder, err := ModularAdder(tt.args.N)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModularAdder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sum := 0
			for _, n := range tt.args.addends {
				sum = adder(n)
			}
			if sum != tt.want {
				t.Errorf("ModularAdder() sum = %v, want %v", sum, tt.want)
				return
			}
		})
	}
}

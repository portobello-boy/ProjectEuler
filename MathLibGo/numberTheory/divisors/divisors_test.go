package divisors

import (
	"reflect"
	"testing"
)

func TestDivisorList(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Divisors of 10",
			args:    args{10},
			want:    []int{1, 2, 5, 10},
			wantErr: false,
		},
		{
			name:    "Divisors of 60",
			args:    args{60},
			want:    []int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60},
			wantErr: false,
		},
		{
			name:    "Divisors of 7",
			args:    args{7},
			want:    []int{1, 7},
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DivisorList(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("DivisorList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DivisorList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProperDivisorList(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Proper divisors of 10",
			args:    args{10},
			want:    []int{1, 2, 5},
			wantErr: false,
		},
		{
			name:    "Proper divisors of 60",
			args:    args{60},
			want:    []int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30},
			wantErr: false,
		},
		{
			name:    "Proper divisors of 7",
			args:    args{7},
			want:    []int{1},
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProperDivisorList(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProperDivisorList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProperDivisorList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimeDivisorMap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]int
		wantErr bool
	}{
		{
			name: "Prime divisors of 6",
			args: args{6},
			want: map[int]int{
				2: 1,
				3: 1,
			},
			wantErr: false,
		},
		{
			name: "Prime divisors of 8",
			args: args{8},
			want: map[int]int{
				2: 3,
			},
			wantErr: false,
		},
		{
			name: "Prime divisors of 11",
			args: args{11},
			want: map[int]int{
				11: 1,
			},
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PrimeDivisorMap(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrimeDivisorMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeDivisorMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

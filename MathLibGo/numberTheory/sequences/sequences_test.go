package sequences

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "n = 0",
			args:    args{0},
			want:    0,
			wantErr: false,
		},
		{
			name:    "n = 2",
			args:    args{2},
			want:    1,
			wantErr: false,
		},
		{
			name:    "n = 5",
			args:    args{5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "n = 10",
			args:    args{10},
			want:    55,
			wantErr: false,
		},
		{
			name:    "Invalid n",
			args:    args{-1},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fibonacci(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciSequence(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Valid case",
			args:    args{10},
			want:    []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			wantErr: false,
		},
		{
			name:    "Invalid length",
			args:    args{0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FibonacciSequence(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("FibonacciSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoundedFibonacciSequence(t *testing.T) {
	type args struct {
		bound int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "Bound 35",
			args:    args{35},
			want:    []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
			wantErr: false,
		},
		{
			name:    "Bound 34",
			args:    args{34},
			want:    []int{0, 1, 1, 2, 3, 5, 8, 13, 21},
			wantErr: false,
		},
		{
			name:    "Bound 0",
			args:    args{0},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "Invalid length",
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BoundedFibonacciSequence(tt.args.bound)
			if (err != nil) != tt.wantErr {
				t.Errorf("BoundedFibonacciSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundedFibonacciSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFibonacciClosure(t *testing.T) {
	tests := []struct {
		name        string
		invocations int
		result      int
	}{
		{
			name:        "Valid closure, invoked 1 time",
			invocations: 1,
			result:      1,
		},
		{
			name:        "Valid closure, invoked 2 times",
			invocations: 2,
			result:      2,
		},
		{
			name:        "Valid closure, invoked 10 times",
			invocations: 10,
			result:      89,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FibonacciClosure()
			for i := 0; i < tt.invocations-1; i++ {
				_ = got()
			}
			if result := got(); result != tt.result {
				t.Errorf("FibonacciClosure() = %v, want %v", result, tt.result)
			}
		})
	}
}

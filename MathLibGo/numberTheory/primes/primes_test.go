package primes

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_boundedPrimeSequenceProducer(t *testing.T) {
	type args struct {
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Primes below 10",
			args: args{10},
			want: []int{2, 3, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := 0
			for p := range boundedPrimeSequenceProducer(tt.args.bound) {
				if p != tt.want[index] {
					t.Errorf("boundedPrimeSequenceProducer() = %v, want %v", p, tt.want[index])
				}
				index += 1
			}
		})
	}
}

func TestBoundedPrimeSequence(t *testing.T) {
	type args struct {
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Primes below 10",
			args: args{10},
			want: []int{2, 3, 5, 7},
		},
		{
			name: "Primes below 100",
			args: args{100},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoundedPrimeSequence(tt.args.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundedPrimeSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_primeSequenceProducer(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First 10 primes",
			args: args{10},
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			name: "0 primes",
			args: args{0},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := 0
			for p := range primeSequenceProducer(tt.args.n) {
				fmt.Println(p)

				if p != tt.want[index] {
					t.Errorf("primeSequenceProducer() = %v, want %v", p, tt.want[index])
				}
				index += 1
			}
		})
	}
}

func TestPrimeSequence(t *testing.T) {
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
			name:    "First 11 primes",
			args:    args{11},
			want:    []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31},
			wantErr: false,
		},
		{
			name:    "0 primes",
			args:    args{0},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "Invalid bound",
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := PrimeSequence(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrimeSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, p := range p {
				if p != tt.want[i] {
					t.Errorf("PrimeSequence() = %v, want %v", p, tt.want[i])
				}
			}
		})
	}
}

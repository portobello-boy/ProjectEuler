package numberproperties

import (
	"testing"
)

func TestGcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Gcd of 6 and 8",
			args: args{6, 8},
			want: 2,
		},
		{
			name: "Gcd of 8 and 6",
			args: args{8, 6},
			want: 2,
		},
		{
			name: "Gcd of 8 and 0",
			args: args{8, 0},
			want: 8,
		},
		{
			name: "Gcd of 0 and 8",
			args: args{0, 8},
			want: 8,
		},
		{
			name: "Gcd of 16 and 8",
			args: args{16, 8},
			want: 8,
		},
		{
			name: "Gcd of 16 and 17",
			args: args{16, 17},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Lcm of 8 and 16",
			args: args{8, 16},
			want: 16,
		},
		{
			name: "Lcm of 5 and 7",
			args: args{5, 7},
			want: 35,
		},
		{
			name: "Lcm of 6 and 20",
			args: args{6, 20},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotient(t *testing.T) {
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
			name:    "Totient of 8",
			args:    args{8},
			want:    4,
			wantErr: false,
		},
		{
			name:    "Totient of 15",
			args:    args{15},
			want:    8,
			wantErr: false,
		},
		{
			name:    "Totient of 11",
			args:    args{11},
			want:    10,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{-1},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Totient(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Totient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Totient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPerfect(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "IsPerfect 6",
			args:    args{6},
			want:    true,
			wantErr: false,
		},
		{
			name:    "IsPerfect 8",
			args:    args{8},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPerfect(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPerfect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPerfect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAmicable(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "IsAmicable 220",
			args:    args{220},
			want:    true,
			wantErr: false,
		},
		{
			name:    "IsAmicable 240",
			args:    args{240},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAmicable(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAmicable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAmicable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAmicablePair(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "IsAmicablePair 220, 284",
			args:    args{220, 284},
			want:    true,
			wantErr: false,
		},
		{
			name:    "IsAmicablePair 220, 285",
			args:    args{220, 285},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0, 0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAmicablePair(tt.args.m, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAmicablePair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAmicablePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAbundant(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "IsAbundant 12",
			args:    args{12},
			want:    true,
			wantErr: false,
		},
		{
			name:    "IsAbundant 14",
			args:    args{14},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAbundant(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAbundant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAbundant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDeficient(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "IsDeficient 5",
			args:    args{5},
			want:    true,
			wantErr: false,
		},
		{
			name:    "IsDeficient 12",
			args:    args{12},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Invalid case",
			args:    args{0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDeficient(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDeficient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDeficient() = %v, want %v", got, tt.want)
			}
		})
	}
}

package day1

import "testing"

func Test_calculateRequiredFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example1", args{12}, 2},
		{"Example2", args{14}, 2},
		{"Example3", args{1969}, 654},
		{"Example4", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRequiredFuel(tt.args.mass); got != tt.want {
				t.Errorf("calculateRequiredFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

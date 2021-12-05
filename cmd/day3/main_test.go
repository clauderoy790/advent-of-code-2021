package main

import (
	"reflect"
	"testing"
)

func Test_isBitOne(t *testing.T) {
	type args struct {
		nb  int
		bit int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 5 101",
			args{5, 2},
			true,
		},
		{
			"test 5 101",
			args{5, 4},
			false,
		},
		{
			"test 5 101",
			args{246, 4},
			true,
		},
		{
			"test 5 101",
			args{2405, 11},
			true,
		},
		{
			"test 1 101",
			args{1, 0},
			true,
		},
		{
			"test 2 101",
			args{2, 0},
			false,
		},
		{
			"test 9 101",
			args{9, 0},
			true,
		},
		{
			"test 0 101",
			args{0, 0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBitOne(tt.args.nb, tt.args.bit); got != tt.want {
				t.Errorf("isBitOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNbsWithBit(t *testing.T) {
	type args struct {
		nbs       []int
		bit       int
		removeOne bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"testing 1",
			args{nbs: []int{}, bit: 4, removeOne: true},
			[]int{},
		},
		{
			"testing 1",
			args{nbs: []int{50, 250, 304, 101}, bit: 4, removeOne: true},
			[]int{101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNbsWithBit(tt.args.nbs, tt.args.bit, tt.args.removeOne); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNbsWithBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countZeroOne(t *testing.T) {
	type args struct {
		nbs []int
		bit int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"t1",
			args{nbs: []int{50, 250, 304, 101}, bit: 4},
			1,
			3,
		},
		{
			"t2",
			args{nbs: []int{10, 20, 22, 101, 4, 1, 20}, bit: 2},
			2,
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countZeroOne(tt.args.nbs, tt.args.bit)
			if got != tt.want {
				t.Errorf("countZeroOne() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countZeroOne() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findLeftBit(t *testing.T) {
	type args struct {
		nbs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{20, 500, 54051, 100, 131, 1, 50}},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLeftBit(tt.args.nbs); got != tt.want {
				t.Errorf("findLeftBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partOne()
		})
	}
}

func Test_findLifeSupportRating2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findLifeSupportRating2()
		})
	}
}

func Test_findLifeSupportRating(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			findLifeSupportRating()
		})
	}
}

func Test_findOxygenGenerator(t *testing.T) {
	type args struct {
		nbs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOxygenGenerator(tt.args.nbs); got != tt.want {
				t.Errorf("findOxygenGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCo2(t *testing.T) {
	type args struct {
		nbs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCo2(tt.args.nbs); got != tt.want {
				t.Errorf("findCo2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstBit(t *testing.T) {
	type args struct {
		nb int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstBit(tt.args.nb); got != tt.want {
				t.Errorf("findFirstBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		base int
		pow  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.base, tt.args.pow); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInput(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

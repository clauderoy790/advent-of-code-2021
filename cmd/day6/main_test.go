package main

import (
	"sync"
	"testing"
)

func Test_countAddedFrom(t *testing.T) {
	type args struct {
		startDay   int
		startValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"start at day 1 with val 0 returns",
			args{1, 0},
			6,
		},
		{
			"start at day 3 with val 0 returns",
			args{3, 0},
			4,
		},
		{
			"start at day 4 with val 0 returns",
			args{4, 0},
			3,
		},
		{
			"start at day 10 with val 8 returns",
			args{10, 8},
			0,
		},
		{
			"start at 2 adds on day 0",
			args{2, 0},
			4,
		},
		{
			"start at 3 adds on day 0",
			args{0, 3},
			4,
		},
		{
			"start on day 8 with val 5",
			args{8, 5},
			1,
		},
		{
			"start at 14 with value 8 returns",
			args{14, 8},
			0,
		},
		{
			"start at 1 1 returns",
			args{1, 1},
			4,
		},
		{
			"start at 16 8 returns",
			args{16, 8},
			0,
		},

		{
			"start at 9 with val 8 returns",
			args{9, 8},
			1,
		},
		{
			"start at 3 with val 8 returns",
			args{3, 8},
			1,
		},
		{
			"start at 9 with val 8 returns",
			args{9, 8},
			1,
		},
		{
			"start at 10 with val 8 returns",
			args{10, 8},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAddedFrom(tt.args.startDay, tt.args.startValue); got != tt.want {
				t.Errorf("countAddedFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_simulateM(t *testing.T) {
	type args struct {
		startDay int
		startVal int
		wg       *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulateM(tt.args.startDay, tt.args.startVal, tt.args.wg)
		})
	}
}

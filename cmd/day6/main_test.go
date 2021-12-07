package main

import "testing"

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
			"start at 3 adds on day 0",
			args{0,3},
			3,
		},
		{
			"start at 3 adds on day 0",
			args{0,3},
			3,
		},
		{
			"start at 8 on day 4 returns",
			args{0,3},
			3,
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

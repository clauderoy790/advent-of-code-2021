package main

import "testing"

func Test_removeFromString(t *testing.T) {
	type args struct {
		original string
		substr   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"remove cefbd from abcg",
			args{"cefbd","abcg"},
			"efd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeFromString(tt.args.original, tt.args.substr); got != tt.want {
				t.Errorf("removeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

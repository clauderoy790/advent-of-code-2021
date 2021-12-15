package main

import "testing"

func Test_insertRuneAt(t *testing.T) {
	type args struct {
		str string
		sub string
		pos int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"add at astart",
			args{"salut", "c", 0},
			"csalut",
		},
		{
			"add at end",
			args{"salut", "t", 4},
			"salutt",
		},
		{
			"add in middle",
			args{"clude", "a", 2},
			"claude",
		},
		{
			"add in serge",
			args{"sere", "g", 3},
			"serge",
		},
		{
			"add multiple",
			args{"c lettres", "'est beaucoup de", 1},
			"c'est beaucoup de lettres",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertStringAt(tt.args.str, tt.args.sub, tt.args.pos); got != tt.want {
				t.Errorf("insertRuneAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

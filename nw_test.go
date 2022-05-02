package main

import (
	"testing"
)

func TestNW(t *testing.T) {
	s1 := []string{"A", "A", "G"}
	s2 := []string{"A", "C", "G"}
	match := 1
	mismatch := -1
	gap := -2
	wantMatrix := [][]int{{0, -2, -4, -6}, {-2, 1, -1, -3}, {-4, -1, 0, -2}, {-6, -3, -2, 1}}
	wantDirection := [][][]int{{{-1, -1, -1, -1}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, {{0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, {{0, 0, 1, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, {{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}}}
	gotMatrix, gotDirection := nw(s1, s2, match, mismatch, gap)
	for i := range wantMatrix {
		for j := range wantMatrix[0] {
			if gotMatrix[i][j] != wantMatrix[i][j] {
				t.Errorf("excepted:%#v, got:%#v", wantMatrix, gotMatrix)
			}
		}
	}
	for i := range wantDirection {
		for j := range wantDirection[0] {
			for k := range wantDirection[0][0] {
				if gotDirection[i][j][k] != wantDirection[i][j][k] {
					t.Errorf("excepted:%#v, got:%#v", wantDirection, gotDirection)
				}
			}

		}
	}
}

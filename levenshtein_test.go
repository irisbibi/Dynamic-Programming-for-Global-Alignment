package main

import "testing"

func TestLevenshtein(t *testing.T) {
	s1 := []string{"A", "A", "G"}
	s2 := []string{"A", "C", "G"}
	wantMatrix := [][]int{{0, 1, 2, 3}, {1, 0, 1, 2}, {2, 1, 1, 2}, {3, 2, 2, 1}}
	wantDirection := [][][]int{{{-1, -1, -1, -1}, {0, 0, 1, 0}, {0, 0, 1, 0},{0,0,1,0}}, {{1, 0, 0, 0}, {0, 0, 1, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}}, {{1, 0, 0, 0}, {1, 0, 1, 0}, {0, 0, 1, 0}, {0, 1, 1, 0}}, {{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 1, 0}, {0, 0, 1, 0}}}
	wantScore := 1
	gotMatrix, gotDirection, gotScore := levenshtein(s1, s2)

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
	if gotScore != wantScore {

	}
}

package main

// nw takes two sequences and the scores of match, mismatch, and gap and returns the scoring matrix and distance matrix used for backtrace.
func nw(s, t []string, match, mismatch, gap int) ([][]int, [][][]int) {

	s1 := len(s)
	s2 := len(t)
	max := 0

	matrix := make([][]int, s1+1)

	for i := range matrix {
		matrix[i] = make([]int, s2+1)
	}

	max = s1
	if s2 > max {
		max = s2
	}

	direction := make([][][]int, s1+1)

	for i := range direction {
		direction[i] = make([][]int, s2+1)
		for j := range direction[i] {
			direction[i][j] = make([]int, max+1)
		}
		direction[0][0][i] = -1
	}

	for i := 1; i < s2+1; i++ {
		matrix[0][i] = gap * i
		direction[0][i][0] = 1
	}
	for j := 1; j < s1+1; j++ {
		matrix[j][0] = gap * j
		direction[j][0][2] = 1
	}
	
	corner := 0

	for a := 1; a < s1+1; a++ {
		for b := 1; b < s2+1; b++ {

			if s[a-1] == t[b-1] {
				corner = match
			} else {
				corner = mismatch
			}

			max := matrix[a-1][b-1] + corner
			left := matrix[a][b-1] + gap
			up := matrix[a-1][b] + gap

			matrix[a][b] = MaxThree(left, max, up)

			listofInt := []int{left, max, up}
			maxIndex := FindMaxIndex(matrix[a][b], listofInt)

			for i := range maxIndex {
				direction[a][b][maxIndex[i]] = 1
			}
		}
	}

	return matrix, direction
}


var resultList1, resultList2 [][]string

//backtraceNW traceback the scoring matrix using the direction matrix and output the optimal sequence alignments.
func backtraceNW(direction [][][]int, i, j int, s, t, finalSeq1, finalSeq2 []string, index int) {
	indicator1, indicator2 := false, false

	if i == 0 || j == 0 {
		if j != 0 {
			finalSeq1 = append(finalSeq1, t[i-1])
			finalSeq2 = append(finalSeq2, "-")
			backtraceNW(direction, i, j-1, s, t, finalSeq1, finalSeq2, index)
		}

		if i != 0 {
			finalSeq1 = append(finalSeq1, "-")
			finalSeq2 = append(finalSeq2, s[i-1])
			backtraceNW(direction, i-1, j, s, t, finalSeq1, finalSeq2, index)
		}
	}

	if i == 0 && j == 0 {
		resultList1 = append(resultList1, finalSeq1)
		resultList2 = append(resultList2, finalSeq2)
	}

	if i > 0 && j > 0 && direction[i][j][0] == 1 {
		newFinalSeq1 := append(finalSeq1, t[j-1])
		newFinalSeq2 := append(finalSeq2, "-")
		backtraceNW(direction, i, j-1, s, t, newFinalSeq1, newFinalSeq2, index)
		indicator1 = true
	}

	if i > 0 && j > 0 && direction[i][j][1] == 1 {
		newFinalSeq1 := append(finalSeq1, t[j-1])
		newFinalSeq2 := append(finalSeq2, s[i-1])

		if indicator1 == false {
			backtraceNW(direction, i-1, j-1, s, t, newFinalSeq1, newFinalSeq2, index)
		} else {
			backtraceNW(direction, i-1, j-1, s, t, newFinalSeq1, newFinalSeq2, index+1)
		}

		indicator2 = true
	}

	if i > 0 && j > 0 && direction[i][j][2] == 1 {
		newFinalSeq1 := append(finalSeq1, "-")
		newFinalSeq2 := append(finalSeq2, s[i-1])

		if indicator1 == false && indicator2 == false {
			backtraceNW(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index)
		} else if indicator1 == true && indicator2 == true {
			backtraceNW(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index+2)
		} else {
			backtraceNW(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index+1)
		}
	}

}

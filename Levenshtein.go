package main

// levenshtein takes two sequences to return the scoring and direction matrix 
// and alignment score between 2 sequences using Levinstein algorithm.
func levenshtein(s1, s2 []string) ([][]int, [][][]int, int) {
	a := len(s1)
	b := len(s2)
	max, score := 0, 0

	matrix := make([][]int, a+1)
	for i := range matrix {
		matrix[i] = make([]int, b+1)
	}

	max = a
	if b > max {
		max = b
	}

	direction := make([][][]int, a+1)
	for i := range direction {
		direction[i] = make([][]int, b+1)
		for j := range direction[i] {
			direction[i][j] = make([]int, max+1)
		}
		direction[0][0][i] = -1
	}

	// if either of the sequence is 0, return the length of the other sequence as the score
	if b == 0 && a != 0 {
		return matrix, direction, a

	} else if a == 0 && b != 0 {
		return matrix, direction, b
	}

	for i := 1; i < a+1; i++ {
		matrix[i][0] = i
		direction[i][0][0] = 1 // deletion
	}
	for j := 1; j < b+1; j++ {
		matrix[0][j] = j
		direction[0][j][2] = 1 // insertion
	}

	for i := 1; i < a+1; i++ {
		for j := 1; j < b+1; j++ {
			if s1[i-1] == s2[j-1] {
				score = 0

			} else {
				score = 1
			}
			insertion := matrix[i-1][j] + 1         // left
			deletion := matrix[i][j-1] + 1          // up
			subsitution := matrix[i-1][j-1] + score //corner

			matrix[i][j] = MinThree(insertion, deletion, subsitution)

			listofInt := []int{insertion, deletion, subsitution}
			minIndex := FindMinIndex(matrix[i][j], listofInt)
			
			for k := range minIndex {
				direction[i][j][minIndex[k]] = 1
			}
		}

	}
	return matrix, direction, matrix[a][b]
}


var insertion, deletion, substitution int

// backtraceLevenshtein tracebacks the scoring matrix and output the optimal sequence alignments.
func backtraceLevenshtein(direction [][][]int, i, j int, s, t, finalSeq1, finalSeq2 []string, index int) {
	indicator1, indicator2 := false, false

	if i == 0 || j == 0 {
		if j != 0 {
			finalSeq1 = append(finalSeq1, "-")
			finalSeq2 = append(finalSeq2, t[i-1])
			backtraceLevenshtein(direction, i, j-1, s, t, finalSeq1, finalSeq2, index)
		}

		if i != 0 {
			finalSeq1 = append(finalSeq1, s[i-1])
			finalSeq2 = append(finalSeq2, "-")
			backtraceLevenshtein(direction, i-1, j, s, t, finalSeq1, finalSeq2, index)
		}
	}

	if i == 0 && j == 0 {
		resultList1 = append(resultList1, finalSeq1)
		resultList2 = append(resultList2, finalSeq2)
	}

	if i > 0 && j > 0 && direction[i][j][1] == 1 {
		newFinalSeq1 := append(finalSeq1, "-")
		newFinalSeq2 := append(finalSeq2, t[j-1])
		backtraceLevenshtein(direction, i, j-1, s, t, newFinalSeq1, newFinalSeq2, index)
		indicator1 = true
		deletion++
	}

	if i > 0 && j > 0 && direction[i][j][2] == 1 {
		newFinalSeq1 := append(finalSeq1, s[i-1])
		newFinalSeq2 := append(finalSeq2, t[j-1])
		if indicator1 == false {
			backtraceLevenshtein(direction, i-1, j-1, s, t, newFinalSeq1, newFinalSeq2, index)
		} else {
			backtraceLevenshtein(direction, i-1, j-1, s, t, newFinalSeq1, newFinalSeq2, index+1)
		}

		indicator2 = true
		substitution++
	}

	if i > 0 && j > 0 && direction[i][j][0] == 1 {
		newFinalSeq1 := append(finalSeq1, s[i-1])
		newFinalSeq2 := append(finalSeq2, "-")
	
		if indicator1 == false && indicator2 == false {
			backtraceLevenshtein(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index)
		} else if indicator1 == true && indicator2 == true {
			backtraceLevenshtein(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index+2)
		} else {
			backtraceLevenshtein(direction, i-1, j, s, t, newFinalSeq1, newFinalSeq2, index+1)
		}
		insertion++
	}

}
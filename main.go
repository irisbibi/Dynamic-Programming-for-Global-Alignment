package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sequence1 := os.Args[1]
	sequence2 := os.Args[2]
	algorithm := os.Args[3]
	matrixFile := os.Args[4]
	resultFile := os.Args[5]

	string1, err := readFasta(sequence1)
	if err != nil {
		panic("cannot read fasta file")
	}

	string2, err := readFasta(sequence2)
	if err != nil {
		panic("cannot read fasta file")
	}

	s1 := strings.Split(string1, "")
	s2 := strings.Split(string2, "")

	fmt.Println("2 sequences are")
	fmt.Println(s1)
	fmt.Println(s2)

	var finalSeq1, finalSeq2 []string
	index := 0

	switch algorithm {

	case "NW", "nw", "needleman":
		scoreFile := os.Args[6]

		score := readScore(scoreFile)
		matchScore := score[0]
		mismatchScore := score[1]
		gapScore := score[2]

		matrix, direction := nw(s1, s2, matchScore, mismatchScore, gapScore)

		Writing(matrix, s1, s2, matrixFile)

		backtraceNW(direction, len(s1), len(s2), s1, s2, finalSeq1, finalSeq2, index)

		for i := range resultList1 {
			for i2, j := 0, len(resultList1[i])-1; i2 < j; i2, j = i2+1, j-1 {
				resultList1[i][i2], resultList1[i][j] = resultList1[i][j], resultList1[i][i2]
			}
		}
		for j := range resultList2 {
			for i, j2 := 0, len(resultList2[j])-1; i < j2; i, j2 = i+1, j2-1 {
				resultList2[j][i], resultList2[j][j2] = resultList2[j][j2], resultList2[j][i]
			}
		}

		fmt.Println("Score:", matrix[len(s1)-1][len(s2)-1])

		results(resultList1, resultList2, resultFile)

	case "l", "levenshtein", "leven":
		matrix, direction, score := levenshtein(s1, s2)

		Writing(matrix, s1, s2, matrixFile)
		
		backtraceLevenshtein(direction, len(s1), len(s2), s1, s2, finalSeq1, finalSeq2, index)

		fmt.Printf("Total insertion: %d\nTotal deletion: %d\nTotal substituion: %d\n", insertion, deletion, substitution)
		fmt.Println("Score:", score)

		for i := range resultList1 {
			for i2, j := 0, len(resultList1[i])-1; i2 < j; i2, j = i2+1, j-1 {
				resultList1[i][i2], resultList1[i][j] = resultList1[i][j], resultList1[i][i2]
			}
		}
		for j := range resultList2 {
			for i, j2 := 0, len(resultList2[j])-1; i < j2; i, j2 = i+1, j2-1 {
				resultList2[j][i], resultList2[j][j2] = resultList2[j][j2], resultList2[j][i]
			}
		}

		results(resultList1, resultList2, resultFile)

	default:
		panic("Please enter correct parameters!")
	}
}

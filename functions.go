package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MaxThree takes three integers and return the maximum integer of the three input integers.
func MaxThree(i, j, k int) int {
	if i > j {
		if i > k {
			return i
		}
		return k
	} else {
		if j > k {
			return j
		}
		return k
	}
}

// MinThree takes three integers and return the minimum integer of the three input integers.
func MinThree(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		}
		return k
	} else {
		if j < k {
			return j
		}
		return k
	}
}

// readScore takes a scoring file and returns the scores for different alignment.
func readScore(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot read the socring file!")
	}
	scanner := bufio.NewScanner(file)
	score := make([]int, 3)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		
		for i := range line {
			score[i], err = strconv.Atoi(line[i])
		}
	}
	return score
}

// readFasta returns the sequence from a FASTA file without comments and newlines.
func readFasta(filename string) (string, error) {
	buffer := bytes.NewBufferString("")
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" && line[0] != '>' { // discard comments
			buffer.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// Writing takes the scoring matrix, two sequences and the filename to save the scoring matrix to filename.
func Writing(d [][]int, l1, l2 []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Couldn't create the file!")
	}
	defer file.Close()

	line0 := "    " + strings.Join(l1, " ")
	fmt.Fprintln(file, line0)

	for i := range d {
		var line string
		if i == 0 {
			line = " "
		} else {
			line = l2[i-1]
		}

		for j := range d[i] {
			line = line + " " + strconv.Itoa(d[i][j])
		}
		fmt.Fprintln(file, line)
	}

}

// results save the optimal alignments into a file.
func results(resultList1, resultList2 [][]string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Couldn't create the file!")
	}
	defer file.Close()

	line0 := "The traceback results are:"
	fmt.Fprintln(file, line0)

	for i := range resultList2 {
		var line string
		line = strconv.Itoa(i + 1)
		fmt.Fprintln(file, line)
		line1 := resultList2[i]
		fmt.Fprintln(file, line1)
		line2 := resultList1[i]
		fmt.Fprintln(file, line2)
	}

}

// printDirection prints the direction matrix.
func printDirection(d [][][]int) {
	for i := range d {
		for j := range d[0] {
			fmt.Println("The current coord is (", i, ",", j, ")")
			fmt.Println("The direction list is ", d[i][j][:])
		}
	}
}

// FndMaxIndex takes a int and the list of the ints to return the position of the max int in the list.
func FindMaxIndex(max int, list []int) []int {
	index := []int{}
	for i := 0; i < len(list); i++ {
		if list[i] == max {
			index = append(index, i)
		}
	}
	return index
}

// FndMinIndex takes a int and the list of the ints to return the position of the min int in the list.
func FindMinIndex(min int, list []int) []int {
	index := []int{}
	for i := 0; i < len(list); i++ {
		if list[i] == min {
			index = append(index, i)
		}
	}
	return index
}

// traceback takes the scoring matrix and direction matrix and return the socring matrix with traceback routes.
// FUTURE STEPS!!
func traceback(matrix [][]int, direction [][][]int) [][]string {
	stringMatrix := make([][]string, len(matrix))
	for i := range stringMatrix {
		stringMatrix[i] = make([]string, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			stringMatrix[i][j] = strconv.Itoa(matrix[i][j])
		}
	}

	for i := 0; i < len(stringMatrix); i++ {
		newline := make([]string, 0)
		for j := 0; j < len(stringMatrix[i]); j++ {
			newline = append(newline, stringMatrix[i][j])
			newline = append(newline, " ")
			//stringMatrix[i][j] = " " + stringMatrix[i][j]
		}
		stringMatrix[i] = newline
	}

	new := make([][]string, 2*len(stringMatrix))
	for i := 0; i < len(new)-1; i = i + 2 {
		new[i] = stringMatrix[i/2]
		//fmt.Println(len(new[i]))
		new[i+1] = make([]string, len(new[i]))
	}
	new = new[:len(new)-1]

	// fmt.Println(stringMatrix)
	//fmt.Println(new)
	for i := 1; i < len(direction); i++ {
		for j := 1; j < len(direction[0]); j++ {
			if direction[i][j][1] == 1 {
				//new[i] = append(new[i][:j-2], new[i][j-1:]...)

				new[2*i-1][2*j] = "|"
				//new[i] = append(new[i][:j-1], temp[i][j-1:]...)
			}
			if direction[i][j][0] == 1 {
				new[2*i][2*j-1] = "-"
			}
			if direction[i][j][2] == 1 {
				new[2*i-1][2*j-1] = "\\"
			}
		}
	}

	return new
}


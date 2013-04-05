// By starting at the top of the triangle below and
// moving to adjacent numbers on the row below,
// the maximum total from top to bottom is 23.
// 
// 3
// 7 4
// 2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom in triangle.txt
// (right click and 'Save Link/Target As...'), a 15K text file containing
// a triangle with one-hundred rows.
//
// NOTE: This is a much more difficult version of Problem 18.
// It is not possible to try every route to solve this problem,
// as there are 299 altogether! If you could check one trillion (1012) routes every second
// it would take over twenty billion years to check them all.
// There is an efficient algorithm to solve it. ;o)
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const TriangleFile = "triangle.txt"

func TriangleToData(s []byte) ([][]int, error) {
	lines := bytes.Split(bytes.TrimSpace(s), []byte("\n"))
	data := make([][]int, len(lines))
	for i, l := range lines {
		fields := bytes.Fields(l)
		row := make([]int, len(fields))
		for col, field := range fields {
			n, err := strconv.Atoi(string(field))
			if err != nil {
				return nil, err
			}
			row[col] = n
		}
		data[i] = row
	}
	return data, nil
}

// LoadPreviousRow add n-1'th row data into nth row so that 
// maximize each field's total number.
func LoadPreviousRow(data [][]int, n int) {
	if n > 0 && n < len(data) {
		for i, _ := range data[n] {
			switch i {
			case 0:
				data[n][i] += data[n-1][0]
			case len(data[n]) - 1:
				data[n][i] += data[n-1][i-1]
			default:
				if data[n-1][i-1] > data[n-1][i] {
					data[n][i] += data[n-1][i-1]
				} else {
					data[n][i] += data[n-1][i]
				}
			}
		}
	}
}

// FindMaximumTotal returns largest sum of path through the triangle data.
func FindMaximumTotal(data [][]int) int {
	for i := 0; i < len(data); i++ {
		LoadPreviousRow(data, i)
	}
	max := 0
	for _, v := range data[len(data)-1] {
		if v > max {
			max = v
		}
	}
	return max
}

func Solver() {
	source, err := ioutil.ReadFile(TriangleFile)
	if err != nil {
		panic(err)
	}
	data, err := TriangleToData(source)
	if err != nil {
		panic(err)
	}
	fmt.Println(FindMaximumTotal(data))
}

func main() {
	Solver()
}

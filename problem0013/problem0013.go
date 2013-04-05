// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
//
// (data is in data.txt)
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

const DataFile = "data.txt"

func ReadAllNumbers(path string) ([]int64, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return BytesToNumbers(data)
}

func BytesToNumbers(b []byte) ([]int64, error) {
	fields := bytes.Fields(b)
	result := make([]int64, len(fields))
	for i, f := range fields {
		n, err := strconv.ParseInt(string(f), 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}

func Solver() {
	nums, err := ReadAllNumbers(DataFile)
	if err != nil {
		panic(err)
	}

	sum := int64(0)
	for _, n := range nums {
		sum += n
	}

	numStr := strconv.FormatInt(sum, 10)
	fmt.Println(numStr[:10])
}

func main() {
	Solver()
}

// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
//
// (data is in data.txt)
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
)

const DataFile = "data.txt"

func ReadAllNumbers(path string) ([]*big.Int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return BytesToNumbers(data)
}

func BytesToNumbers(b []byte) ([]*big.Int, error) {
	fields := bytes.Fields(b)
	result := make([]*big.Int, len(fields))
	for i, f := range fields {
		n := new(big.Int)
		_, err := fmt.Sscan(string(f), n)
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

	sum := big.NewInt(0)
	for _, n := range nums {
		sum.Add(sum, n)
	}

	numStr := sum.String()
	fmt.Println(numStr[:10])
}

func main() {
	Solver()
}

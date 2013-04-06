// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file
// containing over five-thousand first names, begin by sorting it into alphabetical order.
// Then working out the alphabetical value for each name, multiply this value by
// its alphabetical position in the list to obtain a name score.
//
// For example, when the list is sorted into alphabetical order, COLIN,
// which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
// So, COLIN would obtain a score of 938  53 = 49714.
//
// What is the total of all the name scores in the file?
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

const NamesFile = "names.txt"

func LoadNames(path string) ([]string, error) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := bytes.Split(src, []byte(","))
	names := make([]string, len(data))
	for i, d := range data {
		names[i] = string(bytes.Trim(d, `"`))
	}
	sort.Strings(names)
	return names, nil
}

// WordScore returns total score of name's alphabetical score.
// e.g.) COLIN : 3 + 15 + 12 + 9 + 14 = 53
func WordScore(name string) int {
	score := 0
	for _, c := range name {
		if c > 0 {
			score += int(c) - 64
		}
	}
	return score
}

func Solver() {
	names, err := LoadNames(NamesFile)
	if err != nil {
		panic(err)
	}
	total := 0
	for i, n := range names {
		total += (i + 1) * WordScore(n)
	}
	fmt.Println(total)
}

func main() {
	Solver()
}

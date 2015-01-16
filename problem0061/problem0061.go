package main

func triangle(n int) int {
	return (n*n + n) / 2
}

func square(n int) int {
	return n * n
}

func pentagonal(n int) int {
	return (3*n*n - n) / 2
}

func hexagonal(n int) int {
	return 2*n*n - n
}

func heptagonal(n int) int {
	return (5*n*n - 3*n) / 2
}

func octagonal(n int) int {
	return 3*n*n - 2*n
}

// findFourDigits lists up all four digits polygonal numbers.
func findFourDigits(f func(int) int) []int {
	ret := []int{}
	i := 0
	for {
		n := f(i)
		if n < 1000 {
			i++
			continue
		} else if n >= 10000 {
			return ret
		}
		ret = append(ret, n)
		i++
	}
}

// Candidate is cyclic polygonal candidate.
type Candidate struct {
	p int // polygonal number
	i int // index
}

// next finds cyclic polygonal candidates of c from polygonals.
func next(c Candidate, polygonals [][]int) []Candidate {
	ret := []Candidate{}
	for i := range polygonals {
		if i == c.i {
			continue
		}
		for _, n := range polygonals[i] {
			if c.p%100 == n/100 {
				ret = append(ret, Candidate{p: n, i: i})
			}
		}
	}
	return ret
}

func main() {
	polygonals := make([][]int, 6)
	for i := range polygonals {
		polygonals[i] = []int{}
	}

	polygonals[0] = findFourDigits(triangle)
	polygonals[1] = findFourDigits(square)
	polygonals[2] = findFourDigits(pentagonal)
	polygonals[3] = findFourDigits(hexagonal)
	polygonals[4] = findFourDigits(heptagonal)
	polygonals[5] = findFourDigits(octagonal)

	// TODO(ymotongpoo): implement cyclic number detection algorithm.
	for _, p := range polygonals[0] {
		cs := next(Candidate{p, 0}, polygonals)
		_ = cs
	}
}

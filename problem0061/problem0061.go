package main

import "fmt"

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

func main() {
	fmt.Println("hoge")
}

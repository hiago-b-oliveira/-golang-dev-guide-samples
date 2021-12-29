package main

import "fmt"

func main() {
	ns := createNumbers(10)

	for _, n := range ns {
		if n%2 == 0 {
			fmt.Printf("%+v is even\n", n)
		} else {
			fmt.Printf("%+v is odd\n", n)
		}
	}
}

func createNumbers(n int) []int {
	ns := make([]int, n+1)
	i := 0
	for {
		if i > n {
			break
		}
		ns[i] = i
		i = i + 1
	}
	return ns
}

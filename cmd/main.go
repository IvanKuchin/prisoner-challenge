package main

import "fmt"

// returns all permutations of numbers 0..n-1
// example: n=3 -> [0,1,2], [0,2,1], [1,0,2], [1,2,0], [2,0,1], [2,1,0]
func permutations(n int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		permute(n, []int{}, ch)
	}()
	return ch
}

func permute(n int, current []int, ch chan<- []int) {
	if len(current) == n {
		ch <- append([]int(nil), current...)
		return
	}
	for i := 0; i < n; i++ {
		if contains(current, i) {
			continue
		}
		permute(n, append(current, i), ch)
	}
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func longest_cycle(permutation []int) int {
	visited := make([]bool, len(permutation))
	var maxCycleLength int

	for i := 0; i < len(permutation); i++ {
		if !visited[i] {
			cycleLength := 0
			for j := i; !visited[j]; j = permutation[j] {
				visited[j] = true
				cycleLength++
			}
			if cycleLength > maxCycleLength {
				maxCycleLength = cycleLength
			}
		}
	}
	return maxCycleLength
}

func main() {
	crates := 4
	total_permutations := 0
	set_free := 0
	for i := range permutations(crates) {
		total_permutations++
		cycle_len := longest_cycle(i)

		if cycle_len <= crates/2 {
			set_free++
			print(i, " -> ")
			for _, v := range i {
				print(v)
			}

			print(" : longest cycle = ", cycle_len)
			println("")
		}

	}
	println("Total permutations: ", total_permutations)
	println("Set free: ", set_free)
	println("Set free percentage: ", fmt.Sprintf("%.2f", float64(set_free)/float64(total_permutations)*100), "%")
}


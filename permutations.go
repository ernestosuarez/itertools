package  itertools

//GenPermutations generates, given a number n,
//all the n factorial permutations of the integers
//from 0 to n-1
func GenPermutations(n int) <-chan []int {
	if n < 0 {
		panic("Invalid argument")
	}

	ch := make(chan []int)

	go func() {
		var finished bool

		result := make([]int, n)

		for i := range result {
			result[i] = i
		}

		temp := make([]int, n)
		copy(temp, result) // avoid overwriting of result
		ch <- temp

		for {
			finished = true

			for i := n - 1; i > 0; i-- {

				if result[i] > result[i-1] {
					finished = false

					min_greater_index := i
					for j := i + 1; j < n; j++ {
						if result[j] < result[min_greater_index] && result[j] > result[i-1] {
							min_greater_index = j
						}

					}

					result[i-1], result[min_greater_index] = result[min_greater_index], result[i-1]

					//sort from i to n-1
					for j := i; j < n; j++ {
						for k := j + 1; k < n; k++ {
							if result[j] > result[k] {
								result[j], result[k] = result[k], result[j]
							}

						}
					}
					break
				}
			}

			if finished {
				close(ch)
				break
			}
			temp := make([]int, n)
			copy(temp, result) // avoid overwriting of result
			ch <- temp

		}

	}()
	return ch
}

func PermutationsInt(iterable []int, r int) chan []int {

	ch := make(chan []int)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]int, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}

func PermutationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make([]string, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}

func PermutationsList(iterable List, r int) chan List {

	ch := make(chan List)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			for perm := range GenPermutations(r) {
				result := make(List, r)
				for i := 0; i < r; i++ {
					result[i] = iterable[comb[perm[i]]]
				}
				ch <- result
			}
		}

		close(ch)
	}()
	return ch
}

package itertools

//List is a list of elements of any kind/type
type List []interface{}

//GenCombinations generates, from two natural numbers n > r,
//all the possible combinations of r indexes taken from 0 to n-1.
//For example if n=3 and r=2, the result will be:
//[0,1], [0,2] and [1,2]
func GenCombinations(n, r int) <-chan []int {

	if r > n {
		panic("Invalid arguments")
	}

	ch := make(chan []int)

	go func() {
		result := make([]int, r)
		for i := range result {
			result[i] = i
		}

		temp := make([]int, r)
		copy(temp, result) // avoid overwriting of result
		ch <- temp

		for {
			for i := r - 1; i >= 0; i-- {
				if result[i] < i+n-r {
					result[i]++
					for j := 1; j < r-i; j++ {
						result[i+j] = result[i] + j
					}
					temp := make([]int, r)
					copy(temp, result) // avoid overwriting of result
					ch <- temp
					break
				}
			}
			if result[0] >= n-r {
				break
			}
		}
		close(ch)

	}()
	return ch
}

//CombinationsInt generates all the combinations of r elements
//extracted from an slice of integers
func CombinationsInt(iterable []int, r int) chan []int {

	ch := make(chan []int)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]int, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

//CombinationsStr generates all the combinations of r elements
//extracted from an slice of strings
func CombinationsStr(iterable []string, r int) chan []string {

	ch := make(chan []string)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]string, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

//CombinationsList generates all the combinations of r elements
//extracted from a List (an arbitrary list of elements).
//A List can be created for instance, as follows
//myList := List{"a", "b", 13, 3.523}
func CombinationsList(iterable List, r int) chan List {

	ch := make(chan List)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make(List, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

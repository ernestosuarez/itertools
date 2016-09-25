package itertools

import (
	"flag"
	"fmt"
	"testing"
)

//TestCombPerm is just for testing
//the itertools package
func TestCombPerm(t *testing.T) {
	n := flag.Int("n", 7, "total number of elments")
	r := flag.Int("r", 3, "Number of elements to select")
	flag.Parse()

	fmt.Println(*n, *r)
	fmt.Println("\ncombinations")
	for v := range GenCombinations(*n, *r) {
		fmt.Println(v)
	}

	iterable := []int{-1, 2, -3, 4, -5, 6, 7}
	fmt.Println("\ncombinations int")
	for v := range CombinationsInt(iterable, *r) {
		fmt.Println(v)
	}

	iterable2 := []string{"-1", "2", "-3", "4", "-5", "6", "7"}
	fmt.Println("\ncombinations string")
	for v := range CombinationsStr(iterable2, *r) {
		fmt.Println(v)
	}

	mystring := List{"A", "B", "C", "D", "E", "F", 1}
	fmt.Println("\ncombinations List")
	for v := range CombinationsList(mystring, *r) {
		fmt.Println(v)
	}

	fmt.Println("\nPermutations Generator")
	for v := range GenPermutations(3) {
		fmt.Println(v)
	}

	fmt.Println("\nPermutations Int")
	iterable3 := []int{1, 2, 3, 4}
	for v := range PermutationsInt(iterable3, 3) {
		fmt.Println(v)
	}

	fmt.Println("\nPermutations Str")
	iterable4 := []string{"1", "2", "3", "4"}
	for v := range PermutationsStr(iterable4, 3) {
		fmt.Println(v)
	}

	fmt.Println("\nPermutations List")
	iterable5 := List{"1", "2", "3", 5}
	for v := range PermutationsList(iterable5, 3) {
		fmt.Println(v)
	}
}

# itertools

Itertools is a tool kit to generate combinations and permutations of go slices. It works similar to python itertools.

## Installation

    go get -u github.com/ernestosuarez/itertools

## Usage (Examples)
After installation the package has to be imported.
```go
import "github.com/ernestosuarez/itertools"
```

###Generating combinations
*Slices of integers*
```go
// combinations of r = 3 elements chosen from iterable
r := 3
iterable := []int{1, 2, 3, 4 }

for v := range CombinationsInt(iterable, r) {
      fmt.Println(v)
```
output:
```
[1 2 3]
[1 2 4]
[1 3 4]
[2 3 4]
```
*Slices of strings*
```go
// combinations of r = 3 elements chosen from iterable
r := 3
iterable := []string{"A", "B", "C", "D"}

for v := range CombinationsStr(iterable, r) {
      fmt.Println(v)
```
output:
```
[A B C]
[A B D]
[A C D]
[B C D]
```
*Custom type List*

Notice you have to use the right function depending of the type of the slice. For a more general case the user could define a custom type List as follows:
```go
type List []interface{}
```
Then can create a list of heterogeneous types and follow the same procedure. In this case CombinationsList has to be used.

```go
r := 3
myList := List{1, "B", 3, 3.14}

for v := range CombinationsList(myList, r) {
      fmt.Println(v)
```
output:
```
[1 B 3]
[1 B 3.14]
[1 3 3.14]
[B 3 3.14]
```

###Generating permutations
*Slices of integers*

If the number (r) of chosen element is equal to the length of the _iterable_, then we will obtain the r-factorial permutations
```go
// permutations of r = 3 elements chosen from iterable with length 3
r := 3
iterable := []int{1, 2, 3}

for v := range PermutationsInt(iterable, r) {
      fmt.Println(v)
```
output:
```
[1 2 3]
[1 3 2]
[2 1 3]
[2 3 1]
[3 1 2]
[3 2 1]
```

On the other hand we can chose less than len(iterable) elements
```go
// permutations of r = 3 elements chosen from a iterable of length 4
r := 3
iterable := []int{1, 2, 3, 4}

for v := range CombinationsInt(iterable, r) {
      fmt.Println(v)
```
output:
```
[1 2 3]
[1 3 2]
[2 1 3]
[2 3 1]
[3 1 2]
[3 2 1]
[1 2 4]
[1 4 2]
[2 1 4]
[2 4 1]
[4 1 2]
[4 2 1]
[1 3 4]
[1 4 3]
[3 1 4]
[3 4 1]
[4 1 3]
[4 3 1]
[2 3 4]
[2 4 3]
[3 2 4]
[3 4 2]
[4 2 3]
[4 3 2]
```

*Slices of strings*

Same idea, just use the function PermutationsStr

*Custom type List*

Same as in combinations, fist define the new type List (see above) and then use PermutationsList


## Contributions

Contributions are welcome.

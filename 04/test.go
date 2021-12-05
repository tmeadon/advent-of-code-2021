package main

import "fmt"

func main() {
	tst := []string{"a", "b", "c"}
	tst = remove(tst, "b")
	fmt.Println(tst)
}

func blah(input *[]string) {
	for i := 0; i < len(*input); i++ {
		// (*input)[i] = "blah"
		if i == 2 {
			(*input)[i] = "blah"
		}
	}
}

func remove(input []string, str string) []string {
	for i, o := range input {
		if o == str {
			input = append(input[:i], input[i+1:]...)
		}
	}
	return input
}

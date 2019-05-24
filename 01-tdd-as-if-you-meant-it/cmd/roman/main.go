package main

import (
	"fmt"
)

func main() {
	fmt.Println(IntToRoman(27))
	fmt.Println(IntToRoman(3))
}

func IntToRoman(i int) string {
	roman := ""

	if i >= 10 {
		for i := 0; i < i/10; i++ {
			roman += "X"
		}

		i %= 10
	}

	if i >= 1 {
		if i == 9 {
			roman += "IX"
		} else if i >= 5 {
			roman += "V"

			for i := 0; i < i-5; i++ {
				roman += "I"
			}
		} else if i == 4 {
			roman += "IV"
		} else if i >= 1 {
			for i := 0; i < i; i++ {
				roman += "I"
			}
		}
	}

	return roman
}

/**
I 1
V 5
X 10

- I, X, C, M can repeat up to 3 times
  III = 1 + 1 + 1 = 3

- V, L, D cannot repeat

- Numerals that decrease left to right are added
  XVI = 10 + 5 + 1 = 16

- A smaller numeral left of a larger numeral is subtracted
  IV = 5 - 1 = 4
*/

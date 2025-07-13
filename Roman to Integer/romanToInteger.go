package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	fmt.Printf("Starting conversion for Roman numeral: %s\n", s)
	for _, char := range s {
		value := romanMap[char]
		fmt.Printf("Processing character: %c, Value: %d, Previous Value: %d\n", char, value, prevValue)
		if value > prevValue {
			fmt.Printf("Found a larger value: %d, subtracting twice the previous value: %d\n", value, prevValue)
			total += value - 2*prevValue // subtract the previous value twice
		} else {
			fmt.Printf("Adding value: %d\n", value)
			total += value
		}
		fmt.Printf("Current total: %d\n", total)
		prevValue = value
	}

	fmt.Printf("Final total: %d\n", total)
	return total

}

func main() {
	fmt.Println(romanToInt("III")) // Output: 3
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("IV")) // Output: 4
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("IX")) // Output: 9
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("LVIII")) // Output: 58
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("MMXXIII")) // Output: 2023
	fmt.Print("---------------------\n")
	fmt.Println(romanToInt("XLII")) // Output: 42
}

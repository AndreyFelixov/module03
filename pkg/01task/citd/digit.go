package citd

import "01task/wordz"

func Digit() string {
	wordz.Words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return wordz.Random()
}

package helper

import (
	"log"
	"time"
)

// Letters func
func Letters(base int) []string {
	letters := make([]string, base)

	for i := 0; i < base/2; i++ {
		letters[i], letters[i+base/2] = string('a'+i), string('A'+i)
	}

	return letters
}

// Code func
func Code(num int, base int) string {
	code := ""

	letters := Letters(base)

	for num > 0 {
		num--
		code = letters[num%base] + code
		num /= base
	}

	return code
}

// Codes func
func Codes(nums int, base int) []string {
	codes := make([]string, nums)

	for i := 0; i < nums; i++ {
		codes[i] = Code(i, base)
	}

	return codes
}

// Measure func
func Measure(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}

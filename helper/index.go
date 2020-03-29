package helper

import (
	"log"
	"time"
)

// Letters generates different ASCII characters.
func Letters(base int) []string {
	letters := make([]string, base)

	for i := 0; i < base/2; i++ {
		letters[i], letters[i+base/2] = string('a'+i), string('A'+i)
	}

	return letters
}

// Code returns the letter according to the given number.
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

// Measure measures the execution time.
func Measure(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}

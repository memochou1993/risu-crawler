package helper

const (
	base = 52
)

// Letters func
func Letters() []string {
	letters := make([]string, base)

	for i := 0; i < base/2; i++ {
		letters[i], letters[i+base/2] = string('a'+i), string('A'+i)
	}

	return letters
}

// Code func
func Code(num int) string {
	code := ""

	for num > 0 {
		num--
		code = Letters()[num%base] + code
		num /= base
	}

	return code
}

// Codes func
func Codes(nums int) []string {
	codes := make([]string, nums)

	for i := 0; i < nums; i++ {
		codes[i] = Code(i)
	}

	return codes
}

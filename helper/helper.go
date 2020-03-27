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

// Codes func
func Codes(index int) []string {
	codes := []string{}

	for index > 0 {
		index--
		codes = append([]string{Letters()[index%base]}, codes...)
		index /= base
	}

	return codes
}

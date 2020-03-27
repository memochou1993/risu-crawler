package controller

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/memochou1993/crawler/helper"
)

// Handle func
func Handle() {
	nums := int(math.Pow(52, 3))
	codes := getCodes(nums)

	fmt.Println(codes)
}

func getCodes(nums int) []string {
	codes := helper.Codes(nums)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(codes), func(i, j int) {
		codes[i], codes[j] = codes[j], codes[i]
	})

	return codes
}

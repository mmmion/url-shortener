package utils

import "strings"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 62 characters

func Encode62(number uint64) string {
	if number == 0 {
		return string(charset[0])
	}

	str := ""
	for number > 0 {
		remainder := number % 62
		str = string(charset[remainder]) + str
		number /= 62
	}
	return str
}

func Decode62(base62 string) uint64 {
	var total uint64 = 0
	for i := 0; i < len(base62); i++ {
		char := strings.IndexByte(charset, base62[i]) // Or substrat the ASCII values for each category

		total = (total * 62) + uint64(char)
	}
	return total
}

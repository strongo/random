package random

// ID creates a random ID of requested length
func ID(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = chars[r.Int63()%int64(len(chars))]
	}
	return string(b)
}

// Digits creates a string of random digits
func Digits(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = digits[r.Int63()%int64(len(digits))]
	}
	return string(b)
}

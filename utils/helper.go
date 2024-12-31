package utils

func Truncate(input string, maxLength int) string {
	if len(input) > maxLength {
		return input[:maxLength-3] + "..."
	}
	return input
}

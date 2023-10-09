package longest_substring_without_repeating_characters

func longestSubstringWithoutRepeatingCharacters(s string) int {
	var longest int
	var start int

	for i, char := range s {
		if index := findIndex(s[start:i], char); index != -1 {
			start += index + 1
		}

		if length := i - start + 1; length > longest {
			longest = length
		}
	}

	return longest
}

func findIndex(str string, char rune) int {
	for i, c := range str {
		if c == char {
			return i
		}
	}

	return -1
}

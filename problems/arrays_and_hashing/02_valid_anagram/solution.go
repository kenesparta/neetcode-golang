package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, ch := range s {
		counts[ch]++
	}

	for _, ch := range t {
		count, ok := counts[ch]

		if !ok || count == 0 {
			return false
		}

		counts[ch]--
	}

	return true
}

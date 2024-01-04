package palindrome

func LongestPalindrome(s string) string {
	// longest palindromes
	var longest string
	// iterate over the string
	for i := 0; i < len(s); i++ {
		// find the longest odd palindrome
		odd := expand(s, i, i)
		// find the longest even palindrome
		even := expand(s, i, i+1)
		// compare the longest palindromes
		if len(odd) > len(longest) {
			longest = odd
		}
		if len(even) > len(longest) {
			longest = even
		}
	}

	return longest
}

func expand(s string, left, right int) string {
	// iterate over the string
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// expand the substring
		left--
		right++
	}

	return s[left+1 : right]
}

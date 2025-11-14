// 3228. Maximum Number of Operations to Move Ones to the End
package p3228

func maxOperations(s string) int {
	result := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		}

		if s[i] == '0' && (i == len(s) - 1 || s[i + 1] == '1') {
			result += count
		}
	}

	return result
}

package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	// stack
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s); i++ {
		p := s[i]
		if _, ok := m[p]; ok {
			stack = append(stack, p)
		} else if len(stack) > 0 && m[stack[len(stack)-1]] == p {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false

	// other
	// for strings.Contains(s, "{}") || strings.Contains(s, "[]") || strings.Contains(s, "()") {
	// 	s = strings.ReplaceAll(s, "{}", "")
	// 	s = strings.ReplaceAll(s, "[]", "")
	// 	s = strings.ReplaceAll(s, "()", "")
	// }
	// return s == ""
}

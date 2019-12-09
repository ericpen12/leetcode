package to_lower_case

func ToLowerCase(str string) string {
	s := []byte(str)
	for i, v := range s {
		s[i] = v|32
	}
	return string(s)
}

// 1. 问题：转小写字母
// 2. 方法： ascii 码

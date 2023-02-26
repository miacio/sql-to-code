package sqltools

// HumpNaming 驼峰命名
func HumpNaming(s string) string {
	var result string
	for i, v := range s {
		if v == '_' {
			continue
		}
		if i == 0 || (i > 0 && s[i-1] == '_') {
			if v > 96 && v < 123 {
				result += string(v - 32)
				continue
			}
		}
		result += string(v)
	}
	return result
}

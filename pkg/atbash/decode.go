package jjAtbash

func Decode(input string) string {
	var result []rune

	for _, char := range input {
		if char >= 'A' && char <= 'Z' {
			result = append(result, 'Z'-(char-'A'))
		} else if char >= 'a' && char <= 'z' {
			result = append(result, 'z'-(char-'a'))
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

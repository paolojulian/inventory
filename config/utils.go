package config

func StringPointer(s string) *string {
	return &s
}

func IntPointer(i int) *int {
	return &i
}

func BoolPointer(b bool) *bool {
	return &b
}

func Includes(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}

	return false
}

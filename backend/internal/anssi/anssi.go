package anssi

type CheckResult struct {
	IsValid bool
	Missing []Recommendation
}

func CheckPassword(password string) CheckResult {
	missing := []Recommendation{}

	if len(password) < 12 {
		missing = append(missing, MinLength)
	}
	if !containsUppercase(password) {
		missing = append(missing, UppercaseLetters)
	}
	if !containsLowercase(password) {
		missing = append(missing, LowercaseLetters)
	}
	if !containsNumber(password) {
		missing = append(missing, NumericCharacters)
	}
	if !containsSpecialChar(password) {
		missing = append(missing, SpecialCharacters)
	}

	return CheckResult{
		IsValid: len(missing) == 0,
		Missing: missing,
	}
}

func containsUppercase(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func containsNumber(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func containsSpecialChar(s string) bool {
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"
	for _, c := range s {
		if contains(specialChars, c) {
			return true
		}
	}
	return false
}

func contains(s string, c rune) bool {
	for _, sc := range s {
		if sc == c {
			return true
		}
	}
	return false
}

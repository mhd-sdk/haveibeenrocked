package anssi

import (
	"reflect"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     CheckResult
	}{
		{
			name:     "Valid password with all requirements",
			password: "Abcdef12345!",
			want: CheckResult{
				IsValid: true,
				Missing: []Recommendation{},
			},
		},
		{
			name:     "Password too short",
			password: "Abc123!",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{MinLength},
			},
		},
		{
			name:     "Password without uppercase",
			password: "abcdef12345!",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{UppercaseLetters},
			},
		},
		{
			name:     "Password without lowercase",
			password: "ABCDEF12345!",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{LowercaseLetters},
			},
		},
		{
			name:     "Password without numbers",
			password: "AbcdefGHIJK!",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{NumericCharacters},
			},
		},
		{
			name:     "Password without special characters",
			password: "Abcdef123456",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{SpecialCharacters},
			},
		},
		{
			name:     "Empty password",
			password: "",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{MinLength, UppercaseLetters, LowercaseLetters, NumericCharacters, SpecialCharacters},
			},
		},
		{
			name:     "Password with multiple missing requirements",
			password: "abcdef",
			want: CheckResult{
				IsValid: false,
				Missing: []Recommendation{MinLength, UppercaseLetters, NumericCharacters, SpecialCharacters},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPassword(tt.password)
			if got.IsValid != tt.want.IsValid {
				t.Errorf("CheckPassword() IsValid = %v, want %v", got.IsValid, tt.want.IsValid)
			}
			if len(got.Missing) == len(tt.want.Missing) && len(got.Missing) == 0 {
				return
			}
			if !reflect.DeepEqual(got.Missing, tt.want.Missing) {
				t.Errorf("CheckPassword() Missing = %v, want %v", got.Missing, tt.want.Missing)
			}
		})
	}
}

func TestContainsUppercase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"With uppercase", "abcDef", true},
		{"Without uppercase", "abcdef", false},
		{"Only uppercase", "ABCDEF", true},
		{"Empty string", "", false},
		{"Special chars only", "!@#$%^", false},
		{"Numbers only", "12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsUppercase(tt.s); got != tt.want {
				t.Errorf("containsUppercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsLowercase(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"With lowercase", "ABCdef", true},
		{"Without lowercase", "ABCDEF", false},
		{"Only lowercase", "abcdef", true},
		{"Empty string", "", false},
		{"Special chars only", "!@#$%^", false},
		{"Numbers only", "12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsLowercase(tt.s); got != tt.want {
				t.Errorf("containsLowercase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"With number", "abc123", true},
		{"Without number", "abcdef", false},
		{"Only numbers", "12345", true},
		{"Empty string", "", false},
		{"Special chars only", "!@#$%^", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNumber(tt.s); got != tt.want {
				t.Errorf("containsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsSpecialChar(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"With special char", "abc!", true},
		{"With different special char", "abc@123", true},
		{"Without special char", "abc123", false},
		{"Only special chars", "!@#$%^", true},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsSpecialChar(tt.s); got != tt.want {
				t.Errorf("containsSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		s    string
		c    rune
		want bool
	}{
		{"Character present", "abcdef", 'c', true},
		{"Character not present", "abcdef", 'z', false},
		{"Empty string", "", 'a', false},
		{"Special character present", "!@#$%", '@', true},
		{"Special character not present", "!@#$%", '&', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.s, tt.c); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

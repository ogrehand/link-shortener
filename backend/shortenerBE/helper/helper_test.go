package helper

import (
	"regexp"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	want := regexp.MustCompile(`.{32}`)
	msg := GenerateToken()
	if !want.MatchString(msg) {
		t.Fatalf(`GenerateToken() = %q, want match for %#q, nil`, msg, want)
	}
}

func TestGenerateSalt(t *testing.T) {
	want := regexp.MustCompile(`.{32}`)
	msg := GenerateToken()
	if !want.MatchString(msg) {
		t.Fatalf(`GenerateToken() = %q, want match for %#q, nil`, msg, want)
	}
}

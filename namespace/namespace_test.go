package namespace

import (
	"strings"
	"testing"
)

func TestCamelAlnum(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"some-str", "SomeStr"},
		{"some str", "SomeStr"},
		{"SoMe STr", "SomeStr"},
		{"@SoMe!STr", "SomeStr"},
	}
	for _, tt := range tests {
		if got := NewCamelCase(tt.in); got != strings.TrimSpace(tt.want) {
			t.Errorf("ToCamelCaseAlnum(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"n", true},
		{"name-space", true},
		{"name_space", true},
		{"2", false},
		{"NameSpace", false},
		{"NameSpace2", false},
		{"2NameSpace", false},
		{"name space", false},
		{"name_space ", false},
		{" name_space", false},
		{"name_space_", false},
		{"_name_space", false},
		{"name_space-", false},
		{"-name_space", false},
		{"CamelCase ", false},
		{"~abc", false},
		{"a@bc", false},
	}
	for _, tt := range tests {
		if got := ValidSlug(tt.in); got != tt.want {
			t.Errorf("ValidSlug(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestSlug(t *testing.T) {
	if slug, expected := Slug("test->àèâ<-test"), "test-aea-test"; slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}

func TestLowerOption(t *testing.T) {
	if slug, expected := Slug("Test->àèâ<-Test"), "test-aea-test"; slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
}

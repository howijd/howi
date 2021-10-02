package namespace

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
)

const (
	// NamespaceMustCompile against following expression.
	SlugRe = "^[a-z][a-z0-9_-]*[a-z0-9]$"
)

// Replacement structure.
type replacement struct {
	re *regexp.Regexp
	ch string
}

// Build regexps and replacements.
var (
	rExps = []replacement{ //nolint:gochecknoglobals
		{re: regexp.MustCompile(`[\xC0-\xC6]`), ch: "A"},
		{re: regexp.MustCompile(`[\xE0-\xE6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xC8-\xCB]`), ch: "E"},
		{re: regexp.MustCompile(`[\xE8-\xEB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xCC-\xCF]`), ch: "I"},
		{re: regexp.MustCompile(`[\xEC-\xEF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xD2-\xD6]`), ch: "O"},
		{re: regexp.MustCompile(`[\xF2-\xF6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xD9-\xDC]`), ch: "U"},
		{re: regexp.MustCompile(`[\xF9-\xFC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xC7-\xE7]`), ch: "c"},
		{re: regexp.MustCompile(`[\xD1]`), ch: "N"},
		{re: regexp.MustCompile(`[\xF1]`), ch: "n"},
	}
	spacereg       = regexp.MustCompile(`\s+`)
	noncharreg     = regexp.MustCompile(`[^A-Za-z0-9-]`)
	minusrepeatreg = regexp.MustCompile(`\-{2,}`)
	alnum          = &unicode.RangeTable{ //nolint:gochecknoglobals
		R16: []unicode.Range16{
			{'0', '9', 1},
			{'A', 'Z', 1},
			{'a', 'z', 1},
		},
	}
)

// NewCamelCase returns a camel case representation of the string all
// non alpha numeric characters removed. Uppercase characters are mapped
// first alnum in string and after each non alnum character is removed.
func NewCamelCase(s string) string {
	var b bytes.Buffer
	tu := true
	for _, c := range s {
		isAlnum := unicode.Is(alnum, c)
		isSpace := unicode.IsSpace(c)
		isLower := unicode.IsLower(c)
		if isSpace || !isAlnum {
			tu = true

			continue
		}
		if tu {
			if isLower {
				b.WriteRune(unicode.ToUpper(c))
			} else {
				b.WriteRune(c)
			}
			tu = false

			continue
		} else {
			if !isLower {
				c = unicode.ToLower(c)
			}
			b.WriteRune(c)
		}
	}
	return b.String()
}

// Slug function returns slugifies string "s".
func Slug(s string) string {
	for _, r := range rExps {
		s = r.re.ReplaceAllString(s, r.ch)
	}

	s = strings.ToLower(s)
	s = spacereg.ReplaceAllString(s, "-")
	s = noncharreg.ReplaceAllString(s, "")
	s = minusrepeatreg.ReplaceAllString(s, "-")

	return s
}

// ValidSlug returns true if s is string which is valid slug.
func ValidSlug(s string) bool {
	if len(s) == 1 {
		return unicode.IsLetter(rune(s[0]))
	}
	re := regexp.MustCompile(SlugRe)
	return re.MatchString(s)
}

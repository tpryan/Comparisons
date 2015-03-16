package rules

import (
	"regexp"
	"strings"
)

const (
	minMatch       = 4
	minLength      = 8
	maxLen         = 24
	success        = "Password passes policy"
	failEmpty      = "No password given"
	failUpper      = "At least one UPPERCASE character is required."
	failLower      = "At least one LOWERCASE character is required."
	failNumber     = "At least one NUMERIC character is required."
	failSpecial    = "At least one SPECIAL (~!@#$%^&*) character is required."
	failDictionary = "No dictionary words allowed."
	failMin        = "Password must be at least 8 characters long."
	failMax        = "Password must be no more than 24 characters long."
)

var (
	lr      = regexp.MustCompile("[a-z]")
	ur      = regexp.MustCompile("[A-Z]")
	nr      = regexp.MustCompile("[0-9]")
	sr      = regexp.MustCompile("[~!@#$%^&*]")
	dict    = GetDict()
	dictmap = GetDictMap()
)

type Result struct {
	Pass    bool
	Message string
	Status  string
	Word    string
}

func Validate(c string, m string) Result {

	if len(c) == 0 {
		return Result{false, failEmpty, "FAIL_EMPTY", ""}
	}

	if len(c) < minLength {
		return Result{false, failMin, "FAIL_MIN", ""}
	}

	if len(c) > maxLen {
		return Result{false, failMax, "FAIL_MAX", ""}
	}

	if r := ur.FindStringIndex(c); r == nil {
		return Result{false, failUpper, "FAIL_UPPER", ""}
	}

	if r := lr.FindStringIndex(c); r == nil {
		return Result{false, failLower, "FAIL_LOWER", ""}
	}

	if r := nr.FindStringIndex(c); r == nil {
		return Result{false, failNumber, "FAIL_NUMBER", ""}
	}

	if r := sr.FindStringIndex(c); r == nil {
		return Result{false, failSpecial, "FAIL_SPECIAL", ""}
	}

	w := ""
	if m == "bruteforce" {
		w = match(c)
	} else {
		w = hashMatch(c)
	}

	if w != "" {
		return Result{false, failDictionary, "FAIL_DICTIONARY", w}
	}

	return Result{true, success, "SUCCESS", w}
}

func match(c string) string {
	uc := strings.ToUpper(c)
	for _, w := range dict {

		if l := len(w); l < minMatch || l > len(c) {
			continue
		}

		if strings.Index(uc, w) >= 0 {
			return w
		}
	}

	return ""
}

func hashMatch(c string) string {
	hmap := breakString(c, minMatch)

	for k := range hmap {
		if _, ok := dictmap[k]; ok {
			return k
		}
	}
	return ""
}

// breakstring breaks a string into all substrings with a length
// greater than @min
func breakString(s string, min int) map[string]int {
	res := make(map[string]int)
	ln := len(s)
	for i := min; i <= ln; i++ {
		for j := 0; j < (ln - min); j++ {

			if i+j > ln {
				continue
			}

			p := strings.ToUpper(s[j : i+j])

			if len(p) >= i {
				res[p] = 0
			}
		}
	}

	return res
}

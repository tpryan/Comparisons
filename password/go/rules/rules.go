package rules

import (
	"regexp"
	"strings"
)

const (
	MinMatch       = 4
	MinLength      = 8
	MaxLen         = 24
	Success        = "Password passes policy"
	FailEmpty      = "No password given"
	FailUpper      = "At least one UPPERCASE character is required."
	FailLower      = "At least one LOWERCASE character is required."
	FailNumber     = "At least one NUMERIC character is required."
	FailSpecial    = "At least one SPECIAL (~!@#$%^&*) character is required."
	FailDictionary = "No dictionary words allowed."
	FailMin        = "Password must be at least 8 characters long."
	FailMax        = "Password must be no more than 24 characters long."
	FailError      = "Error in matchmode must be 'hash' or 'bruteforce'"
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

var (
	resultEmpty      = Result{false, FailEmpty, "FAIL_EMPTY", ""}
	resultUpper      = Result{false, FailUpper, "FAIL_UPPER", ""}
	resultLower      = Result{false, FailLower, "FAIL_LOWER", ""}
	resultNumber     = Result{false, FailNumber, "FAIL_NUMBER", ""}
	resultSpecial    = Result{false, FailSpecial, "FAIL_SPECIAL", ""}
	resultDictionary = Result{false, FailDictionary, "FAIL_DICTIONARY", ""}
	resultMin        = Result{false, FailMin, "FAIL_MIN", ""}
	resultMax        = Result{false, FailMax, "FAIL_MAX", ""}
	resultError      = Result{false, FailError, "FAIL_ERROR", ""}
	resultSuccess    = Result{true, Success, "SUCCESS", ""}
)

type MatchMode string

const (
	Bruteforce MatchMode = "bruteforce"
	Hash       MatchMode = "hash"
)

func Validate(c string, m MatchMode) Result {
	switch {
	case len(c) == 0:
		return resultEmpty
	case len(c) < MinLength:
		return resultMin
	case len(c) > MaxLen:
		return resultMax
	case ur.FindStringIndex(c) != nil:
		return resultUpper
	case lr.FindStringIndex(c) != nil:
		return resultLower
	case nr.FindStringIndex(c) != nil:
		return resultNumber
	case sr.FindStringIndex(c) != nil:
		return resultSpecial
	}

	w := ""
	switch m {
	case Bruteforce:
		w = match(c)
	case Hash:
		w = hashMatch(c)
	default:
		return resultError
	}

	if w != "" {
		r := resultDictionary
		r.Word = w
		return r
	}

	return resultSuccess
}

func match(c string) string {
	uc := strings.ToUpper(c)
	for _, w := range dict {

		if l := len(w); l < MinMatch || l > len(c) {
			continue
		}

		if strings.Index(uc, w) >= 0 {
			return w
		}
	}

	return ""
}

func hashMatch(c string) string {
	hmap := breakString(c, MinMatch)

	for k := range hmap {
		if dictmap[k] {
			return k
		}
	}
	return ""
}

// breakstring breaks a string into all substrings with a length
// greater than min, returning ...
func breakString(s string, min int) map[string]bool {
	res := make(map[string]bool)
	ln := len(s)
	for i := min; i <= ln; i++ {
		for j := 0; j < (ln - min); j++ {

			if i+j > ln {
				continue
			}

			p := strings.ToUpper(s[j : i+j])

			if len(p) >= i {
				res[p] = true
			}
		}
	}

	return res
}

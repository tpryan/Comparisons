package Rules

import (
	"regexp"
	"strings"
)

const (
	MINIMUM_MATCH   = 4
	MIN_LENGTH      = 8
	MAX_LENGTH      = 24
	SPECIAL         = "~!@#$%^&*"
	SUCCESS         = "Password passes policy"
	FAIL_EMPTY      = "No password given"
	FAIL_UPPER      = "At least one UPPERCASE character is required."
	FAIL_LOWER      = "At least one LOWERCASE character is required."
	FAIL_NUMBER     = "At least one NUMERIC character is required."
	FAIL_SPECIAL    = "At least one SPECIAL (~!@#$%^&*) character is required."
	FAIL_DICTIONARY = "No dictionary words allowed."
	FAIL_MIN        = "Password must be at least 8 characters long."
	FAIL_MAX        = "Password must be no more than 24 characters long."
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
		return Result{false, FAIL_EMPTY, "FAIL_EMPTY", ""}
	}

	if len(c) < MIN_LENGTH {
		return Result{false, FAIL_MIN, "FAIL_MIN", ""}
	}

	if len(c) > MAX_LENGTH {
		return Result{false, FAIL_MAX, "FAIL_MAX", ""}
	}

	if r := ur.FindStringIndex(c); r == nil {
		return Result{false, FAIL_UPPER, "FAIL_UPPER", ""}
	}

	if r := lr.FindStringIndex(c); r == nil {
		return Result{false, FAIL_LOWER, "FAIL_LOWER", ""}
	}

	if r := nr.FindStringIndex(c); r == nil {
		return Result{false, FAIL_NUMBER, "FAIL_NUMBER", ""}
	}

	if r := sr.FindStringIndex(c); r == nil {
		return Result{false, FAIL_SPECIAL, "FAIL_SPECIAL", ""}
	}

	w := ""
	if m == "bruteforce" {
		w = match(c)
	} else {
		w = hashMatch(c)
	}

	if w != "" {
		return Result{false, FAIL_DICTIONARY, "FAIL_DICTIONARY", w}
	}

	return Result{true, SUCCESS, "SUCCESS", ""}
}

func match(c string) string {
	uc := strings.ToUpper(c)
	for _, w := range dict {

		if len(w) < MINIMUM_MATCH {
			continue
		}
		if len(w) > len(c) {
			continue
		}
		if strings.Index(uc, w) >= 0 {
			return w
		}
	}

	return ""
}

func hashMatch(c string) string {
	hmap := breakString(c, MINIMUM_MATCH)

	for k := range hmap {
		if _, ok := dictmap[k]; ok {
			return k
		}
	}
	return ""
}

func breakString(s string, m int) map[string]int {
	res := make(map[string]int)
	ln := len(s)
	for i := m; i <= ln; i++ {
		for j := 0; j < (ln - m); j++ {

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

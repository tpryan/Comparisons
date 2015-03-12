package Rules

import (
	"testing"
)

func TestHandlesEmptyInput(t *testing.T) {

	in := ""
	got := Validate(in)
	want := Result{false, FAIL_EMPTY, "FAIL_EMPTY", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_EMPTY {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesMinCheck(t *testing.T) {

	in := "dasdsfg"
	got := Validate(in)
	want := Result{false, FAIL_MIN, "FAIL_MIN", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_MIN {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesMaxCheck(t *testing.T) {

	in := "1234567890123456789012345"
	got := Validate(in)
	want := Result{false, FAIL_MAX, "FAIL_MAX", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_MAX {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesNoUpper(t *testing.T) {

	in := "dasdasdasdasd"
	got := Validate(in)
	want := Result{false, FAIL_UPPER, "FAIL_UPPER", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_UPPER {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesNoLower(t *testing.T) {

	in := "DKRKASDKEKASKD"
	got := Validate(in)
	want := Result{false, FAIL_LOWER, "FAIL_LOWER", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_LOWER {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesNoNumeric(t *testing.T) {

	in := "Drdfjflrmg"
	got := Validate(in)
	want := Result{false, FAIL_NUMBER, "FAIL_NUMBER", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_NUMBER {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesNoSpecial(t *testing.T) {

	in := "Drdfjflr9mg"
	got := Validate(in)
	want := Result{false, FAIL_SPECIAL, "FAIL_SPECIAL", ""}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_SPECIAL {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

func TestHandlesDictionaryPresent(t *testing.T) {

	in := "Drdfjflr9mg&Apple"
	got := Validate(in)
	want := Result{false, FAIL_DICTIONARY, "FAIL_DICTIONARY", "APPLE"}

	if got.Pass != false {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != FAIL_DICTIONARY {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
	if got.Word != want.Word {
		t.Errorf("Validate(%q).Word == %q, want %q", in, got.Word, want.Word)
	}
}

func TestHandlesValid(t *testing.T) {

	in := "Drdfjflr9mg&"
	got := Validate(in)
	want := Result{true, SUCCESS, "SUCCESS", ""}

	if got.Pass != true {
		t.Errorf("Validate(%q).Pass == %q, want %q", in, got.Pass, want.Pass)
	}
	if got.Message != SUCCESS {
		t.Errorf("Validate(%q).Message == %q, want %q", in, got.Message, want.Message)
	}
}

/*
   Copyright 2015, Google, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package rules

import (
	"testing"
)

func TestValidation(t *testing.T) {
	cases := []struct {
		input string
		mode  string
		err   string
		word  string
		pass  bool
	}{
		{"", "bruteforce", failEmpty, "", false},
		{"dasdsfg", "bruteforce", failMin, "", false},
		{"1234567890123456789012345", "bruteforce", failMax, "", false},
		{"dasdasdasdasd", "bruteforce", failUpper, "", false},
		{"DKRKASDKEKASKD", "bruteforce", failLower, "", false},
		{"Drdfjflrmg", "bruteforce", failNumber, "", false},
		{"Drdfjflr9mg", "bruteforce", failSpecial, "", false},
		{"Drdfjflr9mg&Apple", "bruteforce", failDictionary, "APPLE", false},
		{"Drdfjflr9mg&", "bruteforce", success, "", true},
		{"Drdfjflr9mg&Apple", "hash", failDictionary, "APPLE", false},
		{"Drdfjflr9mg&", "hash", success, "", true},
	}

	for _, c := range cases {
		got := Validate(c.input, c.mode)

		if got.Pass != c.pass {
			t.Errorf("Validate(%q).Pass == %q, want %q", c.input, got.Pass, c.pass)
		}
		if got.Message != c.err {
			t.Errorf("Validate(%q).Message == %q, want %q", c.input, got.Message, c.err)
		}
		if got.Word != c.word {
			t.Errorf("Validate(%q).Word == %q, want %q", c.input, got.Word, c.word)
		}

	}

}

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
package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/tpryan/comps/password/go/rules"
)

var (
	count  = flag.Int("count", 1, "the max number of passwords to process.")
	method = flag.String("method", "bruteforce", "the way to process the passwords")
)

func main() {
	flag.Parse()

	f, err := os.Open("password/data/test_passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	i := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		if i > *count {
			break
		}

		l := s.Text()
		switch *method {
		case "bruteforce":
			_ = rules.Validate(l, rules.Bruteforce)
		case "hash":
			_ = rules.Validate(l, rules.Hash)
		default:
			log.Fatalf("%v\n", rules.FailError)
		}

		i++
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

}

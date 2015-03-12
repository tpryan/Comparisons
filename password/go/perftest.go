package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/tpryan/comps/password/go/rules"
)

func main() {

	loopcount, err := strconv.Atoi(os.Args[1])

	file, err := os.Open("password/data/test_passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i > loopcount {
			break
		}

		l := scanner.Text()
		res := Rules.Validate(l)
		_ = res
		//fmt.Printf("%v\n", res)

		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

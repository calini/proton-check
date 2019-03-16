package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// pushes strings to stdin
func main() {
	flag.Parse()

	alphabet := flag.Arg(0)

	l := flag.Arg(1)
	// string to int
	length, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range All(length, strings.Split(alphabet, "")) {
		fmt.Println(str)
	}
}

// All returns all strings of legnth len formed from the alphabet
func All(len int, alphabet []string) []string {
	if len <= 0 {
		return nil
	}

	if len == 1 {
		return alphabet
	}

	var res []string
	tails := All(len-1, alphabet)
	for _, chi := range alphabet {
		for _, tail := range tails {
			res = append(res, chi+tail)
		}
	}
	return res
}

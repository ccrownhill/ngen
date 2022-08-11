/*
A simple random name generator.
The only criterium is that a human can still speak the name.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const vowels string = "aeiou"
const consonants string = "bcdfghjklmnpqrstvwxyz"
const defaultLen int = 7

func main() {
	var outname string = ""
	var chooseVowel bool
	var maxlen int = defaultLen

	if len(os.Args) >= 2 {
		inputLen, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, os.Args[0], "[max name length]")
			os.Exit(1)
		}
		maxlen = inputLen
	}

	rand.Seed(time.Now().UnixNano())

	chooseVowel = rand.Intn(2) == 1

	for i := 0; i < maxlen; i++ {
		if chooseVowel {
			outname += string(getRandChar(vowels))
		} else {
			outname += string(getRandChar(consonants))
		}
		chooseVowel = !chooseVowel
	}

	fmt.Println(outname)
}

// Return one random rune in a string
func getRandChar(charset string) rune {
	runeCharset := []rune(charset)
	return runeCharset[rand.Intn(len(runeCharset))]
}

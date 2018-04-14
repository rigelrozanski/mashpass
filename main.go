package main

import "fmt"

/*
Creation:
 - collect a number of mashes
 - determine the common keyset (aka, the 1st, 3rd, 5th, and 6th keys)
   - this is the private key
 - save the hashes of the all the mashes
 - print the security level (aka time to crack)

Authentication:
 - collect a mash
 - iterate through all possibilities stepwise starting with 1 key inccorect, then 2, and soforth
   - stop after 10 seconds of searching to request a retrial
 - once a correct hash has been found, create the common keyset aka the private key
*/

var (
	macbookKeys = map[string][]string{
		"1": {"2", "q"}, // top row
		"2": {"1", "q", "w", "3"},
		"3": {"2", "w", "e", "4"},
		"4": {"3", "e", "r", "5"},
		"5": {"4", "r", "t", "6"},
		"6": {"5", "t", "y", "7"},
		"7": {"6", "y", "u", "8"},
		"8": {"7", "u", "i", "9"},
		"9": {"8", "i", "o", "0"},
		"0": {"9", "o", "p", "-"},
		"-": {"0", "p", "[", "="},
		"=": {"-", "[", "]"},
		"q": {"1", "2", "w", "a", "s"}, // second row
		"w": {"2", "3", "q", "e", "a", "s"},
		"e": {"3", "4", "w", "r", "s", "d"},
		"r": {"4", "5", "e", "t", "d", "f"},
		"t": {"5", "6", "r", "y", "f", "g"},
		"y": {"6", "7", "t", "u", "g", "h"},
		"u": {"7", "8", "y", "i", "h", "j"},
		"i": {"8", "9", "u", "o", "j", "k"},
		"o": {"9", "0", "i", "p", "k", "l"},
		"p": {"0", "-", "o", "[", "l", ";"},
		"[": {"-", "=", "p", "]", ";", "'"},
		"]": {"=", "[", ";", "'"},
		"a": {"q", "w", "s", "z", "x"}, // third row
		"s": {"w", "e", "a", "d", "z", "x"},
		"d": {"e", "r", "s", "f", "x", "c"},
		"f": {"r", "t", "d", "g", "c", "v"},
		"g": {"t", "y", "f", "h", "v", "b"},
		"h": {"y", "u", "g", "j", "b", "n"},
		"j": {"u", "i", "h", "k", "n", "m"},
		"k": {"i", "o", "j", "l", "m", ","},
		"l": {"o", "p", "k", ";", ",", "."},
		";": {"p", "[", "l", "'", ".", "/"},
		"'": {"[", "]", ";", ".", "/"},
		"z": {"a", "s", "x"}, //bottom row
		"x": {"s", "d", "z", "c"},
		"c": {"d", "f", "x", "v"},
		"v": {"f", "g", "c", "b"},
		"b": {"g", "h", "v", "n"},
		"n": {"h", "j", "b", "m"},
		"m": {"j", "k", "n", ","},
		",": {"k", "l", "m", "."},
		"/": {";", "'", "."},
	}
)

func main() {
	fmt.Println("vim-go")
}

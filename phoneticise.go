// Package phoneticise provides a library for expanding strings to NATO phonetics
package phoneticise

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

// Phoneticise takes a string and outputs the phonetic alphabet expansiuon of
// the string
func Phoneticise(s string) string {
	results := make([]string, 0)

	// remove leading and trailing whitespace
	s = strings.TrimSpace(s)
	// collapse duplicate whitespace
	spaceRegexp := regexp.MustCompile(`\s+`)
	s = spaceRegexp.ReplaceAllString(s, " ")

	var val string

	for _, c := range s {
		// mark spaces
		if c == ' ' {
			val = "(space)"
		} else {
			// grab the appropriate value and add to our results
			val = Lookup(c)
		}
		results = append(results, val)
	}

	// join results, trim any leading or trailing whitespace
	joined := strings.Join(results, " ")
	trimmed := strings.TrimSpace(joined)

	return trimmed
}

// Lookup takes a rune and returns the appropriate word from the phonetic alphabet
func Lookup(r rune) string {
	mapper := map[rune]string{
		'0': "Zero",
		'1': "One",
		'2': "Two",
		'3': "Three",
		'4': "Four",
		'5': "Five",
		'6': "Six",
		'7': "Seven",
		'8': "Eight",
		'9': "Nine",
		'A': "Alfa",
		'B': "Bravo",
		'C': "Charlie",
		'D': "Delta",
		'E': "Echo",
		'F': "Foxtrot",
		'G': "Golf",
		'H': "Hotel",
		'I': "India",
		'J': "Juliett",
		'K': "Kilo",
		'L': "Lima",
		'M': "Mike",
		'N': "November",
		'O': "Oscar",
		'P': "Papa",
		'Q': "Quebec",
		'R': "Romeo",
		'S': "Sierra",
		'T': "Tango",
		'U': "Uniform",
		'V': "Victor",
		'W': "Whiskey",
		'X': "X-ray",
		'Y': "Yankee",
		'Z': "Zulu",
	}

	val, ok := mapper[unicode.ToUpper(r)]
	if !ok {
		log.Fatalf("couldn't find %q", r)
	}

	return val
}

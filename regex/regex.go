package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	// Read in the word list
	data, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")

	// Compile a regular expression pattern to match the word list format
	r, _ := regexp.Compile(`::(.+?)::(.+)`)

	// Create a map to store the corrected words
	corrections := make(map[string]string)

	// Iterate over the lines in the word list
	for _, line := range lines {
		// Match the line against the regular expression pattern
		match := r.FindStringSubmatch(line)
		if match != nil {
			// If a match is found, add the misspelled word as the key and the corrected word as the value to the map
			corrections[match[1]] = match[2]
		}
	}

	// Use the map to autocorrect text
	text := "I always type airtbale instead of Airtable"
	for key, value := range corrections {
		text = strings.Replace(text, key, value, -1)
	}
	fmt.Println(text)
	// Output: "I always type Airtable instead of Airtable"
}

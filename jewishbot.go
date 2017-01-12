package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type GlossaryEntry struct {
	Transliterations []string `json:"transliterations"`
	Description      string   `json:"description"`
}

func StripString(in string) string {
	// convert all strings to lowercase
	in = strings.ToLower(in)
	// strip out nikkudot
	var out string
	for _, r := range in {
		if !(r >= '\u0591' && r < 'א') {
			out += string(r)
		}
	}
	return out
}

/*
A bot for /r/judaism
*/
func ReadGlossary(filename string) (map[string]GlossaryEntry, error) {
	var glossary = make(map[string]GlossaryEntry)
	if glossaryFile, err := os.Open(filename); err != nil {
		return nil, err
	} else {
		jsonParser := json.NewDecoder(glossaryFile)
		if err = jsonParser.Decode(&glossary); err != nil {
			return nil, err
		}
	}
	return glossary, nil
}

func main() {
	if glossary, err := ReadGlossary(os.Args[1]); err != nil {
		fmt.Printf("Error reading glossary: %s\n", err.Error())
	} else {
		for key, value := range glossary {
			fmt.Printf("%s:%s\n", key, value)
		}
	}
}

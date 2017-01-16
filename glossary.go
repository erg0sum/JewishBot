package main

import (
	"encoding/json"
	"os"
	"strings"
)

// GlossaryEntry is an entry in the Glossary
type GlossaryEntry struct {
	Transliterations []string `json:"transliterations"`
	Description      string   `json:"description"`
	MatchValues      []string `json:"-"`
}

//Glossary is the read in glossary
type Glossary map[string]GlossaryEntry

func stripString(in string) string {
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

// ReadGlossary reads a glossary file
// Glossary file has the format
// {
//	"term": {"transliterations":[], "description": "..."}
//}
func ReadGlossary(filename string) (Glossary, error) {
	var glossary = make(Glossary)
	glossaryFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(glossaryFile)

	if err = jsonParser.Decode(&glossary); err != nil {
		return nil, err
	}

	for key, entry := range glossary {
		var matchValues = make([]string, 0)
		matchValues = append(matchValues, stripString(key))
		for _, transliteration := range entry.Transliterations {
			matchValues = append(matchValues, stripString(transliteration))
		}
		entry.MatchValues = matchValues
	}

	return glossary, nil
}

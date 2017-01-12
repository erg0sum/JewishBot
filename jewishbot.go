package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlossaryEntry struct {
	Transliterations []string `json:"transliterations"`
	Description      string   `json:"description"`
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

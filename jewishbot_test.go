package main

import (
	"testing"
)

var testFile string = "test/glossary.json"

func TestReadGlossary(t *testing.T) {
	var glossary, err = ReadGlossary(testFile)
	if err != nil {
		t.Error("error should not be null")
	}
	if len(glossary) != 2 {
		t.Error("wrong length for glossary")
	}
	if val, ok := glossary["test"]; !ok {
		t.Error("test should be in glossary")
	} else if val != "A test string" {
		t.Error("Bad value for test string")
	}
}

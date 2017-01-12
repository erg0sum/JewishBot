package main

import "testing"

var testFile string = "test/glossary.json"

func TestReadGlossaryNoFile(t *testing.T) {
	var _, err = ReadGlossary("")
	if err == nil {
		t.Error("error should not be null")
	}
}

func TestReadGlossaryFromFile(t *testing.T) {
	var glossary, err = ReadGlossary(testFile)
	if err != nil {
		t.Log(err)
		t.Error("error should be null")
	}
	if len(glossary) != 3 {
		t.Error("wrong length for glossary")
	}
}

func TestReadEnglishEntry(t *testing.T) {
	var glossary, _ = ReadGlossary(testFile)
	if val, ok := glossary["test"]; !ok {
		t.Error("test should be in glossary")
		if val.Description != "A test string" {
			t.Error("Bad description for test")
		}

		if len(val.Transliterations) != 0 {
			t.Error("Transliterations for test string should be empty")
		}
	}
}

func TestReadHebrewEntry(t *testing.T) {
	var glossary, _ = ReadGlossary(testFile)
	if val, ok := glossary["שָׁלוֹם"]; !ok {
		t.Error("שָׁלוֹם should be in glossary")
	} else {
		if val.Description != "Peace" {
			t.Error("bad description for שָׁלוֹם")
		}

		if len(val.Transliterations) == 0 {
			t.Error("Transliterations for שָׁלוֹם should not be empty")
		}
	}
}

func TestReadHebrewWithMultipleTransliterations(t *testing.T) {

}

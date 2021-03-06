package main

import (
	"testing"
)

const testFile = "test/glossary.json"
const Shalom = "שָׁלוֹם"
const BeitMidrash = "בית מידרש"

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
	if val, ok := glossary["TEST"]; !ok {
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
	if val, ok := glossary[Shalom]; !ok {
		t.Error("שָׁלוֹם should be in glossary")
	} else {
		if val.Description != "Peace" {
			t.Error("bad description for שָׁלוֹם")
		}

		if len(val.Transliterations) == 0 {
			t.Error("Transliterations for שָׁלוֹם should not be empty")
		} else {
			if val.Transliterations[0] != "shalom" {
				t.Error("Could not find transliteration for shalom")
			}
		}
	}
}

func TestReadHebrewWithMultipleTransliterations(t *testing.T) {
	var glossary, _ = ReadGlossary(testFile)
	if val, ok := glossary[BeitMidrash]; !ok {
		t.Errorf("%s should be in glossary", BeitMidrash)
	} else {
		if len(val.Transliterations) != 4 {
			t.Errorf("Transliterations for %s should not be empty", BeitMidrash)
		}
	}
}

func TestStripStringWithHebrew(t *testing.T) {
	s := stripString(Shalom)
	if s != "שלום" {
		t.Error("Unable to strip nikkudot from shalom")
	}
}

func TestStripStringWithEnglish(t *testing.T) {
	s := stripString("TESTING123")
	if s != "testing123" {
		t.Error("Unable to lowercase string")
	}
}

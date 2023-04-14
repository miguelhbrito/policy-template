package main

import (
	"testing"

	"github.com/mailru/easyjson"
)

func TestParsingSettingsWithNoValueProvided(t *testing.T) {
	rawSettings := []byte(`{}`)
	settings := &Settings{}
	if err := easyjson.Unmarshal(rawSettings, settings); err != nil {
		t.Errorf("Unexpected error %+v", err)
	}

	if len(settings.DeniedNames) != 0 {
		t.Errorf("Expecpted DeniedNames to be empty")
	}

	valid, err := settings.Valid()
	if !valid {
		t.Errorf("Settings are reported as not valid")
	}
	if err != nil {
		t.Errorf("Unexpected error %+v", err)
	}
}

func TestIsPalindrome(t *testing.T) {
	settings := Settings{}

	if !settings.IsPalindrome("arara") {
		t.Errorf("string should be denied")
	}

	if settings.IsPalindrome("zelda") {
		t.Errorf("string should not be denied")
	}
}

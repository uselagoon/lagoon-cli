package app

import (
	"testing"
)

func TestFileExists(t *testing.T) {
	var got bool

	got = fileExists("fixtures/does_not_exist/.lagoon.yml")
	if got == true {
		t.Error("File should not have existed.", got)
	}
	got = fileExists("fixtures/basic/.lagoon.yml")
	if got == false {
		t.Error("File should have existed.", got)
	}
}

func TestFindLocalProjectRoot(t *testing.T) {
	path, err := findLocalProjectRoot("fixtures/basic/sub/sub/sub")
	if err != nil {
		t.Error("Should have found a Lagoon file")
	}
	if path != "fixtures/basic" {
		t.Error("Unexpected path", path)
	}

	_, err = findLocalProjectRoot("fixtures/nope/nope/nope/nope")
	if err == nil {
		t.Error("Should have errored as there is nothing")
	}
}

func TestGetProjectFromPath(t *testing.T) {
	app, err := getProjectFromPath("fixtures/basic/sub/dir")
	if err != nil {
		t.Error("Should have found a Lagoon file")
	}
	if app.Name != "basic-lagoon" {
		t.Error("App should have been named 'basic-lagoon'", app.Name)
	}
	if app.Dir != "fixtures/basic" {
		t.Error("App dir was not correct", app.Dir)
	}
}

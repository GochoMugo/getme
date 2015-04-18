package test

import (
	"github.com/GochoMugo/gimme/lib"
	"os"
	"path"
	"testing"
)

func TestGetTemplateDirectories(t *testing.T) {
	var paths []string
	key := "GIMME_PATH"

	pathVar := os.Getenv(key)
	defer os.Setenv(key, pathVar)

	os.Unsetenv(key)
	paths = lib.GetTemplateDirectories()
	if len(paths) != 0 {
		t.Errorf("empty slice not returned if env var %s not set", key)
	}

	os.Setenv(key, "dirA")
	paths = lib.GetTemplateDirectories()
	if paths[0] != "dirA" {
		t.Errorf("single path in env var %s", key)
	}

	os.Setenv(key, "dirA:dirB")
	paths = lib.GetTemplateDirectories()
	if len(paths) != 2 || paths[0] != "dirA" || paths[1] != "dirB" {
		t.Errorf("separator (:) not respected in env var %s", key)
	}
}

func TestFilterTemplateDirectories(t *testing.T) {
	target1 := "home/user/utils/go/lib"
	paths := []string{
		"/home/user/gimme",
		"/home/user/lib/gimme",
		target1,
	}

	var filtered []string
	test := func(searchString string, errorMessage string) {
		filtered = lib.FilterTemplateDiretories(paths, searchString)
		if filtered[0] != target1 {
			t.Errorf(errorMessage)
		}
	}

	test("go/lib", "straight-forward filtering")
	test("go/lib/", "right-slash used")
}

func TestSearchItem(t *testing.T) {
	cwd, _ := os.Getwd()
	dataDir := path.Join(cwd, "data")
	paths := []string{dataDir}
	var result string

	result = lib.SearchItem(paths, "sample.txt")
	if result != path.Join(dataDir, "sample.txt") {
		t.Errorf("searching for an exisiting file/dir")
	}

	result = lib.SearchItem(paths, "non-existing-file")
	if result != "" {
		t.Errorf("searching for non-exisitng file")
	}
}

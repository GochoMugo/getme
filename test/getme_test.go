package test

import (
	"github.com/GochoMugo/getme/lib"
	"os"
	"path"
	"testing"
)

func TestGetTemplateDirectories(t *testing.T) {
	var paths []string
	var err error
	key := "GETME_PATH"

	pathVar := os.Getenv(key)
	defer os.Setenv(key, pathVar)

	os.Setenv(key, "") // same as os.Unsetenv, but compatible with go 1.3
	paths, err = lib.GetTemplateDirectories()
	if len(paths) != 1 && err != nil {
		t.Errorf("slice with only homedir not returned if env var %s not set", key)
	}

	os.Setenv(key, "dirA")
	paths, err = lib.GetTemplateDirectories()
	if paths[1] != "dirA" && err != nil {
		t.Errorf("single path in env var %s", key)
	}

	os.Setenv(key, "dirA:dirB")
	paths, err = lib.GetTemplateDirectories()
	if paths[1] != "dirA" || paths[2] != "dirB" && err != nil {
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

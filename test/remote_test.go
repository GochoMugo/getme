package test

import (
	"bytes"
	"github.com/GochoMugo/getme/lib"
	"io/ioutil"
	"os"
	"path"
	"testing"
)


func TestFetch(t *testing.T) {
  actualUrl := "https://raw.githubusercontent.com/GochoMugo/getme/master/LICENSE"
  bogusUrl := "https://raw.githubusercontent.com/GochoMugo/getme/master/index.js"

  _, err := lib.Fetch(actualUrl)
  if err != nil {
    t.Errorf("Error dowloading from URL")
    return
  }

  _, err = lib.Fetch(bogusUrl)
  if err == nil {
    t.Errorf("Error not thrown on a bogus url")
  }
}

func TestDownloadFromGithub(t *testing.T) {
	shorthand := "GochoMugo/getme"
	branch := "master"
	filePath := "LICENSE"
	cwd, _ := os.Getwd()

	content, err := lib.DownloadFromGithub(shorthand, branch, filePath)
	if err != nil {
		t.Errorf("Error downloading from github: %s", err)
		return
	}
	realcontent, err := ioutil.ReadFile(path.Join(cwd, "..", "LICENSE"))
	if err != nil {
		t.Errorf("internal error: %s", err)
		return
	}
	if bytes.Compare(content, realcontent) != 0 {
		t.Errorf("content downloaded does not seem right")
		return
	}

  _, err = lib.DownloadFromGithub(shorthand, branch, "missing")
  if err == nil {
    t.Errorf("error is not thrown for missing file from remote")
  }
}


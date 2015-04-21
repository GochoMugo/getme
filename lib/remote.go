package lib

import (
  "errors"
	"io/ioutil"
	"net/http"
)

/**
* DownloadFromGithub retrieves a file from a github repo
 */
func DownloadFromGithub(shorthand, branch, filePath string) (content []byte, err error) {
	downloadUrl := "https://raw.githubusercontent.com/" + shorthand + "/" +
		branch + "/" + filePath
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
	  err = errors.New("not found")
	  return
	}
	content, err = ioutil.ReadAll(resp.Body)
	return
}


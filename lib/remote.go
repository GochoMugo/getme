package lib

import (
  "errors"
	"io/ioutil"
	"net/http"
)


/**
* Fetch is a convenience wrapper for fetching data from a Url
*/
func Fetch(url string) (content []byte, err error) {
	resp, err := http.Get(url)
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

/**
* DownloadFromGithub retrieves a file from a github repo
 */
func DownloadFromGithub(shorthand, branch, filePath string) ([]byte, error) {
	downloadUrl := "https://raw.githubusercontent.com/" + shorthand +
	  "/" + branch + "/" + filePath
  return Fetch(downloadUrl)
}

/**
* DownloadFromUrl retrives a file hosted at a URL
* This works exactly as Fetch. Only exists to main consistency in
* naming of remote-handling functions
*/
func DownloadFromUrl(url string) (content []byte, err error) {
  return Fetch(url)
}


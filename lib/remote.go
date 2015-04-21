package lib

import (
	"io/ioutil"
	"net/http"
)

/**
* DownloadFromGithub retrieves a file from a github repo
 */
func DownloadFromGithub(shorthand, branch, filepath string) (content []byte, err error) {
	downloadUrl := "https://raw.githubusercontent.com/" + shorthand + "/" +
		branch + "/" + filepath
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	return
}

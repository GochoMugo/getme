package lib

import (
	shutil "github.com/termie/go-shutil"
	"os"
	"path"
	"strings"
)

/**
* GetTemplateDirectories returns a slice of paths where templates
* may be placed in. Note: returns an empty slice if env var $GIMME_PATH
* is not set
 */
func GetTemplateDirectories() []string {
	pathVar := os.Getenv("GIMME_PATH")
	if pathVar == "" {
		return []string{}
	}
	return strings.Split(pathVar, ":")
}

/**
* FilterTemplateDiretories returns a slice of paths having the suffix
* passed
 */
func FilterTemplateDiretories(paths []string, suffix string) []string {
	var filtered []string
	suffix = strings.TrimRight(suffix, "/")
	for _, dir := range paths {
		if strings.HasSuffix(dir, suffix) {
			filtered = append(filtered, dir)
		}
	}
	return filtered
}

/**
* SearchItem looks for an item in a list of paths. Returns path to the
* first matched item. Returns empty string if not found.
 */
func SearchItem(paths []string, searchTerm string) string {
	var fullpath string
	for _, dir := range paths {
		fullpath = path.Join(dir, searchTerm)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath
		}
	}
	return ""
}

/**
* CopyItem copies file from source to destination
 */
func CopyItem(itemPath string, destination string) error {
	_, err := shutil.Copy(itemPath, destination, true)
	return err
}

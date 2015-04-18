package main

import (
	"github.com/GochoMugo/gimme/lib"
	"os"
)

const helpInfo = `
  gimme by GochoMugo <mugo@forfuture.co.ke>

  gimme                   show this help information
  gimme <file1> ...       give me those files

  notes:
    1. gimme relies on the environment variable, GIMME_PATH
    2. On *nix, $GIMME_PATH is delimited using : (full-colon)
    3. On Windows, GIMME_PATH is delimited using ; (semi-colon)
`

func main() {
	files := os.Args[1:]
	cwd, _ := os.Getwd()
	if len(files) == 0 {
		// show help information
		lib.LogRaw(helpInfo)
		return
	}
	for _, itemname := range os.Args[1:] {
		paths := lib.GetTemplateDirectories()
		if len(paths) == 0 {
			lib.LogError("not set: GIMME_PATH")
			return
		}
		itemPath := lib.SearchItem(paths, itemname)
		if itemPath == "" {
			lib.LogError("not found: " + itemname)
			continue
		}
		err := lib.CopyItem(itemPath, cwd)
		if err != nil {
			lib.LogError("not copied: " + itemname)
			continue
		}
		lib.LogSuccess("given you: " + itemname)
	}
}

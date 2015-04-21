package lib

import (
  "path/filepath"
  "os"
)

/**
* LocalGet gets a file from local file system
*/
func LocalGet(paths []string, itemname string, destination string) {
  itemPath := SearchItem(paths, itemname)
  if itemPath == "" {
    LogError(TextUnderline("not"), " found")
    return
  }
  err := CopyItem(itemPath, destination)
  if err != nil {
    LogError(TextUnderline("not"), " copied: ", TextBold(err.Error()))
    return
  }
  LogSuccess("given you ", TextBold(itemname))
}

/**
* LocalGetMany allows passing many files to get locally
*/
func LocalGetMany(itemnames []string) {
  cwd, _ := os.Getwd()
  paths, err := GetTemplateDirectories()
  if err != nil {
    LogError("could not get template directories: ", TextBold(err.Error()))
    return
  }
  for _, itemname := range itemnames {
    LocalGet(paths, itemname, cwd)
  }
}

/**
* RemoteGet gets a file from a remote source
*/
func RemoteGet(whichRemote string, args []string) {
  var content []byte
  var err error

  switch whichRemote {
  case "github":
    shorthand, branch, filePath := args[0], args[1], args[2]
    LogNormal("downloading from ", TextUnderline("Github"))
    content, err = DownloadFromGithub(shorthand, branch, filePath)
    if err != nil {
      LogError("failed to download from github: ", TextBold(err.Error()))
      return
    }
    filename := filepath.Base(filePath)
    err = WriteFile(filename, content)
    if err != nil {
      LogError("failed to write file: ", TextBold(err.Error()))
      return
    }
    LogSuccess("downloaded and written to ", TextBold(filename))
  }
}

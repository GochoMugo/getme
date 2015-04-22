package main

import (
	"github.com/GochoMugo/getme/lib"
	"github.com/spf13/cobra"
)

var getmeCmd = &cobra.Command{
  Use: "getme",
  Short: "getme is an easy way to retrieve files",
  Run: func(cmd *cobra.Command, args []string) {
    lib.LogNormal("nothing to do")
  },
}

var localCmd = &cobra.Command{
  Use: "local",
  Short: "use local file-system",
  Run: func(cmd *cobra.Command, args[] string) {
    lib.LocalGetMany(args)
  },
}

var remoteGithub bool
var remoteUrl bool
var remoteCmd = &cobra.Command{
  Use: "remote",
  Short: "use remote source",
  Run: func(cmd *cobra.Command, args []string) {
    var whichRemote string
    if remoteGithub {
      whichRemote = "github"
    }
    if remoteUrl {
      whichRemote = "url"
    }
    lib.RemoteGet(whichRemote, args)
  },
}


func main() {
  // adding flags
  remoteCmd.Flags().BoolVarP(&remoteGithub, "github", "g", false,
    "from github")
  remoteCmd.Flags().BoolVarP(&remoteUrl, "url", "u", false,
    "from an url")

  // adding commands
  getmeCmd.AddCommand(localCmd)
  getmeCmd.AddCommand(remoteCmd)
  getmeCmd.Execute()
}

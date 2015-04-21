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
var remoteCmd = &cobra.Command{
  Use: "remote",
  Short: "use remote source",
  Run: func(cmd *cobra.Command, args []string) {
    var whichRemote string
    if remoteGithub {
      whichRemote = "github"
    }
    lib.RemoteGet(whichRemote, args)
  },
}


func main() {
  // adding flags
  remoteCmd.Flags().BoolVarP(&remoteGithub, "github", "g", false,
    "from github")

  // adding commands
  getmeCmd.AddCommand(localCmd)
  getmeCmd.AddCommand(remoteCmd)
  getmeCmd.Execute()
}

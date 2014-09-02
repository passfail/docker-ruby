package cmdopts

import "flag"

var Link bool
var Verbose bool

func init(){
  // --link
  linkHelp := "link to current version"
  flag.BoolVar(&Link, "link", false, linkHelp)
  flag.BoolVar(&Link, "l", false, linkHelp)

  // --verbose
  verboseHelp := "be verbose"
  flag.BoolVar(&Verbose, "verbose", false, verboseHelp)
  flag.BoolVar(&Verbose, "v", false, verboseHelp)

  flag.Parse()
}

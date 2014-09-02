package main

import "os"
import "fmt"
import "github.com/passfail/docker-ruby/commands"

const version string = "0.0.1"

func main() {
  var command string
  var args []string

  if len(os.Args) >= 2 {
    command = os.Args[1]
  }

  if len(os.Args) >= 3 {
    args = os.Args[2:]
  }

  switch command {
    case "install":
      commands.Install(args)
    case "link":
      commands.Link(args)
    case "version":
      fmt.Println("docker-ruby", version)
    default:
      errAndExit("invalid command")
  }
}

func errAndExit(msg string){
  fmt.Println(msg)
  os.Exit(1)
}

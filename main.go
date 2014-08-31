package main

import "os"
import "fmt"
import "github.com/passfail/docker-ruby/commands"

const version string = "0.0.1"

func main() {
  command := os.Args[1]
  args := os.Args[2:]

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

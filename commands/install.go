package commands

import "fmt"
import "net/http"
import "io/ioutil"
import "os"

func Install(args []string){
  version := args[0]
  fmt.Println("Install Ruby", getVersionString(version))
}

func downloadPackage(version string){
  version = getVersionString(version)
}

func getVersionString(partial string) string {
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v.txt", partial)
  return string(getBody(url))
}

func getBody(url string) []byte {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  } else if resp.StatusCode >= 400 {
    fmt.Println(fmt.Sprintf("%s", resp.Status))
    os.Exit(1)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  }
  return body
}

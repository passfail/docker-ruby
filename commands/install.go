package commands

import "fmt"
import "os"
import "github.com/passfail/docker-ruby/lib/tarutils"
import "github.com/passfail/docker-ruby/lib/httputils"

func rubiesDir() string {
  var path string
  if len(os.Getenv("RUBIES_DIR")) > 0 {
    path = os.Getenv("RUBIES_DIR")
  } else {
    path = "./tmp/rubies"
  }
  var dirPerm os.FileMode
  dirPerm = 0755
  os.MkdirAll(path, dirPerm)
  return path
}

func Install(args []string){
  version := args[0]
  downloadPackage(version)
}

func downloadPackage(version string){
  version = getVersionString(version)
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v/binary", version)
  fmt.Println("Downloading and installing ruby", version)
  tarutils.UnarchiveGzip(httputils.GetBodyIO(url), rubiesDir())
}

func getVersionString(partial string) string {
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v.txt", partial)
  return string(httputils.GetBodyString(url))
}

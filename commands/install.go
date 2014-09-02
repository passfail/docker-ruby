package commands

import "strings"
import "fmt"
import "log"
import "os"
import "os/exec"
import "github.com/passfail/docker-ruby/lib/tarutils"
import "github.com/passfail/docker-ruby/lib/httputils"

func Install(args []string){
  version := args[0]
  downloadPackage(version)
}

func downloadPackage(version string){
  version = getVersionString(version)
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v/binary", version)
  log.Println("Downloading and installing ruby", version)
  tarutils.UnarchiveGzip(httputils.GetBodyIO(url), versionsDir())
  installDependecies(version)
}

func getVersionString(partial string) string {
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v.txt", partial)
  return string(httputils.GetBodyString(url))
}

func installDependecies(version string){
  path := strings.Join([]string{versionsDir(), version, "bin", "install-ruby-dependencies"}, "/")
  cmd := exec.Command(path)
  err := cmd.Start()
  if err != nil { log.Fatal(err) }
  log.Println("Installing Dependencies...")
  cmd.Wait()
}

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

func versionsDir() string {
  var dirPerm os.FileMode
  dirPerm = 0755
  dir := strings.Join([]string{rubiesDir(), "versions"}, "/")
  os.MkdirAll(dir, dirPerm)
  return dir
}

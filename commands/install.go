package commands

import "fmt"
import "net/http"
import "io/ioutil"
import "io"
import "os"
import "archive/tar"
import "compress/gzip"
import "log"
import "bytes"

func Install(args []string){
  version := args[0]
  // fmt.Println("Install Ruby", getVersionString(version))
  downloadPackage(version)
}

func downloadPackage(version string){
  version = getVersionString(version)
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v/binary", version)
  fmt.Println("Downloading and installing ruby", version)
  tarUnarchive(gzipExtract(getBodyIO(url)))
}

func tarUnarchive(buffer *bytes.Buffer, dest string){
    // Read the tar
    tr := tar.NewReader(buffer)
    for {
      hdr, err := tr.Next()
      if err == io.EOF {
        break
      }
      if err != nil {
        log.Fatalln(err)
      }
      // fmt.Println(hdr.Typeflag, hdr.Name)
      switch {
        case hdr.FileInfo().IsDir():
          fmt.Println("Dir:", hdr.Name)
        default:
          fmt.Println("File:", hdr.Name)
      }
    }
}

func gzipExtract(ioVar io.ReadCloser) *bytes.Buffer {
  defer ioVar.Close()
  gr, err := gzip.NewReader(ioVar)
  defer gr.Close()
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  }
  outBytes, err := ioutil.ReadAll(gr)
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  }
  return bytes.NewBuffer(outBytes)
}

func getVersionString(partial string) string {
  url := fmt.Sprintf("http://ruby.itri.ag/versions/%v.txt", partial)
  return string(getBodyString(url))
}

func getBodyIO(url string) io.ReadCloser {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  } else if resp.StatusCode >= 400 {
    fmt.Println(fmt.Sprintf("%s", resp.Status))
    os.Exit(1)
  }
  return resp.Body
}

func getBodyString(url string) string {
  ioBody := getBodyIO(url)
  defer ioBody.Close()
  body, err := ioutil.ReadAll(ioBody)
  if err != nil {
    fmt.Println(fmt.Sprintf("%s", err))
    os.Exit(1)
  }
  return string(body)
}

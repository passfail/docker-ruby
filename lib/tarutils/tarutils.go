package tarutils

import "bytes"
import "io"
import "log"
import "fmt"
import "io/ioutil"
import "os"
import "archive/tar"
import "compress/gzip"
import "strings"

func Unarchive(ioVar io.ReadCloser, dest string) {
  outBytes, err := ioutil.ReadAll(ioVar)
  if err != nil { log.Fatalln(err) }
  unarchive(bytes.NewBuffer(outBytes), dest)
}

func UnarchiveGzip(ioVar io.ReadCloser, dest string) {
  unarchive(extractGzip(ioVar), dest)
}

func unarchive(buffer *bytes.Buffer, dest string){
    wrkdir, err := os.Getwd()
    if err != nil { log.Fatalln(err) }
    os.Chdir(dest)

    // Read the tar
    tr := tar.NewReader(buffer)
    for {
      hdr, err := tr.Next()
      if err == io.EOF {
        break
      } else if err != nil {
        log.Fatalln(err)
      }

      // Create file
      fullpath := strings.Join([]string{dest, hdr.Name}, "/")
      if hdr.FileInfo().IsDir()  {
        fmt.Println("dir: ", fullpath)
        err := os.Mkdir(hdr.Name, hdr.FileInfo().Mode())
        if err != nil && !os.IsExist(err) { log.Fatalln(err) }
      } else {
        fmt.Println("file:", fullpath)
        file, err := os.Create(hdr.Name)
        if err != nil { log.Fatalln(err) }

        // Copy Data
        io.Copy(file, tr)

        // Set Permissions
        err = file.Chmod(hdr.FileInfo().Mode())
        if err != nil { log.Fatalln(err) }

        // Close the File
        err = file.Close()
        if err != nil { log.Fatalln(err) }
      }

    }
    os.Chdir(wrkdir)
}

func extractGzip(ioVar io.ReadCloser) *bytes.Buffer {
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

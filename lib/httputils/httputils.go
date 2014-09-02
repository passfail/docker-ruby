package httputils

import "net/http"
import "io/ioutil"
import "io"
import "log"

func GetBodyIO(url string) io.ReadCloser {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatalln(err)
  } else if resp.StatusCode >= 400 {
    log.Fatalln(resp.Status)
  }
  return resp.Body
}

func GetBodyString(url string) string {
  ioBody := GetBodyIO(url)
  defer ioBody.Close()
  body, err := ioutil.ReadAll(ioBody)
  if err != nil { log.Fatalln(err) }
  return string(body)
}

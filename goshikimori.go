package goshikimori

import (
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "strings"
)

const (
  bearer       = "Bearer "
  urlOrig      = "%s://%s"
  protocol     = "https"
  urlShiki     = "shikimori.one/api/"
)

func Parameters(s ...string) string {
  p := strings.Join(s, "/")
  return p
}

func NewRequest(oauth2, access_token, method, input string) ([]byte, error) {
  req, err := http.NewRequest(
    method, fmt.Sprintf(urlOrig, protocol, urlShiki)+input, nil)
  req.Header.Add("User-Agent", oauth2)
  req.Header.Add("Authorization", bearer+access_token)
  if err != nil {
    log.Fatal(err)
  }

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}
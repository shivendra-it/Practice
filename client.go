package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
  "log"
)
func main() {

	url := "http://localhost:8080/"
  var username string = "john"
  var passwd string = "cena"
  client := &http.Client{}
  req, err := http.NewRequest("GET",url, nil)
  req.SetBasicAuth(username, passwd)
  resp, err := client.Do(req)
  if err != nil{
      log.Fatal(err)
  }
  bodyText, err := ioutil.ReadAll(resp.Body)
  s := string(bodyText)
  fmt.Println(s)
  fmt.Println(resp.StatusCode)
}

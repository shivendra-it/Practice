package main
import (
 "encoding/base64"
 "net/http"
 "fmt"
 //"strings"
)
//var w http.ResponseWriter
var s string
func handler(w http.ResponseWriter, r *http.Request){
// fmt.Fprintf(w, "I am here %s!\n", r.URL.Path[1:])
 vat := r.Header.Get("authorization")
//Encoded data
 fmt.Println("\nEncoded username and password\n");
 fmt.Println(vat)
 data, err := base64.StdEncoding.DecodeString(vat[6:])
 if err != nil {
  fmt.Println("error:", err)
  return
 }
//Decoded data
 fmt.Println("\nDecoded username and password\n");
 fmt.Printf("%q\n", data)
 s = string(data[:])
}


func Auth() string{
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
  return s//,w
}

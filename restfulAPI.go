//Importing main Package 
package main
//Importing all other needed packages
import (
 "encoding/base64"
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
)
//Request Handler
func handler(w http.ResponseWriter, r *http.Request) {
 fmt.Fprintf(w, "I am here %s!\n", r.URL.Path[1:])
 vat := r.Header.Get("authorization")
 
 fmt.Println("\nEncoded username and password\n");
 fmt.Println(vat)

 data, err := base64.StdEncoding.DecodeString(vat[6:])
 if err != nil {
  fmt.Println("error:", err)
  return
 }

 fmt.Println("\nDecoded username and password\n");
 fmt.Printf("%q\n", data)

 robots, err := ioutil.ReadAll(r.Body)
 r.Body.Close()
 if err != nil {
  panic("there is an error")
 }
 if robots != nil {
  fmt.Println("\nJSON Data is given below\n");
  fmt.Printf("%s", robots)
  var dat map[string]interface{}
  if err := json.Unmarshal(robots, &dat); err != nil {
   panic(err)
  }
  fmt.Printf("\n")
  fmt.Println("\nMap of JSON Data is given below\n");
  fmt.Println(dat)
  datab, _ := json.Marshal(dat)
  fmt.Fprintln(w, string(datab))
 }

}
//main Function
func main() {
 http.HandleFunc("/", handler)
 http.ListenAndServe(":8080", nil)
}

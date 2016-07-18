//Importing main Package
package main
//Importing all other needed packages
import (
 "encoding/base64"
 "fmt"
 "net/http"
 "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "log"
 "strings"
)
const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "authentication"
    DB_USER = "root"
    DB_PASS = "12345"
)
//Request Handler
func handler(w http.ResponseWriter, r *http.Request) {
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

 s := string(data[:])
 st := strings.Split(s, ":")
 st1, st2 := st[0], st[1]

 dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    var str string
rows, err := db.Query("select userpass from T")
if err != nil {
	log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&str)
	if err != nil {
		log.Fatal(err)
	}
//	log.Println(str)
	stp := strings.Split(str, ":")
  stp1, stp2 := stp[0], stp[1]
if st1==stp1{
    if st2==stp2 {
	fmt.Println("\nPassword Matched\n")
	http.Error(w, `Successfully login`, http.StatusOK)
	return
    }else{
	fmt.Println("\nPassword didn't Matched\n")
	http.Error(w, `Invalid input parameters!`, http.StatusUnauthorized)
	return
}
}
}
  fmt.Println("\nNo such user exist\n")
	http.Error(w, `No user exist`, http.StatusBadRequest)
	return
err = rows.Err()
if err != nil {
	log.Fatal(err)
}

}
//main Function
func main() {
 http.HandleFunc("/", handler)
 http.ListenAndServe(":8080", nil)
}

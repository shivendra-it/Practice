//Importing main Package
package main
//Importing all other needed packages
import (
  "fmt"
  "log"
  "strings"
//  "net/http"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)
const (
    DB_HOST = "tcp(127.0.0.1:3306)"
    DB_NAME = "authentication"
    DB_USER = "root"
    DB_PASS = "12345"
)
//Request Handler

//main Function
func oldDB() {
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
      s := Auth()
      st := strings.Split(s, ":")
      st1, st2 := st[0], st[1]
      stp := strings.Split(str, ":")
      stp1, stp2 := stp[0], stp[1]
    if st1==stp1{
        if st2==stp2 {
      fmt.Println("\nPassword Matched\n")
    //  http.Error(w, `Successfully login`, http.StatusOK)
      return
        }else{
      fmt.Println("\nPassword didn't Matched\n")
  //    http.Error(w, `Invalid input parameters!`, http.StatusUnauthorized)
      return
    }
    } else {

      fmt.Println("\nNo such user exist\n")
    //  http.Error(w, `No user exist`, http.StatusBadRequest)
      return
    err = rows.Err()
    if err != nil {
      log.Fatal(err)
    }
      stmtIns, err := db.Prepare("INSERT INTO T VALUES(?)") // ? = placeholder
          if err != nil {
              panic(err.Error()) // proper error handling instead of panic in your app
          }
          _, err = stmtIns.Exec(s) // Insert tuples (i, i^2)
                  if err != nil {
                      panic(err.Error()) // proper error handling instead of panic in your app
                  }
                }

}
}

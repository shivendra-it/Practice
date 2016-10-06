package main


import (
    "os"
    "fmt"
    "log"
    "encoding/json"
    "io"
    "net/http"
    "strconv"
    "time"
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    // Read the body into a string for json decoding
    var content = &PayloadCollection{}

    err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Go through each payload and queue items individually to be posted to S3
    for _, payload := range content.Payloads {

        // let's create a job with the payload
        work := Job{Payload: payload}

        // Push the work onto the queue.
        JobQueue <- work
    }

    w.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/", payloadHandler)
    err := http.ListenAndServe(":8080", nil)
    log.Println("listening on localhost:8080")
    fmt.Println(err)
}

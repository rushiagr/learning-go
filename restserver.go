package main

// From: http://thenewstack.io/make-a-restful-json-api-go/

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    // http.HandleFunc() will automatically register the request handler. So
    // it is not required to provide this as an argument to
    // http.ListenAndServe()
    // 'w' is the response object to write to
    // 'r' is the request object
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // fprintf writes to the object passed as first param, in this case 'w'
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

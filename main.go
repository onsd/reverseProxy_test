package main

import(
    "fmt"
    "net/http"
)

func main() {
    // register function
    // Hello
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9091", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World at 9091!, I FIXED")
}

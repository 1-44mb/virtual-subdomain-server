package main

/**
 * [ ] Wildcard subdomain routing from any subdomain
 * [ ] Lookup to check if subdomain is in use, return default display if not
 * [ ] If subdomain is in use and request is for something that doesn't exist return 404
 */

import (
  "fmt"
  "net/http"
  "strings"
)

type Mux struct {
  http.Handler
}

func handler(w http.ResponseWriter, r *http.Request) {
  domainParts := strings.Split(r.Host, ".")
  if len(domainParts) > 2 {
    fmt.Fprintf(w, "Hello, %q", domainParts[0:1])
  } else {
    fmt.Fprintf(w, "Hello, %q", r.Host)
  }
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  mux.ServeHTTP(w, r)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)
  http.ListenAndServe(":8080", mux)
}

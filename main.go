package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"web2smb"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8445", nil))
}

// this handler is returning component path of URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	tp := &Smb{}
	tp.Datafs = strings.TrimPrefix(r.URL.Path, "/")
	go web2smb.SmbGo(tp.Datafs)
}

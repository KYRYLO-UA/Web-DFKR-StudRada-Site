package main

import (
	"flag"
	"net/http"
)

func main() {
	listenAddr := flag.String("listenaddr", "8080", "todo")
	flag.Parse()

	http.ListenAndServe(*listenAddr, nil)
}

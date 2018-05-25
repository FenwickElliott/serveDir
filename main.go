package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", "80", "port to serve on")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

package others

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func Server_without_gin() { //https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //r is from where we get request details sent by user, w is where we send our response
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)) // fprintf(whereToWrite, whatToWrite)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

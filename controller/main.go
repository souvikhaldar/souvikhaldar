package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./view")))

	debug := flag.Bool("debug", false, "Set debug mode")
	flag.Parse()
	if !*debug {
		// debug mode off
		log.Println("Serving on port 443")
		log.Fatal(http.ListenAndServeTLS(
			":443",
			"/etc/letsencrypt/live/souvikhaldar.in/fullchain.pem",
			"/etc/letsencrypt/live/souvikhaldar.in/privkey.pem",
			nil,
		))
	} else {
		// debug mode on
		log.Println("Serving on port 8192")
		log.Fatal(http.ListenAndServe(":8192", nil))
	}
}

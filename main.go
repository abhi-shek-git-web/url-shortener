package main

import (
	urlshortner "URL_Shoner/urlShortner"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// configuring roouter
	router := mux.NewRouter()


	// making route
	router.HandleFunc("/urlShortner", urlshortner.UrlShortner).Methods(http.MethodPost)
	router.HandleFunc("/redirect/{shortUrl}", urlshortner.Redirect).Methods(http.MethodGet)
	router.HandleFunc("/metrics", urlshortner.Metrics).Methods(http.MethodGet)

	// making machine up and listening
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Print("error occured during listen and serve. error =", err)
		return
	}

}

package urlshortner

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"net/http"
)


func UrlShortner(w http.ResponseWriter, r *http.Request) {

	// consuming input and validating it
	var url struct {
		Url string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Print("error occured during decoding input into url structure. error =", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// parsing url
	parsedUrl, err := r.URL.Parse(url.Url)
	if err != nil {
		log.Print("error occured during parsing url. error =", err)
		http.Error(w, "url not allowed", http.StatusBadRequest)
		return
	}

	// extracting hostname
	host := parsedUrl.Hostname()
	if host == "" {
		log.Print("error occured during extracting hostname, error =", err)
		http.Error(w, "url not allowed", http.StatusBadRequest)
		return
	}

	//  checking if host is valid
	_, err = net.LookupHost(host)
	if err != nil {
		log.Print("error occured during checking host. error =", err)
		http.Error(w, "url host not allowed", http.StatusBadRequest)
		return
	}

	// checking url path or scheme
	switch {
	case parsedUrl.Path == "":
		log.Print("empty path not allowed")
		http.Error(w, "empty path not allowed", http.StatusBadRequest)
		return

	case parsedUrl.Scheme == "":
		log.Print("url scheme not allowed")
		http.Error(w, "url scheme not allowed", http.StatusBadRequest)
		return
	}

	

	if len(url.Url) <= 8 {
		_, err := w.Write([]byte("url is already short"))
		log.Print("error occured during sending response. error =", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// processing input
	// generate short url
	shortUrl := generateShortUrl(url.Url)

	// sending response back
	_, err = w.Write([]byte(shortUrl))
	if err != nil {
		log.Print("error occured during writing response back to user. error =", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

}


func generateShortUrl(originalUrl string) string {
	// generating hash
	h := sha256.Sum256([]byte(originalUrl))
	hashStr := base64.URLEncoding.EncodeToString(h[:])
	return hashStr[:8]
}

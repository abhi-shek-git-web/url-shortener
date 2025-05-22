package urlshortner

import (
	"URL_Shoner/models"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/url"
	"sort"
	"sync"
)

var (
	savedUrl      = make(map[string]models.Url)
	DomainCounter = make(map[string]int)
	mu            sync.Mutex
)

func UrlShortner(w http.ResponseWriter, r *http.Request) {

	// consuming input and validating it
	var urlInput struct {
		Url string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&urlInput)
	if err != nil {
		log.Print("error occured during decoding input into url structure. error =", err)
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// checking if url is already short
	if len(urlInput.Url) <= 8 {
		log.Print("url is already short")
		http.Error(w, "url is already short", http.StatusBadRequest)
		return
	}

	// parsing url
	parsedUrl, err := url.ParseRequestURI(urlInput.Url)
	if err != nil {
		log.Printf("error occured during parsing url. error = %v", err)
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	// extracting host
	host := parsedUrl.Hostname()
	if host == "" {
		log.Print("host is empty")
		http.Error(w, "empty host", http.StatusBadRequest)
		return
	}

	// checking if url syntax is valid
	_, err = net.LookupHost(host)
	if err != nil {
		log.Print("error occured during checking host. error =", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// validating url path
	if parsedUrl.Path == "" {
		log.Print("empty path")
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	// processing input
	// generate short url
	shortUrl := generateShortUrl(urlInput.Url)
	if shortUrl == "" {
		log.Print("short url is nil")
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

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
	hash := sha256.Sum256([]byte(originalUrl))
	hashStr := base64.URLEncoding.EncodeToString(hash[:])
	shortenedUrl := hashStr[:8]

	mu.Lock()

	// saving origional url and short url for further access purpose
	savedUrl[shortenedUrl] = models.Url{
		Url:      originalUrl,
		ShortUrl: shortenedUrl,
	}
	mu.Unlock()

	// updating domain counter
	parsedUrl, err := url.Parse(originalUrl)
	if err != nil {
		log.Print("error occured during parsing url", err)
		return ""
	}

	host := parsedUrl.Hostname()
	DomainCounter[host]++

	return shortenedUrl
}

func Redirect(w http.ResponseWriter, r *http.Request) {

	// checking if url input is correct
	if len(r.URL.Path) < 18 {
		http.Error(w, "input url is not correct", http.StatusBadRequest)
		return
	}

	// extracting short url from incoming request
	shortUrl := r.URL.Path[len("/redirect/"):]

	// find origional url corresponding to given short url

	mu.Lock()

	// find origional url from saved urls
	savedUrl := savedUrl[shortUrl]

	mu.Unlock()

	if savedUrl.Url == "" {
		http.Error(w, "url not found", http.StatusBadRequest)
		return
	}

	// redirect to origional url
	http.Redirect(w, r, savedUrl.Url, http.StatusFound)
}

func Metrics(w http.ResponseWriter, r *http.Request) {

	// place all data into a slice
	type domainStats struct {
		Domain string
		Count  int
	}

	var stats []domainStats

	mu.Lock()

	for dom, cnt := range DomainCounter {
		stats = append(stats, domainStats{Domain: dom, Count: cnt})
	}

	mu.Unlock()

	// sort slice
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count > stats[j].Count
	})

	// extract top 3 domain
	if len(stats) > 3 {
		stats = stats[:3]
	}

	// send response

	err := json.NewEncoder(w).Encode(stats)
	if err != nil {
		log.Print("error occured during sending response. error = ", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

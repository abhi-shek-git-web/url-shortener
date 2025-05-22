package urlshortner_test

import (
	urlshortner "URL_Shoner/urlShortner"
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_urlShortnerOk(t *testing.T) {

	body := `{
		"url":"http://youtube.com/"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v ", err)
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusOK {
		log.Fatalf("want response code %d and got %d", http.StatusOK, res.Code)
	}
}

func Test_urlShortnerEmptyBody(t *testing.T) {

	body := ``

	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusBadRequest {
		t.Fatalf("want code %d and got %d ", http.StatusBadRequest, res.Code)
	}
}

func Test_urlShortnerWrongUrlHost(t *testing.T) {

	body := `{
		"url":"http://example/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("want code %d and got %d", http.StatusBadRequest, res.Code)
	}
}

func Test_urlShortnerWrongUrlScheme(t *testing.T) {

	body := `{
		"url":"123://example.com/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("want code %d and got %d", http.StatusBadRequest, res.Code)
	}
}

func Test_urlShortnerEmptyPath(t *testing.T) {

	body := `{
		"url":"http://example.com"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("want code %d and got %d", http.StatusBadRequest, res.Code)
	}
}

func Test_urlShortnerWrongJsonMapping(t *testing.T) {

	body := `{
		"ur":"http://example.com/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)

	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("want code %d and got %d", http.StatusBadRequest, res.Code)
	}
}

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
		"url":"http://example.com/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusOK{
		t.Fail()
	}
}

func Test_urlShortnerEmptyBody(t *testing.T) {

	body := ``

	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusBadRequest{
		t.Fail()
	}
}

func Test_urlShortnerWrongUrlHost(t *testing.T) {

	body := `{
		"url":"http://example/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusBadRequest{
		t.Fail()
	}
}

func Test_urlShortnerWrongUrlScheme(t *testing.T) {

	body := `{
		"url":"123://example.com/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusBadRequest{
		t.Fail()
	}
}

func Test_urlShortnerEmptyPath(t *testing.T) {

	body := `{
		"url":"http://example.com"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusBadRequest{
		t.Fail()
	}
}

func Test_urlShortnerWrongJsonMapping(t *testing.T) {

	body := `{
		"ur":"http://example.com/abc"
	}`
	reqBody := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(http.MethodPost, "/urlShortner", reqBody)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	urlshortner.UrlShortner(res, req)
	if res.Code!= http.StatusBadRequest{
		t.Fail()
	}
}
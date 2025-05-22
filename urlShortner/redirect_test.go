package urlshortner

import (
	"URL_Shoner/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RedirectOk(t *testing.T) {
	url := "http://localhost/abcd"
	shortUrl := "O4WVmCTH"
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirect/O4WVmCTH", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusOK {
		log.Fatalf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
	}
}

func Test_RedirectNilInput(t *testing.T) {
	url := "http://localhost/abcd"
	shortUrl := ""
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirect/", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
	}
}
func Test_RedirectNilWrongInput(t *testing.T) {
	url := "http://localhost/abcd"
	shortUrl := "yu098"
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirect/O4WVmjklmn", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
	}
}
func Test_RedirectNilWrongPath(t *testing.T) {
	url := "http://localhost/abcd"
	shortUrl := "O4WVmCTH"
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirectO4WVmCTH", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
	}
}
func Test_RedirectNilUrl(t *testing.T) {
	url := ""
	shortUrl := "O4WVmCTH"
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirect/O4WVmCTH", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v", err)
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Fatalf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
	}
}

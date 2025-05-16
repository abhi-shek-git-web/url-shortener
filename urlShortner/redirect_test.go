package urlshortner

import (
	"URL_Shoner/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Redirect(t *testing.T) {
	url := "http://localhost/abcd"
	shortUrl := "O4WVmCTH"
	savedUrl[shortUrl] = models.Url{
		Url:      url,
		ShortUrl: shortUrl,
	}

	req, err := http.NewRequest(http.MethodGet, "/redirect/O4WVmCTH", nil)
	if err != nil {
		log.Print("error occured during making req. error =", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusOK {
		log.Printf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
		t.Fail()
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
		log.Print("error occured during making req. error =", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Printf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
		t.Fail()
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
		log.Print("error occured during making req. error =", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Printf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
		t.Fail()
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
		log.Print("error occured during making req. error =", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Printf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
		t.Fail()
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
		log.Print("error occured during making req. error =", err)
		t.Fail()
	}
	res := httptest.NewRecorder()
	Redirect(res, req)
	if res.Code != http.StatusBadRequest {
		log.Printf("response status code doesn't match. Want %d but got %d", http.StatusOK, res.Code)
		t.Fail()
	}
}

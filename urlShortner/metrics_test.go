package urlshortner

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Metrics(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/metrics", nil)
	if err != nil {
		log.Print("error occured during making req. error = ", err)
		t.Fail()
	}
	res := httptest.NewRecorder()

	Metrics(res, req)
	if res.Code != http.StatusOK {
		log.Printf("want status code %d and get %d", http.StatusOK, res.Code)
		t.Fail()
	}
}

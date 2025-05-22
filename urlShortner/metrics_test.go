package urlshortner

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Metrics(t *testing.T) {

	DomainCounter = map[string]int{
		"example.com":  10,
		"google.com":   25,
		"youtube.com":  20,
		"linkedin.com": 15,
	}

	req, err := http.NewRequest(http.MethodGet, "/metrics", nil)
	if err != nil {
		log.Fatalf("error occured during making req. error = %v ", err)
	}
	res := httptest.NewRecorder()

	Metrics(res, req)
	if res.Code != http.StatusOK {
		log.Fatalf("want status code %d and get %d", http.StatusOK, res.Code)
	}

	type domainCount struct {
		Domain string `json:"domain"`
		Count  int    `json:"count"`
	}

	var result []domainCount
	err = json.Unmarshal(res.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("error occured during unmarshalling data. error = %v", err)
	}

	expected := map[string]int{
		"google.com":   25,
		"youtube.com":  20,
		"linkedin.com": 15,
	}

	if len(result) != 3 {
		t.Fatalf("expected len of result 3, but got %d", len(result))
	}

	for dom, cnt := range expected {
		found := false
		for _, rslt := range result {
			if rslt.Domain == dom {
				if rslt.Count != cnt {
					t.Fatalf("expected %s to have count %d, but got %d", dom, cnt, rslt.Count)
				}
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("expected domain %s not found in result", dom)
		}
	}

}

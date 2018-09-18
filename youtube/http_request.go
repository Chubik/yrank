package youtube

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func httpRequest(url string) ([]byte, int, error) {

	var httpClient = &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Cannont prepare the HTTP request", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Cannot process the HTTP request", err)
	}

	defer resp.Body.Close()
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal("Cannont dump the body of HTTP response", err)
	}
	statusProcessing(resp.StatusCode, url)
	return body, resp.StatusCode, err
}

func statusProcessing(statusCode int, url string) {
	if statusCode == 403 {
		log.Fatalf("Looks like the rate limit is exceeded, please try again later")
	} else if statusCode == 202 {
		log.Printf("Looks like the server need some time to prepare request.")
	} else if statusCode != 200 {
		log.Fatalf("The status code of URL %s is not OK : %d", url, statusCode)
	}
}
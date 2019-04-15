package requester

import (
	"io/ioutil"
	"net/http"
)

type HttpRequester struct {
	url    string
	client *http.Client
}

func NewHttpRequester(u string) *HttpRequester {
	client := &http.Client{}
	return &HttpRequester{
		client: client,
		url:    u,
	}
}

// make a single HTTP request to the url. return the response body string.
func (h *HttpRequester) OneRequest() (string, error) {
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		return "", err
	}

	// TODO -- allow custom headers

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

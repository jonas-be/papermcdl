package paper_api

import (
	"io"
	"net/http"
)

type PapermcAPI struct {
	URL string
}

func (p PapermcAPI) sendRequest(url string) ([]byte, error) {
	resp, err := http.Get(p.URL + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

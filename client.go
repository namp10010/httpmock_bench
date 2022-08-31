package httpmock_bench

import (
	"io"
	"net/http"
)

var serverURL = "http://httpmock"

func get(url string) (string, error) {
	var u string

	if url != "" {
		u = url
	} else {
		u = serverURL
	}

	res, err := http.Get(u)

	if err != nil {
		return "", nil
	}

	defer res.Body.Close()

	s, err := io.ReadAll(res.Body)

	return string(s), err
}

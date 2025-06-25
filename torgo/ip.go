package torgo

import (
	"fmt"
	"io"
	"net/http"
)

func ShowIP() (string, error) {
	const url = "https://api.ipify.org"
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s responded with status code of :- %d %s", url, resp.StatusCode, resp.Status)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

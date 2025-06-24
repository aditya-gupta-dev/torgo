package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ShowIP() {
	const url = "https://api.ipify.org"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Check your network connection.")
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("%s responded with status code of :- %d %s", url, resp.StatusCode, resp.Status))
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\033[32m%s\033[0m", string(data))
	os.Exit(0)
}

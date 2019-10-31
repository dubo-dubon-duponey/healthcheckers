package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	uri := fmt.Sprintf("%s", os.Getenv("HEALTHCHECK_URL"))
	if uri == "" {
		os.Exit(1)
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(2)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res.StatusCode)
	fmt.Println()
	fmt.Println(string(body))
}

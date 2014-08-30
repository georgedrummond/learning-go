package pages

import (
	"io/ioutil"
	"net/http"
)

func Page() string {
	resp, err := http.Get("http://www.google.co.jp")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// do something
	}

	return string(body)
}

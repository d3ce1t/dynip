package dynip

import (
	"io/ioutil"
	"net/http"
)

// https://api.ipify.org/
func currentServiceIP() string {
	resp, err := http.Get("https://api.ipify.org/")
	manageError(err)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	manageError(err)
	return string(result)
}

func manageError(err error) {
	if err != nil {
		panic(err)
	}
}

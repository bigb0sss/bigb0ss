package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // Closing the body is important
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}
	return string(body)

}

func main() {
	fmt.Println(Get("https://google.com/robots.txt"))
}

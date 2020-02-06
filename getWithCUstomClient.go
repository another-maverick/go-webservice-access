package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	customClient := http.Client{Timeout: time.Second * 20}

	resp, _ := customClient.Get("http://google.com")

	data, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	fmt.Println(string(data))

}

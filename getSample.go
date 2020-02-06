package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp , _ := http.Get("http://adobe.com")
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(data))
}

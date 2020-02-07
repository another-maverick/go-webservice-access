package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MyError struct {
	HTTPCode int `json: "-"`
	MyCode int `json: "code, omitempty"`
	Message string `json: "message"`
}

// above struct implements error interface by defining Error method
func (e MyError) Error() string {
	return fmt.Sprintf("HTTPCode - %d, my-code - %d, Message - %s \n", e.HTTPCode, e.MyCode, e.Message)
}

func customGet(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	if res.StatusCode < 200 || res.StatusCode > 300 {
		if res.Header.Get("Content-Type") != "application/json" {
			return res, fmt.Errorf("content type is not application/json. cannot proceeed further, status is %s", res.Status)
		}

		resData, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		errData := struct {
			WrapperError MyError `json: "error"`
		}{}
		fmt.Println(string(resData))
		err := json.Unmarshal(resData, &errData)
		if err != nil {
			fmt.Println(err)
			return res, fmt.Errorf("Unable to parse the json. Error is  - %s", err)
		}
		errData.WrapperError.HTTPCode = res.StatusCode

		return res, errData.WrapperError
	}
	return res, nil
}

func main() {
	res, err := customGet("http://localhost:12345/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respData, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(respData))

}
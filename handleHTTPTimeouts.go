package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	tmpFile, err := os.Create("/tmp/tmpFile.dmg")
	if err != nil  {
		fmt.Println("cannot create the file for writing")
	}
	defer tmpFile.Close()

	url := "https://download-installer.cdn.mozilla.net/pub/firefox/releases/40.0.3/mac/en-US/Firefox%2040.0.3.dmg"

	err = downloadFile(url, tmpFile, 10)

	if err != nil {
		fmt.Println(err)
	}

	tmpFileStat, err := tmpFile.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v bytes have been doanloaded and filename is %v", tmpFileStat.Size(), tmpFileStat.Name())

}

func downloadFile(url string, tmpFile *os.File, retries int) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fileStat, err := tmpFile.Stat()
	if err != nil {
		return err
	}

	currentSize := fileStat.Size()

	if currentSize > 0 {
		downloadBegin := strconv.FormatInt(currentSize, 10)
		req.Header.Set("Range", "bytes="+downloadBegin+"-")
	}

	customClient := &http.Client{Timeout: 4 * time.Minute}
	resp, err := customClient.Do(req)

	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return downloadFile(url, tmpFile, retries-1)
		}
		return err
	}else if err != nil {
		return err
	}

	if resp.StatusCode < 200 && resp.StatusCode > 300 {
		return fmt.Errorf("Unsuccessful error code - %v", resp.StatusCode)
	}

	if resp.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	_, err = io.Copy(tmpFile, resp.Body)

	if err != nil && hasTimedOut(err) {
		if retries > 0 {
			return downloadFile(url,  tmpFile, retries-1)
		}
		return err
	}else if err != nil{
		return err
	}
	return nil

}

func hasTimedOut(err error) bool {
	switch err := err.(type){
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout(){
			return true
		}
	case *net.OpError:
		if err.Timeout(){
			return true
		}
	}
	errContent := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errContent) {
		return true
	}
return false
}

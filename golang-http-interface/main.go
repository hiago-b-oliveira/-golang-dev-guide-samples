package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// bytesRead, _ := resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// fmt.Println(bytesRead)

	//io.Copy(os.Stdout, resp.Body)
	io.Copy(logWriter{}, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	myBs := make([]byte, len(bs))
	copy(myBs, bs)

	return len(bs), nil
}

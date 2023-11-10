package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	// bytes := make([]byte, 99999) // 99999 is the max size of the byte slice
	// resp.Body.Read(bytes)
	// fmt.Println(string(bytes))

	// Another way to do the above
	// io.Copy(os.Stdout, resp.Body)

	// Implementing a custom Writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

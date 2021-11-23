package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetHttpMethod(url string) io.ReadCloser {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error in creating new request : ", err)
	}

	req.SetBasicAuth("yxp200011@utdallas.edu", "Sonata@678")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Body
}

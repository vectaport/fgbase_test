package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var i int
	for {
		resp, err := http.Get("http://localhost:8080/count/a")
		if err != nil {
			fmt.Printf("Err from server: %v,%v\n", resp,err)
		}
		defer resp.Body.Close()
		
		body,err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Err from server: %v,%v\n", body,err)
		}
                i++
                if i%10000==0 {
                        fmt.Printf("request %d\n", i)
                }
	}
}

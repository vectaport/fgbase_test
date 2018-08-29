package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	portp := flag.Int("port", 8080, "port to use")
	testp := flag.Bool("test", false, "test mode")
	flag.Parse()
	port := *portp
	test := *testp

	if test {
		time.Sleep(4 * time.Second)
	}

	var i int
	for {
		req := fmt.Sprintf("http://localhost:%d/count/a", port)
		resp, err := http.Get(req)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Err from server: %v\n", err)
				break
			}
		}
		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Err from server: %v\n", err)
		}
		i++
	}
}

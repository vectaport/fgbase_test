package main

import (
	"flag"
	"fmt"
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
		time.Sleep(4*time.Second) 
	}

	var i int
	for {
		req := fmt.Sprintf("http://localhost:%d/count/a", port)
		resp, err := http.Get(req)
		if err != nil {
			if err.Error() == "Get "+req+": EOF" {
				break
			} else {
				fmt.Printf("Err from server: %v,%v\n", resp,err)
			}
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

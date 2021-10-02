package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main(){
	for {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			panic(err)
		}
		s,err:=ioutil.ReadAll(resp.Body)
		if s !=nil{}
		time.Sleep(time.Nanosecond)
		resp.Body.Close()
	}
}
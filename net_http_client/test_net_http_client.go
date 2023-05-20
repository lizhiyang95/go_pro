package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

var url = "http://httpbin.org"

func get() {
	resp, _ := http.Get(url + "/get")
	//fmt.Printf("r: %v\n", r)

	defer func() {
		resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)

}

func post() {
	resp, _ := http.Post(url+"/post", "", nil)
	//fmt.Printf("r: %v\n", r)

	defer func() {
		resp.Body.Close()
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)

}

func put() {
	req, err := http.NewRequest(http.MethodPut, url+"/put", nil)

	resp, err2 := http.DefaultClient.Do(req)

	defer func() {
		resp.Body.Close()
	}()
	if err2 != nil {
		panic(err2)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)
}

func getWithQuery() {
	req, err := http.NewRequest(http.MethodGet, url+"/get", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//query := make(url2.Values)
	var query = url2.Values{}
	query.Add("name", "tom")

	req.URL.RawQuery = query.Encode()
	fmt.Printf("req.URL.RawQuery: %v\n", req.URL.RawQuery)

	//增加header
	req.Header.Add("token", "dksajdkjsalkda")

	resp, err2 := http.DefaultClient.Do(req)

	defer func() {
		resp.Body.Close()
	}()
	if err2 != nil {
		panic(err2)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)
}

func getResponse() {
	req, err := http.NewRequest(http.MethodGet, url+"/get", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	resp, err2 := http.DefaultClient.Do(req)

	defer func() {
		resp.Body.Close()
	}()
	if err2 != nil {
		panic(err2)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b: %s\n", b)
	fmt.Printf("resp.Status: %v\n", resp.Status)
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	fmt.Printf("resp.Header.Get(\"content-type\"): %v\n", resp.Header.Get("content-type"))

}

func main() {
	//get()
	//post()
	//put()
	//getWithQuery()
	getResponse()
}

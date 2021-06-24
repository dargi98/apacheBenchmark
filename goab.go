package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var n, c, totalErrors int = 1000, 36, 0
var totalTime float64 = 0
var k bool = false
var address string = "https://www.docker.com/"
var client *http.Client

func concurrentRequests(errorsCH chan int, it int) {
	var resp *http.Response
	var err error

	for i := 0; i < it; i++ {
		if k {
			resp, err = client.Get(address)
		} else {
			resp, err = http.Get(address)
		}

		if err != nil {
			errorsCH <- 1
		} else {
			errorsCH <- 0
			resp.Body.Close()
		}
	}
}

func goab() {
	errorsCH := make(chan int, n)
	for i := 0; i < c; i++ {
		if n%c != 0 && n%c > i {
			go concurrentRequests(errorsCH, n/c+1)
		} else {
			go concurrentRequests(errorsCH, n/c)
		}
	}
	// Retrieve error values, and make sure that goroutines finish
	for i := 0; i < n; i++ {
		totalErrors += <-errorsCH
	}
}

func initClient() {
	tr := &http.Transport{
		MaxIdleConnsPerHost: n,
		TLSHandshakeTimeout: 0 * time.Second,
	}
	client = &http.Client{Transport: tr}
}

func setParameters() {
	var cs, ns, ks, ad string

	fmt.Println("Input c parameter (default is 1)")
	fmt.Scanln(&cs)
	c, _ = strconv.Atoi(cs)

	fmt.Println("Input n parameter (default is 1)")
	fmt.Scanln(&ns)
	n, _ = strconv.Atoi(ns)

	fmt.Println("Input 1 for k = true (default is false)")
	fmt.Scanln(&ks)
	if ks == "1" {
		k = true
		initClient()
	}

	fmt.Println("If the go server is running, input 1 for http://localhost:8080/ (default is https://www.docker.com/)")
	fmt.Scanln(&ad)
	if ad == "1" {
		address = "http://localhost:8080/"
	}
}

func printResults() {
	totalTime = totalTime / 1000000.0 // convert ns to ms
	// fmt.Println("Time of all transactions:", totalTime, "ms")
	fmt.Printf("Number of transactions per second, TPS: %.3f #/s \n", float64(n)/(totalTime/1000.0))
	fmt.Printf("Average latency: %.3f ms \n", totalTime/float64(n))
	fmt.Println("Amount of errors:", totalErrors)
	fmt.Printf("Percentatge of errors: %.2f", float64(totalErrors)/float64(n)*100.0)
	fmt.Print("% \n")
}

func main() {
	setParameters()
	fmt.Println("Benchmarking...")
	start := time.Now()
	goab()
	totalTime = float64(time.Since(start))
	printResults()
}

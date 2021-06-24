package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var n, c, totalErrors int = 1000, 300, 0
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
	// Retrieve error values, and make sure that goroutines end
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

	fmt.Println("Input c parameter")
	fmt.Scanln(&cs)
	c, _ = strconv.Atoi(cs)

	fmt.Println("Input n parameter")
	fmt.Scanln(&ns)
	n, _ = strconv.Atoi(ns)

	fmt.Println("Input 'y' for k = true, 'n' otherwise")
	fmt.Scanln(&ks)
	if ks == "y" {
		k = true
		initClient()
	}

	fmt.Println("Input server address or press 'enter' for default address: https://www.docker.com")
	fmt.Scanln(&ad)
	if ad != "" {
		address = ad
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
	// setParameters()

	start := time.Now()
	goab()
	totalTime = float64(time.Since(start))

	printResults()
}

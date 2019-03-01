package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func printPage(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", b)
	ch <- string(b)
}

func main() {

	ch := make(chan string)
	for i := 1; i <= 10; i++ {
		go printPage("https://google.com", ch)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}

}

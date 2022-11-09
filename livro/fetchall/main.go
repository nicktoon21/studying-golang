package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprint(fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url))
}

func archive(name string, text string){
	file, err := os.OpenFile(name,os.O_APPEND|os.O_CREATE| os.O_WRONLY, 0644)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		log.Println(err)
	}
}

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:]{
		archive("out.txt",fmt.Sprintln(<-ch)) 
	}
	fmt.Fprintf(os.Stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
	
}



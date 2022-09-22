// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chen-huicheng/GSWGo/crawl/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	filename := "./dst/" + url[len(*prefix)+1:]
	SaveFile(url, filename)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func SaveFile(url, filename string) {
	// fmt.Println(filename)
	// return
	if !strings.HasSuffix(filename, ".html") {
		filename += "-.html"
	}
	folderPath := filename[:strings.LastIndex(filename, "/")]
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, os.ModePerm)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("url:%s error:%s", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("url:%s StatusCode:%d", url, resp.StatusCode)
		return
	}
	f, err1 := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		log.Printf("file:%s error:%d", filename, err1)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		f.WriteString(string(buf[:n]))

	}
}

var url = flag.String("url", "", "web url")
var prefix = flag.String("prefix", "", "web url prefix")

//!+
func main() {
	flag.Parse()
	if *url == "" {
		fmt.Println("please use -url [web url]")
		return
	}
	if *prefix == "" {
		if strings.HasPrefix(*url, "http") {
			idx := strings.LastIndex(*url, "/")
			*prefix = (*url)[:idx]
		}
	}
	// https: //books.studygolang.com/gopl-zh/ch8/ch8-08.html
	var goNum int = 20
	worklist := make(chan []string, goNum)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string, goNum) // de-duplicated URLs

	// Add init url
	go func() { worklist <- []string{*url} }()

	var runGoNum sync.WaitGroup
	// Create goNum crawler goroutines to fetch each unseen link.
	for i := 0; i < goNum; i++ {
		go func() {
			for link := range unseenLinks {
				runGoNum.Add(1)
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
				runGoNum.Done()
			}
		}()
	}
	go func() {
		time.Sleep(time.Second)
		runGoNum.Wait()
		fmt.Println("wait run")
		close(worklist)

	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if ok := strings.HasPrefix(link, *prefix); !ok {
				continue
			}
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
	close(unseenLinks)
}

//!-

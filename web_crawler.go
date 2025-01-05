// Solution to the web crawler exercise from the Go Tour
// Used:
// -- wait groups to wait for all parallel calls to finish
// -- mutex locks to prevent concurrent access to the URLs cache
// Source: https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
)

// cache for dedup urls
type fetchState struct {
	crawled map[string]bool
	mux     sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// GOAL: Don't fetch the same URL twice.
// create a URL cache that checks if a url as been crawled
// needs to have locks for each url to prevent concurrency issues

func (fs *fetchState) isCrawled(url string) bool {
	fs.mux.Lock()
	// usage of defer here to run after func completion
	defer fs.mux.Unlock()

	// if url not crawled
	if _, ok := fs.crawled[url]; ok == false {
		return false
	}
	return true
}

func (fs *fetchState) setCrawled(url string) {
	fs.mux.Lock()
	defer fs.mux.Unlock()
	fs.crawled[url] = true
}

func (fs *fetchState) testAndSet(url string) bool {
	fs.mux.Lock()
	defer fs.mux.Unlock()
	r := fs.crawled[url]
	fs.crawled[url] = true
	return r
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
func Crawl(url string, depth int, fetcher Fetcher, fs *fetchState) {
	// defer wg.Done()
	// Check if already crawled
	if fs.isCrawled(url) == true {
		return
	}
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	// Marked as crawled, doesn't retry on error
	fs.setCrawled(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var done sync.WaitGroup
	for _, u := range urls {
		// wg.Add(1)
		done.Add(1)
		// go Crawl(u, depth-1, fetcher)
		go func(u string) {
			Crawl(u, depth-1, fetcher, fs)
			done.Done()
		}(u)
	}
	done.Wait()
	return
}

func makeState() *fetchState {
	return &fetchState{crawled: make(map[string]bool)}
}

func main() {
	// GOAL: Fetch URLs in parallel.
	// use wait group to store all crawl calls, and wait for completion
	fmt.Println(("Starting web crawler"))
	Crawl("https://golang.org/", 4, fetcher, makeState())
	fmt.Println(("Finished web crawler"))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		fmt.Printf("found:   %s\n", url)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

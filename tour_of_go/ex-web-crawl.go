package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns body of URL and
	// a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(
	url string,
	depth int,
	fetcher Fetcher,
	cache *SafeMap,
	ch chan string,
	wg *sync.WaitGroup) {

	defer wg.Done()

	// base case when we run out of depth
	if depth <= 0 {
		return
	}

	// avoiding repetitive urls
	if flag := cache.GetValue(url); flag {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	cache.Cache(url) // Caching visited url

	ch <- fmt.Sprintf("Found %s %q\n", url, body)
	for _, u := range urls {
		// avoiding repetitive urls
		if flag := cache.GetValue(u); flag {
			continue
		}
		cache.Cache(u) // Caching visited url

		go Crawl(u, depth-1, fetcher, cache, ch, wg)

	}

	return
}

func main() {
	d := 4
	// Creating a waitgroup to get data from all goros
	var wg sync.WaitGroup
	wg.Add(d + 1)
	cache := SafeMap{val: make(map[string]bool)}
	ch := make(chan string)
	go func() {
		defer close(ch)
		Crawl("https://golang.org/", d, fetcher, &cache, ch, &wg)
	}()

	go func() {
		wg.Wait()
	}()

	// Get our results from the queue
	for elem := range ch {
		fmt.Println(elem)
	}
}

// SafeMap is safe to use concurrently
// Based on https://go.dev/tour/concurrency/9
type SafeMap struct {
	mu  sync.Mutex
	val map[string]bool
}

func (sm *SafeMap) Cache(key string) {
	sm.mu.Lock()
	sm.val[key] = true
	sm.mu.Unlock()
}

func (sm *SafeMap) GetValue(key string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.val[key]
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
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

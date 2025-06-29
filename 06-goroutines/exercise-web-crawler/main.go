package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// 訪問済みURLを記録するマップ
var visited = make(map[string]bool)
var mu sync.Mutex // マップを安全に使うためのmutex

// Crawl uses fetcher to recursively crawl
// Pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done() // この関数が終わったらwgに完了を通知

	if depth <= 0 {
		return
	}

	// 既に訪問済みかチェック
	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}
	visited[url] = true
	mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// 見つかったURLを並列で処理
	for _, u := range urls {
		wg.Add(1)                         // 新しいgoroutineを追加することをwgに伝える
		go Crawl(u, depth-1, fetcher, wg) // 並列実行
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait() // すべてのgoroutineが完了するまで待つ

	time.Sleep(time.Millisecond * 100) // 出力が完了するまで少し待つ
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

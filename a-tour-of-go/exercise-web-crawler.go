// この演習では、ウェブクローラ( web crawler )を並列化するため、Goの並行性の特徴を使います。

// 同じURLを2度取ってくることなく並列してURLを取ってくるように、 Crawl 関数を修正してみてください(注1)。

// 補足: 工夫すれば Crawl 関数のみの修正で実装できますが、無理に Crawl 関数内部に収める必要はありません。

// ひとこと: mapにフェッチしたURLのキャッシュを保持できますが、mapだけでは並行実行時の安全性はありません!

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	wg := new(sync.WaitGroup)
	defer wg.Wait()

	mux := new(sync.Mutex)
	history := map[string]int{}

	var checkOrUpdateFetchedURL func(url string) bool
	checkOrUpdateFetchedURL = func(url string) bool {
		mux.Lock()
		defer mux.Unlock()
		if _, ok := history[url]; ok {
			return true
		}
		history[url]++
		return false
	}

	var crawler func(url string, depth int)
	crawler = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}

		if checkOrUpdateFetchedURL(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			wg.Add(1)
			go crawler(u, depth-1)
		}

		return
	}

	wg.Add(1)
	go crawler(url, depth)
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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

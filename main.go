package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func main() {
	var limit int
	flag.IntVar(&limit, "parallel", 10, "set limit of parallel requests")
	flag.Parse()

	if limit < 1 {
		limit = 10
	}

	args := os.Args
	if args[1] == "-parallel" {
		args = args[3:]
	}

	urls := make(chan url.URL, 20)
	go parseArgs(args[1:], urls)

	var waiter sync.WaitGroup
	for i := 0; i < limit; i++ {
		waiter.Add(1)
		go func() {
			printHashes(urls)
			waiter.Done()
		}()
	}

	waiter.Wait()
}

func parseArgs(args []string, addrCh chan<- url.URL) {
	for _, arg := range args {
		res, err := url.Parse(arg)
		if err != nil {
			fmt.Printf("\nparse argument %v with error %v", arg, err)
			continue
		}
		if res.Scheme == "" {
			res.Scheme = "http"
		}
		addrCh <- *res
	}
	close(addrCh)
}

func printHashes(ch <-chan url.URL) {
	for u := range ch {
		path := u.String()
		hash := getHash(path)
		fmt.Printf("\n%v %v", path, hash)
	}
}

func getHash(path string) (hash string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("\nsend request %v with error %v", path, err)
		return
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nread result of %v with error %v", path, err)
		return
	}

	return fmt.Sprintf("%x", md5.Sum(rawBody))
}

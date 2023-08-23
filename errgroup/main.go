package main

import (
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	urls := []string{
		"http://www.golang.org/",
		"http://www.google.com/",
	}
	for _, url := range urls {
		u := url
		eg.Go(func() error {
			resp, err := http.Get(u)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

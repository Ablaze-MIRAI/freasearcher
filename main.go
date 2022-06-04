package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	word := os.Args[1]

	ctns := getResp(word, "&language=ja-JP&safesearch=0")

	idx, err := fuzzyfinder.FindMulti(
		ctns,
		func(i int) string {
			return ctns[i].title
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return insertIndent(ctns[i].content)
		}))
	if err != nil {
		log.Fatal(err)
	}
	url := ctns[idx[0]].url
	openBowser(url)
}

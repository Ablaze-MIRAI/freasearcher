package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/mattn/go-runewidth"
)

var (
	isURL bool
	word  string
)

func init() {
	flag.BoolVar(&isURL, "u", false, "ブラウザを開かずURLを表示するか指定します。")
	flag.StringVar(&word, "w", "", "検索するキーワードを指定します。")
}

func main() {
	flag.Parse()

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
			return runewidth.Wrap(ctns[i].content, setWidth())
		}))
	if err != nil {
		log.Fatal(err)
	}
	url := ctns[idx[0]].url

	if isURL == true {
		fmt.Println(url)
	} else {
		openBowser(url)
	}
}

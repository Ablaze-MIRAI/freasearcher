package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/tidwall/gjson"
)

type Result struct{
	title string
	content string
	url string
}

func getResp(word string, opt string) ([]Result) {
	url := "https://freasearch.org/search?q=" + word + "&format=json" + opt

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	results := gjson.Get(string(bArray), "results")

	ctns := []Result{}

	for _, result := range results.Array(){
		// fmt.Println(gjson.Get(result.String(), "title"))
		// fmt.Println(gjson.Get(result.String(), "content"))
		// fmt.Println(gjson.Get(result.String(), "url"))
		// fmt.Println("============================================")

		title := gjson.Get(result.String(), "title").String()
		content := gjson.Get(result.String(), "content").String()
		url := gjson.Get(result.String(), "url").String()

		tmp := Result{title: title, content: content, url: url}
		ctns = append(ctns, tmp)

	}


	return ctns
}

func main() {
	word := os.Args[1]

	ctns := getResp(word, "&language=ja-JP&safesearch=0",)

	// for _, ctn := range ctns {
	// 	fmt.Println(ctn.title,)
	// }
	// idx, err := fuzzyfinder.FindMulti(
	// 	ctns,
	// 	func(i int,) string  {
	// 		return ctns[i].title
	// 	},
	// 	fuzzyfinder.WithPreviewWindow(func(i, w, h, int,) string  {
	// 		if i == -1{
	// 			return ""
	// 		}
	// 		return fmt.Sprintf("Title: %s\nDescription: %s", ctns[i].title, ctns[i].content,)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Printf(ctns[idx].url)

	idx, err := fuzzyfinder.FindMulti(
		ctns,
		func(i int) string {
			return ctns[i].title
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf("%s",
			ctns[i].content,)
		}))
		if err != nil {
			log.Fatal(err)
		}
		url := ctns[idx[0]].url
		fmt.Println(url)
}

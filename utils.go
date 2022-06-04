package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"github.com/tidwall/gjson"
	// "github.com/yitsushi/go-misskey/services/app"
)

type Result struct {
	title   string
	content string
	url     string
}

func splitWords(word string) string {
	splited := strings.Split(word, " ")
	return strings.Join(splited, "&")
}

func getResp(word string, opt string) []Result {
	url := "https://freasearch.org/search?q=" + splitWords(word) + "&format=json" + opt
	fmt.Println(url)

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

	for _, result := range results.Array() {

		title := gjson.Get(result.String(), "title").String()
		content := gjson.Get(result.String(), "content").String()
		url := gjson.Get(result.String(), "url").String()

		tmp := Result{title: title, content: content, url: url}
		ctns = append(ctns, tmp)

	}

	return ctns
}

func insertIndent(str string) string {
	runes := []rune(str)
	size := len(runes)

	var result string

	thrice := [][]rune{}

	for i := 0; i < size; i += 41 {
		end := i + 41
		if size < end {
			end = size
		}
		thrice = append(thrice, runes[i:end])
	}

	for _, k := range thrice {
		result = result + string(k) + "\n"
	}
	return result
}

func openBowser(url string) {
	fmt.Printf("Open the %s in your browser...", url)
	open.Run(url)
}

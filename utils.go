package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"syscall"

	"github.com/skratchdot/open-golang/open"
	"github.com/tidwall/gjson"
	"golang.org/x/crypto/ssh/terminal"
)

type Result struct {
	title   string
	content string
	url     string
}

func splitWords(word string) string {
	splited := strings.Split(word, " ")
	return strings.Join(splited, "+")
}

func getResp(word string, opt string) ([]Result, error) {
	url := "https://freasearch.org/search?q=" + splitWords(word) + "&format=json" + opt

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエストの構築に失敗しました: %w", err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエストの取得に失敗しました: %w", err)
	}

	defer resp.Body.Close()

	bArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスボディの読み込みに失敗しました: %w", err)
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

	return ctns, nil
}

func setWidth() int {
	width, _, err := terminal.GetSize(syscall.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	return (width / 2) - 7
}

func openBowser(url string) {
	fmt.Printf("Open the %s in your browser...", url)
	open.Run(url)
}

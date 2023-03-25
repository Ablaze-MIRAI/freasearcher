// Files with general processing such as URL generation
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/tidwall/gjson"
	"golang.org/x/term"
)

type result struct {
	title   string
	content string
	url     string
}

type param struct {
	Query      string
	Language   string
	SafeSearch int
}

func newURL(param param) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.freasearch.org",
		Path:   "search",
	}
	q := u.Query()
	q.Set("q", param.Query)
	q.Set("language", "ja-JP")
	u.RawQuery = q.Encode()

	return u.String()
}

func getResp(param param) ([]result, error) {
	req, err := http.NewRequest("GET", newURL(param), nil)
	if err != nil {
		return nil, fmt.Errorf("リクエストの作成に失敗しました: %w", err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエストの取得に失敗しました: %w", err)
	}

	defer resp.Body.Close()

	bArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスボディの読み込みに失敗しました: %w", err)
	}

	results := gjson.Get(string(bArray), "results")

	ctns := []result{}

	for _, r := range results.Array() {

		title := gjson.Get(r.String(), "title").String()
		content := gjson.Get(r.String(), "content").String()
		url := gjson.Get(r.String(), "url").String()

		tmp := result{title: title, content: content, url: url}
		ctns = append(ctns, tmp)

	}

	return ctns, nil
}

func setWidth() int {
	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}

	return (width / 2) - 7
}

func openBrowser(url string) {
	fmt.Printf("Opening %s in your browser...", url)
	open.Run(url)
}

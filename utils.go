package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"

	"github.com/code-raisan/gouseragent"

	"github.com/skratchdot/open-golang/open"
	"github.com/tidwall/gjson"
	"golang.org/x/term"
)

type Result struct {
	title   string
	content string
	url     string
}

type Param struct {
	Query      string
	Language   string
	SafeSearch int
}

func newURL(param Param) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "freasearch.org",
		Path:   "search",
	}
	q := u.Query()
	q.Set("format", "json")
	q.Set("q", param.Query)
	q.Set("language", param.Language)
	q.Set("safesearch", strconv.Itoa(param.SafeSearch))
	u.RawQuery = q.Encode()

	return u.String()
}

func getUserAgent() string {
	switch runtime.GOOS {
	case "windows":
		return gouseragent.MozillaFirefoxWindows
	case "linux":
		return gouseragent.MozillaFirefoxLinux
	case "darwin":
		return gouseragent.MozillaFirefoxMac
	default:
		return gouseragent.MozillaFirefoxLinux
	}
}

func getResp(param Param) ([]Result, error) {
	fmt.Println(newURL(param))
	req, err := http.NewRequest("GET", newURL(param), nil)
	ua := getUserAgent()
	req.Header.Add("User-Agent", ua)

	if err != nil {
		return nil, fmt.Errorf("リクエストの作成に失敗しました: %w", err)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("リクエストの取得に失敗しました: %w", err)
	}

	defer resp.Body.Close()
	fmt.Println(resp.Body)

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
	fd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}

	return (width / 2) - 7
}

func openBrowser(url string) {
	fmt.Printf("Open the %s in your browser...", url)
	open.Run(url)
}

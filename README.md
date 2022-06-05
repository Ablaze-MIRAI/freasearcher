<div align="center">

# 🔍 freasearcher

</div>

<div align="center">

[Frea Search](https://freasearch.org/)のAPIを用いて素早く検索し、ターミナル上で閲覧できるコマンドラインツール

</div>

![実行中のgif画像](./img/t-rec.gif)

## 🔰 使い方

```
# キーワードで検索し、ブラウザで開く
freasearcher -w keyword

# 複数キーワードを指定する
freasearcher -w "search some keyword"

# URLを表示して終了する
freasearcher -u -w keyword
```

**注意⚠️ : デフォルトの検索言語は日本語に設定されています。また、現状変更出来ません。**

1. 引数に検索したいキーワードを指定してコマンドを実行します。
1. 検索結果をfuzzyfinderで絞り込みます。`Ctrl-N`, `Ctrl-P` または `Ctrl-J`, `Ctrl-K` でフォーカスを移動します。 `Tab`キーで選択し `Enter` キーで確定します。
1. 選択したサイトのURLが出力されます。

### 🌏 ブラウザで記事のページを開く

ブラウザで記事のページを開く機能が実装されました！🎉
また、`-u`オプションを追加するとURLを表示して終了することも可能です。

### ⚡ 高度な検索

Not yet... 💤

## ⬇️  インストール

リリースページから実行可能なバイナリをダウンロードしてください。

> [Latest Release]()

Not yet... 💤

ソースからビルドする場合は、このリポジトリをクローンして `go install` を実行してください。
`go1.18.2 linux/amd64`にて開発しています。

`go install`をする場合は以下のように実行してください。
`go install https://github.com/Ablaze-MIRAI/freasearcher@latest`

## ⛏️   開発

```sh
# install dependencies
go mod tidy

# Run freasearcher
go run main.go

# Build
fo build
```
## 📝 Todo

- [ ] 検索オプションの追加
- [ ] 英語のドキュメントの準備
- [ ] FreaSearch APIのドキュメント

## 📜 ライセンス

[MIT](LICENSE)

## 👏 影響を受けたプロジェクト 

- [sheepla/fzwiki](https://github.com/sheepla/fzwiki)
- [sheepla/fzenn](https://github.com/sheepla/fzenn)

## 💕  スペシャルサンクス
mattn
- [mattn/go-runewidth](https://github.com/mattn/go-runewidth)

ktr0731
- [ktr0731/go-fuzzyfinder](https://github.com/ktr0731/go-fuzzyfinder)

Markus Heiser & Alexandre Flament
- [searxng/searxng](https://github.com/searxng/searxng)

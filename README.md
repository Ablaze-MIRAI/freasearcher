<div align="center">

# 🔍 freasearcher

</div>

<div align="center">

[Frea Search](https://freasearch.org/)のAPIを用いて素早く検索し、ターミナル上で閲覧できるコマンドラインツール

</div>

## 🔰 使い方

```
freasearcher keyword
```
:::note info
デフォルトの検索言語は日本語に設定されています。
:::

1. 引数に検索したいキーワードを指定してコマンドを実行します。
1. 検索結果をfuzzyfinderで絞り込みます。`Ctrl-N`, `Ctrl-P` または `Ctrl-J`, `Ctrl-K` でフォーカスを移動します。 `Tab`キーで選択し `Enter` キーで確定します。
1. 選択したサイトのURLが出力されます。

### 🌏 ブラウザで記事のページを開く

freasearcher自体にブラウザを開く機能は搭載していません。お使いのブラウザで実行できるようなスクリプトをご自身でお書きください。
例として、Fishで指定したサイトを開く関数を例示します。

```fish
function __freasearch
	firefox (freasearcher $argv[1])
end

```

### ⚡ 高度な検索

Not yet... 💤

## ⬇️  インストール

リリースページから実行可能なバイナリをダウンロードしてください。

Not yet... 💤
<!-- > [Latest Release]() -->

ソースからビルドする場合は、このリポジトリをクローンして `go install` を実行してください。
`go1.18.2 linux/amd64`にて開発しています。

## ⛏️   開発

```sh
# install dependencies
go mod tidy

# Run freasearcher
go run main.go

# Build
fo build
```

## 📜 ライセンス

[MIT](LICENSE)

## 👏 影響を受けたプロジェクト 

- [sheepla/fzwiki](https://github.com/sheepla/fzwiki)
- [sheepla/fzenn](https://github.com/sheepla/fzenn)

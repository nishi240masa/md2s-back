# MD2S-back
[MD2S](https://github.com/watagassa/md2s) のバックエンド部分です。

## utils/slide の使い方
```go
type RequestBody struct {
	Md    string `json:"md"`    // 変換したいマークダウン
	Title string `json:"title"` // スライドのタイトル
	Style int    `json:"style"` // スライドのテーマ（0~5）
}
```
変換に必要な情報を上記のように定義しているので、合致する形式で json データを送ると
```ts
{
  "style": 変換したスライド（string型）
}
```
という形式の json が送信されます。
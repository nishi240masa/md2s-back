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
上記の形式でjsonデータを送ると、**slide** というキーでスライド変換後のデータが文字列型で送信されます。
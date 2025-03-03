package slide

import (
	"context"
	"fmt"
	"log"
	"math"
	"md2s/dto"
	"md2s/utils/slide/slidethemes"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"google.golang.org/api/option"
)

// スライド1ページの型指定
type Slide struct {
	Title   string
	Content string
}

// ノード内のテキストを再帰的に抽出する関数
func extractText(n ast.Node, content []byte) string {
	var result string
	ast.Walk(n, func(child ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if child.Kind() == ast.KindText || child.Kind() == ast.KindString {
				result += string(child.Text(content))
			}
		}
		return ast.WalkContinue, nil
	})
	return result
}

// Qiita独自マークダウンを判定する関数
func isQiitaBlock(content string) bool {
	// Qiita独自のマークダウン構文をチェック
	return strings.Contains(strings.TrimSpace(content), ":::")
}

// Qiita独自ブロックからテキストを抽出する関数
func extractTextFromQiitaBlock(blockText string) string {
	var result strings.Builder
	lines := strings.Split(blockText, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Qiita独自マークダウンのタグ行（:::で始まる行）は無視
		if strings.HasPrefix(trimmed, ":::") {
			continue
		}
		result.WriteString(trimmed + "\n")
	}
	return result.String()
}

// マークダウンをページ（ヘッダー基準）ごとに分ける
var images []string    // 画像のURL分離用
var images_index []int // 分離した画像があった配列番号
func parseMarkdown(content []byte) ([]*Slide, error) {

	// Goldmarkの初期化
	mdParser := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // GitHub Flavored Markdown
		),
	)
	reader := text.NewReader([]byte(content))
	doc := mdParser.Parser().Parse(reader)

	var slides []*Slide
	var currentSlide *Slide

	// ASTを歩いてスライドを構築
	var count = 0
	var afterOption = false
	err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			fmt.Println(n.Kind())
			switch n.Kind() {
			case ast.KindHeading:
				heading := n.(*ast.Heading)
				headingText := extractText(heading, content)
				if heading.Level <= 4 { // h1,h2,h3,h4 to title
					if currentSlide != nil {
						slides = append(slides, currentSlide)
					}
					currentSlide = &Slide{
						Title:   headingText,
						Content: "",
					}
					count++
				}
				afterOption = true
			case ast.KindTextBlock, ast.KindText:
				// すべてのテキストベースのノードを検査
				var textContent string
				if afterOption {
					afterOption = false
				} else {
					textContent = extractText(n, content)
					if isQiitaBlock(textContent) {
						// Qiita独自のマークダウンブロックからテキストを抽出
						text := extractTextFromQiitaBlock(textContent)
						if currentSlide != nil {
							currentSlide.Content += text + "\n"
						}
						return ast.WalkSkipChildren, nil
					} else if currentSlide != nil {
						currentSlide.Content += textContent + "\n"
					}
				}
			case ast.KindRawHTML:
				if currentSlide != nil {
					rawHtml := n.(*ast.RawHTML)
					currentSlide.Content += "\n" + string(rawHtml.Text(content))
				}
			case ast.KindHTMLBlock:
				if currentSlide != nil {
					html := n.(*ast.HTMLBlock)
					currentSlide.Content += "\n" + string(html.Text(content)) + "\n"
				}
			case ast.KindListItem:
				if currentSlide != nil {
					// listItem := n.(*ast.ListItem)
					// currentSlide.Content += "- " + extractText(listItem, content) + "\n"
					afterOption = true
				}
			case ast.KindCodeBlock:
				if currentSlide != nil {
					codeBlock := n.(*ast.CodeBlock)
					currentSlide.Content += "\n```\n" + string(codeBlock.Text(content)) + "\n```\n"
				}
			case ast.KindCodeSpan:
				if currentSlide != nil {
					codeBlock := n.(*ast.CodeSpan)
					currentSlide.Content += "`" + string(codeBlock.Text(content)) + "`\n"
				}
			case ast.KindFencedCodeBlock:
				if currentSlide != nil {
					codeBlock := n.(*ast.FencedCodeBlock)
					currentSlide.Content += "\n```\n" + string(codeBlock.Text(content)) + "\n```\n"
				}
			case ast.KindImage:
				if currentSlide != nil {
					image := n.(*ast.Image)
					imageSrc := string(image.Destination) // 画像のURL
					images = append(images, fmt.Sprintf("\n---\n![bg fit](%s)\n", imageSrc))
					images_index = append(images_index, count)
					afterOption = true
				}
			case ast.KindLink:
				if currentSlide != nil {
					link := n.(*ast.Link)
					linkDest := string(link.Destination) // リンク先
					linkText := extractText(n, content)  // リンクテキスト
					currentSlide.Content += fmt.Sprintf("\n[%s](%s)\n", linkText, linkDest)
					afterOption = true
				}
			case ast.KindAutoLink:
				if currentSlide != nil {
					link := n.(*ast.AutoLink)
					linkDest := string(link.URL(content)) // リンク先
					currentSlide.Content += fmt.Sprintf("\n[リンク](%s)\n", linkDest)
					afterOption = true
				}
			}
		}
		return ast.WalkContinue, nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to walk AST: %w", err)
	}

	// 最後のスライドを追加
	if currentSlide != nil {
		slides = append(slides, currentSlide)
	}
	return slides, nil
}

// Gemini でページごとの内容をスライドっぽくする
func AnalyzeContentWithGemini(slides []*Slide) ([]*Slide, error) {
	ctx := context.Background()

	// Gemini APIクライアントを作成する
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Gemini のモデル指定
	model := client.GenerativeModel("gemini-1.5-flash")

	// スライドを15個ずつに分割する
	fmt.Println("[slide length]:", len(slides))
	var s_size = 13
	var slide_parts [][]*Slide // 分割したスライドの二次元配列 1要素あたり15個のページ
	if len(slides) > s_size {
		block := math.Ceil(float64(len(slides)) / float64(s_size))
		for i := 0; i < int(block); i++ {
			start := i * s_size
			end := start + s_size
			if end > len(slides) {
				end = len(slides)
			}
			slide_parts = append(slide_parts, slides[start:end])
		}
	} else {
		slide_parts = append(slide_parts, slides[0:])
	}

	for j, slide_part := range slide_parts {
		var wg sync.WaitGroup

		fmt.Println() // ログを見やすくするために改行
		for i, slide := range slide_part {

			wg.Add(1)
			go func() {
				defer wg.Done()
				// プロンプト設定するとこ
				prompt := fmt.Sprintf("コンテンツを箇条書きプレゼン調に要約。コンテンツがない場合は空白を2個出力。それ以外は要約のみ出力 \n\n以下コンテンツ\n\n%s", slide.Content)
				// Gemini API を使用してコンテンツを最適化
				fmt.Println("[send] index:", i)
				resp, err := model.GenerateContent(ctx, genai.Text(prompt))
				if err != nil {
					fmt.Println("[ERROR] at index:", i, "\n", err)
					return
				}
				// レスポンスをスライドに代入
				for _, part := range resp.Candidates[0].Content.Parts {
					slide.Content = fmt.Sprintln(part)
				}
			}()
		}
		wg.Wait()
		if len(slides) > s_size && j != len(slide_parts)-1 {
			time.Sleep(62 * time.Second) // 送信時に若干時間がズレるため少し余裕を持たせる
		}
		// 分離しておいた画像を代入
		var image_counter = 0
		for i, slide := range slides {
			if slices.Contains(images_index, i+1) {
				slide.Content += fmt.Sprintln(images[image_counter])
				image_counter++
			}
		}
	}

	return slides, nil
}

// marpタグを冒頭に追加、ページの分かれたスライドを連結
func ConvertToMarp(slides []*Slide, title []byte, style int) string {
	var marpBuilder strings.Builder
	marpBuilder.WriteString("---\nmarp: true\n") // Marpタグ
	marpBuilder.WriteString(slidethemes.ThemeList[style])
	marpBuilder.WriteString("---\n# ")
	marpBuilder.WriteString(string(title))
	marpBuilder.WriteString("\n<style scoped>section{font-size:50px;text-align:center}</style>")

	for _, slide := range slides {
		marpBuilder.WriteString("\n---\n")
		marpBuilder.WriteString(fmt.Sprintf("# %s\n\n", slide.Title))
		marpBuilder.WriteString(fmt.Sprintf("%s\n", slide.Content))
	}

	return marpBuilder.String()
}

func DeleteEscape(content []byte) (result []byte) {
	// strc := string(content)

	// unescaped, err := strconv.Unquote(strc)
	// if err != nil {
	// 	fmt.Println("[ERROR] parse failed")
	// }
	result = []byte(content)
	return result
}

func generateTitle(content []byte) (title []byte) {
	ctx := context.Background()
	// Gemini APIクライアントを作成する
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Gemini のモデル指定
	model := client.GenerativeModel("gemini-1.5-flash")

	// プロンプト設定するとこ
	prompt := fmt.Sprintf("コンテンツをもとに短いタイトルを1つ作ってください。作ったタイトルだけ出力してください。コンテンツがない場合は何も出力しないでください。\n\n以下コンテンツ\n\n%s", string(content))
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		fmt.Println("[ERROR] ", err)
		return
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		title = []byte(fmt.Sprintln(part))
	}

	return title
}

func MD2S(content []byte, title []byte, style int) (marpContent string, err error) {
	// マークダウンをページごとに変換
	slides, err := parseMarkdown(content)
	if err != nil {
		return "convert failed.", err
	}

	// Gemini で内容をスライドっぽくする
	analyzedSlides, err := AnalyzeContentWithGemini(slides)
	if err != nil {
		return "convert failed", err
	}

	// 連結＆marpタグ追加
	marpContent = ConvertToMarp(analyzedSlides, title, style)

	return marpContent, nil
}

func SlideConverter(input dto.RequestBody) (marp string, err error) {

	content := []byte(input.Md)
	title := []byte(input.Title)
	style := input.Style

	// タイトルが空の時AIで生成
	if string(title) == "" {
		fmt.Println("Title is empty. Generating title...")
		title = generateTitle(content)
	}

	// インデックス外参照対策
	if style > 5 {
		style = 0
	}

	decoded := DeleteEscape(content)

	// 文字列変換の例（全て大文字に変換）
	marp, err = MD2S(decoded, title, style)

	return marp, err
}

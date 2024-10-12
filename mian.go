package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"os"
	"path"
	"strings"
)

var (
	markdownFile string
	themes       string
	newHtmlFile  string
)

func init() {
	flag.StringVar(&markdownFile, "f", "", "markdown文件")
	flag.StringVar(&newHtmlFile, "o", "", "指定写入的文件夹")
	flag.StringVar(&themes, "t", "typorawiki", fmt.Sprintf("主题文件,支持：%s", strings.Join(getThemesName(), ",")))
	flag.Parse()
}

func main() {
	if markdownFile == "" {
		fmt.Println("请指定markdown文件")
		os.Exit(1)
	}
	htmlFile := strings.TrimSuffix(path.Base(markdownFile), path.Ext(markdownFile)) + ".html"
	if newHtmlFile != "" {
		htmlFile = path.Join(newHtmlFile, htmlFile)
	}
	dir := path.Dir(htmlFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	source, err := os.ReadFile(markdownFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile(htmlFile, append(buf.Bytes(), getThemes(themes)[:]...), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

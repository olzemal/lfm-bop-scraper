// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package scrape

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

func ScrapeBopTable(url string, wait time.Duration) ([]byte, error) {
	dom, err := discoverDom(url, wait)
	if err != nil {
		return nil, err
	}
	info, err := bopInfo(dom)
	if err != nil {
		return nil, err
	}

	w := new(bytes.Buffer)
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "h3" {
			fmt.Fprintf(w, "\n%s", node.FirstChild.Data)
		}
		if node.Type == html.ElementNode && node.Data == "tr" {
			fmt.Fprint(w, "\n")
		}
		if node.Type == html.ElementNode && node.Data == "td" {
			fmt.Fprintf(w, "%s\t", node.FirstChild.Data)
		}
		if node.Type == html.ElementNode && node.Data == "th" {
			fmt.Fprintf(w, "%s\t", node.FirstChild.Data)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(info)
	return io.ReadAll(w)
}

func bopInfo(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if isMatchingRoot(node) {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <mat-tab-body id=\"mat-tab-content-1-0\"> in the node tree")
}

func isMatchingRoot(node *html.Node) bool {
	match := node.Type == html.ElementNode && node.Data == "mat-tab-body"
	if !match {
		return false
	}
	for _, a := range node.Attr {
		if a.Key == "id" && a.Val == "mat-tab-content-1-0" {
			return true
		}
	}
	return false
}

func discoverDom(url string, wait time.Duration) (*html.Node, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		return nil, err
	}

	time.Sleep(wait)

	var str string
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			str, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	})
	if err != nil {
		return nil, err
	}

	return html.Parse(strings.NewReader(str))
}

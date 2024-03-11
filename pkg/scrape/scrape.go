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

const (
	loadInterval = time.Second
	initialWait  = time.Second * 5
	url          = "https://lowfuelmotorsport.com/seasonsv2/bop"
)

func BopTable(ctx context.Context) ([]byte, error) {
	cc, cancel := chromedp.NewContext(ctx)
	defer cancel()
	err := chromedp.Run(cc, chromedp.Navigate(url))
	if err != nil {
		return nil, err
	}

	errorChan := make(chan error)
	nodeChan := make(chan *html.Node)
	go func() {
		time.Sleep(initialWait)
		for {
			node, err := loadContent(cc)
			if err != nil {
				errorChan <- err
			}
			nodeChan <- node
			time.Sleep(loadInterval)
		}
	}()

	for {
		select {
		case err := <-errorChan:
			return nil, err
		case node := <-nodeChan:
			table, err := findBopTable(node)
			if err != nil {
				continue
			}
			buf := new(bytes.Buffer)
			crawlTable(table, buf)
			if buf.Len() == 0 {
				continue
			}
			return io.ReadAll(buf)
		case <-ctx.Done():
			return nil, fmt.Errorf("context deadline exceeded")
		}
	}
}

func loadContent(ctx context.Context) (*html.Node, error) {
	var str string
	err := chromedp.Run(ctx, chromedp.Tasks{
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

func crawlTable(node *html.Node, w *bytes.Buffer) {
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
		crawlTable(child, w)
	}
}

func findBopTable(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if isBopTableParent(node) {
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

func isBopTableParent(node *html.Node) bool {
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

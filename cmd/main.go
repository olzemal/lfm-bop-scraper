// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/olzemal/lfmbopscraper/pkg/parse"
	"github.com/olzemal/lfmbopscraper/pkg/scrape"
)

func lfmBopScraper() error {
	outfile := ""
	flag.StringVar(&outfile, "o", "", "output file")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	table, err := scrape.BopTable(ctx)
	if err != nil {
		return err
	}
	cfg, err := parse.TableToConfig(table)
	if err != nil {
		return err
	}
	j, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	var f *os.File
	if outfile == "" {
		f = os.Stdout
	} else {
		f, err = os.Create(outfile)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := lfmBopScraper()
	if err != nil {
		log.Fatal(err)
	}
}

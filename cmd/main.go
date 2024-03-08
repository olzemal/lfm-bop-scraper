// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package main

//go:generate go run -tags generate github.com/google/addlicense -c "Alexander Olzem" -l apache -y 2024 -s=only ..

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

func main() {
	outfile := ""
	flag.StringVar(&outfile, "o", "", "output file")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	table, err := scrape.BopTable(ctx)
	cfg := parse.TableToConfig(table)
	j, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var f *os.File
	if outfile == "" {
		f = os.Stdout
	} else {
		f, err = os.Create(outfile)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()

	_, err = f.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

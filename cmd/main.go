package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/olzemal/lfmbopscraper/pkg/parse"
	"github.com/olzemal/lfmbopscraper/pkg/scrape"
)

const (
	url = "https://lowfuelmotorsport.com/seasonsv2/bop"
)

var (
	wait    = time.Second * 10
	outfile = ""
)

func main() {
	var w string
	flag.StringVar(&w, "w", "", "wait time in go time format (default 10s)")
	flag.StringVar(&outfile, "o", "", "output file")
	flag.Parse()

	if w != "" {
		d, err := time.ParseDuration(w)
		if err != nil {
			log.Fatal(err)
		}
		wait = d
	}

	table, err := scrape.ScrapeBopTable(url, wait)
	if err != nil {
		log.Fatal(err)
	}

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

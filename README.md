# LFM BoP scraper

Program to scrape the current Balance of Performance (BoP) for GT3 cars on
[lowfuelmotorsport.com](https://lowfuelmotorsport.com) and write it to a JSON file.

To generate a fresh `bop.json` simply run:

```bash
$ make
```

or alternatively run the command with go.

```bash
$ go run cmd/main.go -o bop.json
```

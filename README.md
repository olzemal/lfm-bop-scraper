# LFM BoP scraper

Program to scrape the current LFM Bop and write it to a JSON file.

To generate a fresh `bop.json` simply run

```bash
$ make
```

or alternatively run the command with go.

```bash
$ go run cmd/main.go -o bop.json
```

You might run into issues loading the LFM Page. If that happens you can tweak
the wait time with the `-w` parameter. For example increase the wait time to 20
seconds (the default is 10s).

```bash
$ go run cmd/main.go -o bop.json -w 20s
```

Finally please consider supporting LFM by joining their
[Patreon](https://lowfuelmotorsport.com/patreon). This tool is probably not as
accurate as a `bop.json` directly from them.

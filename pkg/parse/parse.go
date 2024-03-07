package parse

import (
	"strconv"
	"strings"

	"github.com/olzemal/lfmbopscraper/pkg/bop"
)

const (
	carNameColumn = 1
	ballastColumn = 3
)

var (
	ignoredPhrases = []string{"Track BoP Version", "Active since", "Class\tCar"}
)

func TableToConfig(b []byte) bop.ServerCfg {
	var bopMap = make(bop.TrackMap)
	var curMap = make(bop.CarMap)

	cur := ""
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			continue
		}
		if containsIgnoredPhrase(line) {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) == 1 {
			cur = bop.TrackNameMap[fields[0]]
			curMap = make(bop.CarMap)
		} else if len(fields) > 4 {
			car := bop.CarNameMap[fields[carNameColumn]]
			ballast := ballastToInt(fields[ballastColumn])
			curMap[car] = ballast
			bopMap[cur] = curMap
		}
	}

	output := bop.ServerCfg{}
	for t := range bopMap {
		for c := range bopMap[t] {
			output.Entries = append(output.Entries, bop.Entry{
				Track:    t,
				CarModel: c,
				Ballast:  bopMap[t][c],
			})
		}
	}

	return output
}

func containsIgnoredPhrase(s string) bool {
	for _, phrase := range ignoredPhrases {
		if strings.Contains(s, phrase) {
			return true
		}
	}
	return false
}

func ballastToInt(s string) int {
	i, err := strconv.Atoi(strings.Fields(s)[0])
	if err != nil {
		panic(err)
	}
	return i
}

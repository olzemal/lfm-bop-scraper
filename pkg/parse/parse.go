// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package parse

import (
	"strconv"
	"strings"
)

const (
	carNameColumn = 1
	carYearColumn = 2
	ballastColumn = 3
)

var (
	ignoredPhrases = []string{"Track BoP Version", "Active since", "Class\tCar"}
)

func TableToConfig(b []byte) ServerCfg {
	var bopMap = make(TrackMap)
	var curMap = make(CarMap)

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
			cur = TrackNameMap[fields[0]]
			curMap = make(CarMap)
		} else if len(fields) > 4 {
			car := CarNameMap[fields[carNameColumn]+" "+fields[carYearColumn]]
			ballast := ballastToInt(fields[ballastColumn])
			curMap[car] = ballast
			bopMap[cur] = curMap
		}
	}

	output := ServerCfg{}
	for t := range bopMap {
		for c := range bopMap[t] {
			output.Entries = append(output.Entries, Entry{
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

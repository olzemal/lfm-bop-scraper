package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type carMap map[int]int
type trackMap map[string]carMap

type serverCfg struct {
	Entries []entry `json:"entries"`
}

type entry struct {
	Track      string `json:"track"`
	CarModel   int    `json:"carModel"`
	Ballast    int    `json:"ballastKg"`
	Restrictor int    `json:"restrictor"`
}

const (
	carNameColumn = 1
	ballastColumn = 3
)

var (
	ignoredPhrases = []string{"Track BoP Version", "Active since", "Class\tCar"}
	carNameMap     = map[string]int{
		"Porsche 991 GT3 R":             0,
		"Ferrari 488 GT3":               2,
		"Audi R8 LMS":                   3,
		"Lamborghini Huracán GT3":       4,
		"McLaren 650S GT3":              5,
		"Nissan GT-R Nismo GT3":         6,
		"BMW M6 GT3":                    7,
		"Bentley Continental GT3":       8,
		"AMR V12 Vantage GT3":           12,
		"Reiter Engineering R-EX GT3":   13,
		"Emil Frey Jaguar G3":           14,
		"Lexus RC F GT3":                15,
		"Lamborghini Huracan GT3 Evo":   16,
		"Honda NSX GT3":                 17,
		"Audi R8 LMS Evo":               19,
		"AMR V8 Vantage":                20,
		"Honda NSX GT3 Evo":             21,
		"McLaren 720S GT3":              22,
		"Porsche 991 II GT3 R":          23,
		"Ferrari 488 GT3 Evo":           24,
		"Mercedes-AMG GT3":              25,
		"BMW M4 GT3":                    30,
		"Audi R8 LMS GT3 Evo 2":         31,
		"Ferrari 296 GT3":               32,
		"Lamborghini Huracan GT3 Evo 2": 33,
		"Porsche 992 GT3 R":             34,
	}
	trackNameMap = map[string]string{
		"Circuit de Catalunya":          "barcelona",
		"BR":                            "brands_hatch",
		"Circuit Of The Americas":       "cota",
		"DO":                            "donington",
		"HU":                            "hungaroring",
		"Autodromo Enzo e Dino Ferrari": "imola",
		"IN":                            "indianapolis",
		"Kyalami":                       "kyalami",
		"LA":                            "laguna_seca",
		"MI":                            "misano",
		"MO":                            "monza",
		"Mount Panorama Circuit":        "mount_panorama",
		"Nürburgring":                   "nurburgring",
		"OU":                            "oulton_park",
		"Circuit de Paul Ricard":        "paul_ricard",
		"SI":                            "silverstone",
		"SN":                            "snetterton",
		"Circuit de Spa Francorchamps":  "spa",
		"Suzuka Circuit":                "suzuka",
		"Watkins Glen":                  "watkins_glen",
		"ZA":                            "zandvoort",
		"ZO":                            "zolder",
		"Circuit Ricardo Tormo":         "valencia",
		"Spielberg - Red Bull Ring":     "red_bull_ring",
	}
)

func main() {
	var bopMap = make(trackMap)
	var curMap = make(carMap)

	b, err := os.ReadFile("./accbop.txt")
	if err != nil {
		panic(err)
	}
	c := string(b)

	cur := ""
	for _, line := range strings.Split(c, "\n") {
		if containsIgnoredPhrase(line) {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) == 1 {
			cur = trackNameMap[fields[0]]
			curMap = make(carMap)
		} else if len(fields) > 4 {
			car := carNameMap[fields[carNameColumn]]
			ballast := ballastToInt(fields[ballastColumn])
			curMap[car] = ballast
			bopMap[cur] = curMap
		}
	}

	output := serverCfg{}

	for t := range bopMap {
		for c := range bopMap[t] {
			output.Entries = append(output.Entries, entry{
				Track:    t,
				CarModel: c,
				Ballast:  bopMap[t][c],
			})
		}
	}

	o, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(o))
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

func (b trackMap) String() string {
	out := ""
	for track := range b {
		out += "Track: " + track + "\n"
		for car := range b[track] {
			out += fmt.Sprintf("%35s\t %vkg\n", car, b[track][car])
		}
	}

	return out
}

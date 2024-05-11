// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package parse

var (
	// Map of LFM Car Names to ACC IDs
	// LFM Names: https://lowfuelmotorsport.com/statistics/carstats
	// ACC IDs: https://www.acc-wiki.info/wiki/Server_Configuration#ID_Lists
	CarNameMap = map[string]int{
		"Porsche 991 GT3 R 2018":       0,
		"Mercedes-AMG GT3 2015":        1,
		"Ferrari 488 GT3 2018":         2,
		"Audi R8 LMS 2015":             3,
		"Lamborghini Huracan GT3 2015": 4,
		"McLaren 650S GT3 2015":        5,
		"Nissan GT-R Nismo GT3 2018":   6,
		"BMW M6 GT3 2017":              7,
		"Bentley Continental 2018":     8,
		// 9 not GT3
		"Nissan GT-R Nismo GT3 2015":       10,
		"Bentley Continental 2015":         11,
		"AMR V12 Vantage GT3 2013":         12,
		"Reiter Engineering R-EX GT3 2017": 13,
		"Emil Frey Jaguar G3 2012":         14,
		"Lexus RC F GT3 2016":              15,
		"Lamborghini Huracan GT3 Evo 2019": 16,
		"Honda NSX GT3 2017":               17,
		// 18 not GT3
		"Audi R8 LMS Evo 2019":     19,
		"AMR V8 Vantage 2019":      20,
		"Honda NSX GT3 Evo 2019":   21,
		"McLaren 720S GT3 2019":    22,
		"Porsche 991II GT3 R 2019": 23,
		"Ferrari 488 GT3 Evo 2020": 24,
		"Mercedes-AMG GT3 2020":    25,
		"BMW M4 GT3 2021":          26,
		// 27 not GT3
		// 28 not GT3
		// 29 not GT3
		// 30 not GT3
		"Audi R8 LMS GT3 evo II 2022":        31,
		"Ferrari 296 GT3 2023":               32,
		"Lamborghini Huracan GT3 EVO 2 2023": 33,
		"Porsche 992 GT3 R 2023":             34,
		"McLaren 720S GT3 Evo 2023":          35,
		"Ford Mustang GT3":                   36,
		// 50+ are GT4
	}
	// Map of LFM to ACC server track names
	// LFM names: https://lowfuelmotorsport.com/tracks/records
	// ACC names: https://www.acc-wiki.info/wiki/Server_Configuration#ID_Lists
	TrackNameMap = map[string]string{
		"Circuit de Catalunya":          "barcelona",
		"Brands Hatch Circuit":          "brands_hatch",
		"Circuit Of The Americas":       "cota",
		"Donington Park":                "donington",
		"Hungaroring":                   "hungaroring",
		"Autodromo Enzo e Dino Ferrari": "imola",
		"Indianapolis":                  "indianapolis",
		"Kyalami":                       "kyalami",
		"Laguna Seca":                   "laguna_seca",
		"Misano":                        "misano",
		"Autodromo Nazionale di Monza":  "monza",
		"Mount Panorama Circuit":        "mount_panorama",
		"Nürburgring":                   "nurburgring",
		"Nürburgring Nordschleife 24h":  "nurburgring_24h",
		"Oulton Park":                   "oulton_park",
		"Circuit de Paul Ricard":        "paul_ricard",
		"Silverstone":                   "silverstone",
		"Snetterton":                    "snetterton",
		"Circuit de Spa Francorchamps":  "spa",
		"Suzuka Circuit":                "suzuka",
		"Watkins Glen":                  "watkins_glen",
		"Zandvoort":                     "zandvoort",
		"Zolder":                        "zolder",
		"Circuit Ricardo Tormo":         "valencia",
		"Spielberg - Red Bull Ring":     "red_bull_ring",
	}
)

type CarMap map[int]int

type TrackMap map[string]CarMap

type ServerCfg struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Track      string `json:"track"`
	CarModel   int    `json:"carModel"`
	Ballast    int    `json:"ballastKg"`
	Restrictor int    `json:"restrictor"`
}

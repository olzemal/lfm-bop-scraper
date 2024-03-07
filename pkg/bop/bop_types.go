// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package bop

var (
	CarNameMap = map[string]int{
		"Porsche 991 GT3 R":             0,
		"Ferrari 488 GT3":               2,
		"Audi R8 LMS":                   3,
		"Lamborghini Huracán GT3":       4,
		"McLaren 650S GT3":              5,
		"Nissan GT-R Nismo GT3":         6,
		"BMW M6 GT3":                    7,
		"Bentley Continental":           8,
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
		"Audi R8 LMS GT3 evo II":        31,
		"Ferrari 296 GT3":               32,
		"Lamborghini Huracan GT3 EVO 2": 33,
		"Porsche 992 GT3 R":             34,
		"McLaren 720S GT3 Evo":          35,
	}
	TrackNameMap = map[string]string{
		"Circuit de Catalunya":          "barcelona",
		"BR":                            "brands_hatch",
		"Circuit Of The Americas":       "cota",
		"DO":                            "donington",
		"Hungaroring":                   "hungaroring",
		"Autodromo Enzo e Dino Ferrari": "imola",
		"Indianapolis":                  "indianapolis",
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

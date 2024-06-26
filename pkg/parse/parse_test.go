// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package parse_test

import (
	"slices"
	"testing"

	"github.com/olzemal/lfmbopscraper/pkg/parse"
)

func TestTableToConfig(t *testing.T) {
	type test struct {
		in  []byte
		out parse.ServerCfg
	}

	tests := []test{
		{
			in:  []byte(""),
			out: parse.ServerCfg{},
		},
		{
			in: []byte("\n" +
				"Kyalami\n" +
				"GT3\tBMW M4 GT3\t2021\t9 kg\t-1 kg\n"),
			out: parse.ServerCfg{
				Entries: []parse.Entry{
					{
						Track:    "kyalami",
						CarModel: 26,
						Ballast:  9,
					},
				},
			},
		},
		{
			in: []byte("\n" +
				"Kyalami\n" +
				"GT3\tBMW M4 GT3\t2021\t9 kg\t-1 kg\n" +
				"Nürburgring\n" +
				"GT3\tBMW M4 GT3\t2021\t12 kg\t-1 kg\n" +
				"GT3\tFerrari 296 GT3\t2023\t8 kg\t-1 kg\n"),
			out: parse.ServerCfg{
				Entries: []parse.Entry{
					{
						Track:    "kyalami",
						CarModel: 26,
						Ballast:  9,
					},
					{
						Track:    "nurburgring",
						CarModel: 26,
						Ballast:  12,
					},
					{
						Track:    "nurburgring",
						CarModel: 32,
						Ballast:  8,
					},
				},
			},
		},
	}

	for _, c := range tests {
		got, err := parse.TableToConfig(c.in)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got == nil {
			t.Fatalf("Got unexpected <nil>")
		}
		if len(got.Entries) != len(c.out.Entries) {
			t.Fatalf("Got\n%+v\nbut want\n%+v", *got, c.out)
		}
		for _, g := range got.Entries {
			if !slices.Contains(c.out.Entries, g) {
				t.Fatalf("Got\n%+v\nbut want\n%+v", got, c.out)
			}
		}
	}
}

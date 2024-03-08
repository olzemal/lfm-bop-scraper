// Copyright 2024 Alexander Olzem
// SPDX-License-Identifier: Apache-2.0

package scrape_test

import (
	"context"
	"testing"
	"time"

	"github.com/olzemal/lfmbopscraper/pkg/scrape"
)

func TestLoadPage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	bop, err := scrape.BopTable(ctx)
	if err != nil {
		t.Fatalf("Unexpected error when loading page: %v", err)
	}
	t.Logf("Could Load BoP Table with %d bytes", len(bop))
}

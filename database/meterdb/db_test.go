// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package meterdb

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stretchr/testify/require"

	"github.com/luxdefi/node/database"
	"github.com/luxdefi/node/database/memdb"
)

func TestInterface(t *testing.T) {
	for _, test := range database.Tests {
		baseDB := memdb.New()
		db, err := New("", prometheus.NewRegistry(), baseDB)
		require.NoError(t, err)

		test(t, db)
	}
}

func FuzzInterface(f *testing.F) {
	for _, test := range database.FuzzTests {
		baseDB := memdb.New()
		db, err := New("", prometheus.NewRegistry(), baseDB)
		if err != nil {
			require.NoError(f, err)
		}
		test(f, db)
	}
}

func BenchmarkInterface(b *testing.B) {
	for _, size := range database.BenchmarkSizes {
		keys, values := database.SetupBenchmark(b, size[0], size[1], size[2])
		for _, bench := range database.Benchmarks {
			baseDB := memdb.New()
			db, err := New("", prometheus.NewRegistry(), baseDB)
			require.NoError(b, err)
			bench(b, db, "meterdb", keys, values)
		}
	}
}

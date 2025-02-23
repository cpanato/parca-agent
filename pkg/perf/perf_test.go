// Copyright 2022-2023 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package perf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerfMapParse(t *testing.T) {
	res, err := ReadPerfMap("testdata/nodejs-perf-map")
	require.NoError(t, err)
	require.Len(t, res.addrs, 28)
	// Check for 4edd3cca B0 LazyCompile:~Timeout internal/timers.js:55
	require.Equal(t, res.addrs[12], MapAddr{0x4edd4f12, 0x4edd4f47, "LazyCompile:~remove internal/linkedlist.js:15"})

	// Look-up a symbol.
	sym, err := res.Lookup(0x4edd4f12 + 4)
	require.NoError(t, err)
	require.Equal(t, sym, "LazyCompile:~remove internal/linkedlist.js:15")

	_, err = res.Lookup(0xFFFFFFFF)
	require.ErrorIs(t, err, ErrNoSymbolFound)
}

func TestPerfMapParseErlangPerfMap(t *testing.T) {
	_, err := ReadPerfMap("testdata/erlang-perf-map")
	require.NoError(t, err)
}

func BenchmarkPerfMapParse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ReadPerfMap("testdata/nodejs-perf-map")
		require.NoError(b, err)
	}
}

func BenchmarkPerfMapParseBig(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ReadPerfMap("testdata/erlang-perf-map")
		require.NoError(b, err)
	}
}

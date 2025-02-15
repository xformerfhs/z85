//
// SPDX-FileCopyrightText: Copyright 2025 Frank Schwab
//
// SPDX-License-Identifier: Apache-2.0
//
// SPDX-FileType: SOURCE
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: Frank Schwab
//
// Version: 1.0.0
//
// Change history:
//    2025-02-14: V1.0.0: Created.
//

package z85_test

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"runtime"
	"testing"
	"z85"
)

// ******** Private constants  ********

// dataSize is the number of bytes in the test data.
const dataSize = 1024

// testData contains the data the benchmarks run one.
var testData = makeData(dataSize)

// ******** Benchmark functions ********

// BenchmarkBase64 runs a benchmark of the Base64 encoding.
func BenchmarkBase64(b *testing.B) {
	enc := base64.StdEncoding.WithPadding(base64.NoPadding)

	runtime.GC()

	for b.Loop() {
		_ = enc.EncodeToString(testData)
	}
}

// BenchmarkBase32 runs a benchmark of the Base32 encoding.
func BenchmarkBase32(b *testing.B) {
	enc := base32.StdEncoding.WithPadding(base32.NoPadding)

	for b.Loop() {
		_ = enc.EncodeToString(testData)
	}
}

// BenchmarkZ85 runs a benchmark of the Z85 encoding.
func BenchmarkZ85(b *testing.B) {
	runtime.GC()

	for b.Loop() {
		_, _ = z85.Encode(testData)
	}
}

// ******** Private functions ********

// makeData builds the test data of a given size.
func makeData(size int) []byte {
	data := make([]byte, size)
	_, _ = rand.Read(data)
	return data
}

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
//    2025-02-15: V1.0.0: Created.
//

package z85_test

import (
	"bytes"
	crand "crypto/rand"
	"github.com/xformerfhs/z85"
	"math/rand"
	"strings"
	"testing"
)

// ******** Private constants ********

// clearTheOne contains the clear bytes of the one test case on the https://rfc.zeromq.org/spec/32 website.
var clearTheOne = []byte{0x86, 0x4f, 0xd2, 0x6f, 0xb5, 0x59, 0xf7, 0x5b}

// encodedTheOne contains the encoded string of the one test case on the https://rfc.zeromq.org/spec/32 website.
var encodedTheOne = `HelloWorld`

// iterationCount contains the iteration count for the general test.
const iterationCount = 100

// maxSliceSize is the maximum slice size for the general test.
const maxSliceSize = 128

// ******** Test functions ********

// TestGeneral is the general encode/decode test.
func TestGeneral(t *testing.T) {
	buffer := make([]byte, maxSliceSize)
	for i := 0; i < iterationCount; i++ {
		chunkLen := rand.Int31n(maxSliceSize>>2) + 1
		testSlice := buffer[:chunkLen<<2]
		_, _ = crand.Read(testSlice)

		encoded, err := z85.Encode(testSlice)
		if err != nil {
			t.Fatal(err)
		}

		var decoded []byte
		decoded, err = z85.Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(testSlice, decoded) {
			t.Fatalf(`decoded bytes don't match`)
		}
	}
}

// TestEncodeTheOne implements the one test case documented on the https://rfc.zeromq.org/spec/32 website.
func TestEncodeTheOne(t *testing.T) {
	encoded, err := z85.Encode(clearTheOne)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if encoded != encodedTheOne {
		t.Fatalf(`Encoding did not result in '%s', but '%s'`, encodedTheOne, encoded)
	}
}

// TestEncodeNil tests the encoding of a nil byte slice.
func TestEncodeNil(t *testing.T) {
	encoded, err := z85.Encode(nil)
	if err != nil {
		t.Fatalf(`Encoding failed: %v`, err)
	}

	if len(encoded) != 0 {
		t.Fatalf(`Encoding nil did not result in an empty string, but '%s'`, encoded)
	}
}

// TestEncodeWithInvalidLength tests if an error occurs encoding with an invalid length.
func TestEncodeWithInvalidLength(t *testing.T) {
	_, err := z85.Encode(clearTheOne[2:5])
	if err == nil {
		t.Fatal(`Invalid length did not result in an error`)
	} else {
		if !z85.IsErrInvalidLength(err) {
			t.Fatalf(`Wrong error when encoding invalid length string: '%v'`, err)
		}

		if !strings.HasSuffix(err.Error(), ` 4`) {
			t.Fatalf(`Invalid length did not result in an error message ending with %d: '%s'`, 4, err)
		}
	}
}

// TestDecodeTheOne implements the one test case documented on the https://rfc.zeromq.org/spec/32 website.
func TestDecodeTheOne(t *testing.T) {
	decoded, err := z85.Decode(encodedTheOne)
	if err != nil {
		t.Fatalf(`Decoding failed: %v`, err)
	}

	if !bytes.Equal(decoded, clearTheOne) {
		t.Fatalf(`Decoding did not result in expected bytes, but '% 02x'`, decoded)
	}
}

// TestDecodeEmpty tests if an empty string is decoded correctly.
func TestDecodeEmpty(t *testing.T) {
	decoded, err := z85.Decode(``)
	if err != nil {
		t.Fatalf(`Decode failed: %v`, err)
	}

	if len(decoded) != 0 {
		t.Fatalf(`Decoding an empty string created a non-empty slice: % 02x`, decoded)
	}
}

// TestDecodeInvalidLength tests if an error occurs with decoding an invalid length.
func TestDecodeInvalidLength(t *testing.T) {
	_, err := z85.Decode(`1234`)
	if err == nil {
		t.Fatal(`Invalid length did not result in an error`)
	} else {
		if !z85.IsErrInvalidLength(err) {
			t.Fatalf(`Wrong error when decoding invalid length string: '%v'`, err)
		}
		if !strings.HasSuffix(err.Error(), ` 5`) {
			t.Fatalf(`Invalid length did not result in an error message ending with %d: '%s'`, 5, err)
		}
	}
}

// TestDecodeInvalidCharTooLarge tests if an error occurs with decoding an invalid character outside the allowed range.
func TestDecodeInvalidCharTooLarge(t *testing.T) {
	_, err := z85.Decode(`123~5`)
	if err == nil {
		t.Fatal(`Invalid character did not result in an error`)
	} else {
		if !z85.IsErrInvalidByte(err) {
			t.Fatalf(`Wrong error when decoding invalid character: '%v'`, err)
		}
		if !strings.HasSuffix(err.Error(), ` 3: '~'`) {
			t.Fatalf(`Correct error with wrong text: '%v'`, err)
		}
	}
}

// TestDecodeInvalidChar tests if an error occurs with decoding an invalid character.
func TestDecodeInvalidChar(t *testing.T) {
	_, err := z85.Decode(`123455432112,45`)
	if err == nil {
		t.Fatal(`Invalid character did not result in an error`)
	} else {
		if !z85.IsErrInvalidByte(err) {
			t.Fatalf(`Wrong error when decoding invalid character: '%v'`, err)
		}
		if !strings.HasSuffix(err.Error(), ` 12: ','`) {
			t.Fatalf(`Correct error with wrong text: '%v'`, err)
		}
	}
}

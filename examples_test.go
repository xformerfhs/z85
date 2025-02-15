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
	"fmt"
	"z85"
)

// ExampleEncode shows how to use the Encode function.
func ExampleEncode() {
	sourceBytes := []byte{0xab, 0xcd, 0xef, 0xaa, 0x55, 0x33, 0x11, 0x00}
	encoded, err := z85.Encode(sourceBytes)
	if err != nil {
		fmt.Printf("Encoding error: %v\n", err)
	}

	fmt.Println(encoded)
	// Output: TiHakrwLi!
}

// ExampleDecode shows how to use the Decode function.
func ExampleDecode() {
	encoded := `W!a&T%e#R0`
	decoded, err := z85.Decode(encoded)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
	}

	fmt.Printf(`%02x`, decoded)
	// Output: b6f47967ffaf05f5
}

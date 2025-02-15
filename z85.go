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

// Package z85 implements Z85 encoding as specified in https://rfc.zeromq.org/spec/32.
package z85

import (
	"encoding/binary"
)

// ******** Private constants ********

// codeSize is the size of the encoding (i.e. the number of encoding characters).
const codeSize = 85

// byteChunkSize is the size of a byte chunk to be processed.
const byteChunkSize = 4

// byteChunkMask is the mask to check a byte chunk index.
const byteChunkMask = byteChunkSize - 1

// byteChunkShift is the shift value used for division by shifting.
const byteChunkShift = 2

// encodedChunkSize is the size of an encoded chunk.
const encodedChunkSize = 5

// decodeOffset is the offset of an encoded byte into the decode table.
// This is the ASCII value of the encoding character with the least value.
const decodeOffset = '!'

// ivEc is the encoding value for an invalid character.
// The name has to have a length of 4 in order to be exactly as long as a hex constant.
const ivEc = 0xff

// encodeTable is the table used for encoding.
var encodeTable = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.-:+=^!/*?&<>()[]{}@%$#`

// decodeTable is the decoding table with an offset of decodeOffset.
var decodeTable = []byte{
	0x44, ivEc, 0x54, 0x53, 0x52, 0x48, ivEc,
	0x4b, 0x4c, 0x46, 0x41, ivEc, 0x3f, 0x3e, 0x45,
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x40, ivEc, 0x49, 0x42, 0x4a, 0x47,
	0x51, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a,
	0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32,
	0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a,
	0x3b, 0x3c, 0x3d, 0x4d, ivEc, 0x4e, 0x43, ivEc,
	ivEc, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
	0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
	0x21, 0x22, 0x23, 0x4f, ivEc, 0x50}

// decodeMaxValue is the maximum acceptable byte value for decoding.
var decodeMaxValue = byte(len(decodeTable)) + decodeOffset - 1

// ******** Public functions ********

// Encode encodes a byte slice into a Z85 encoded string.
// The length of the slice must be a multiple of 4.
func Encode(source []byte) (string, error) {
	sourceLen := uint(len(source))

	if (sourceLen & byteChunkMask) != 0 {
		return ``, ErrInvalidLength(byteChunkSize)
	}

	chunkCount := sourceLen >> byteChunkShift
	result := make([]byte, sourceLen+chunkCount)
	destination := result
	for chunkIndex := uint(0); chunkIndex < chunkCount; chunkIndex++ {
		value := binary.BigEndian.Uint32(source[:byteChunkSize])

		// Generate 5 characters
		for i := byteChunkSize; i >= 0; i-- {
			valueDiv := value / codeSize
			destination[i] = encodeTable[value-(valueDiv*codeSize)]
			value = valueDiv
		}

		destination = destination[encodedChunkSize:]
		source = source[byteChunkSize:]
	}

	return string(result), nil
}

// Decode decodes a Z85 string into a byte slice.
// The length of the string must be a multiple of 5.
func Decode(source string) ([]byte, error) {
	sourceLen := uint(len(source))

	chunkCount := sourceLen / encodedChunkSize
	if sourceLen != chunkCount*encodedChunkSize {
		return nil, ErrInvalidLength(encodedChunkSize)
	}

	result := make([]byte, sourceLen-chunkCount)
	destination := result
	for chunkIndex := uint(0); chunkIndex < chunkCount; chunkIndex++ {
		value := uint32(0)
		for i := uint(0); i < encodedChunkSize; i++ {
			charByte := source[i]
			if charByte < decodeOffset || charByte > decodeMaxValue {
				return nil, &ErrInvalidByte{position: chunkIndex*encodedChunkSize + i, value: charByte}
			}

			encodedValue := decodeTable[charByte-decodeOffset]
			if encodedValue == ivEc {
				return nil, &ErrInvalidByte{position: chunkIndex*encodedChunkSize + i, value: charByte}
			}

			value = value*codeSize + uint32(encodedValue)
		}

		binary.BigEndian.PutUint32(destination, value)

		destination = destination[byteChunkSize:]
		source = source[encodedChunkSize:]
	}

	return result, nil
}

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

package z85

import (
	"errors"
	"fmt"
)

// ******** Private constants ********

// invalidLengthMessage contains the format for the error message when the input
// has a length that is not valid for the operation.
const invalidLengthMessage = `input length is not a multiple of %d`

// invalidByteMessage contains the format for the error message of an invalid byte.
const invalidByteMessage = `invalid byte at position %d: %q`

// ******** Public types and functions ********

// ErrInvalidLength is returned when the input has a length that is not valid for the operation.
type ErrInvalidLength byte

// Error returns the error message for an invalid length error.
func (e ErrInvalidLength) Error() string {
	return fmt.Sprintf(invalidLengthMessage, e)
}

// IsErrInvalidLength reports whether the supplied error is the ErrInvalidLength error.
func IsErrInvalidLength(err error) bool {
	var expectedErr ErrInvalidLength
	return errors.As(err, &expectedErr)
}

// ErrInvalidByte is returned when there is an invalid byte in the encoded string.
type ErrInvalidByte struct {
	position uint
	value    byte
}

// Error returns the error message for an invalid byte error.
func (e *ErrInvalidByte) Error() string {
	return fmt.Sprintf(invalidByteMessage, e.position, e.value)
}

// IsErrInvalidByte reports whether the supplied error is the ErrInvalidByte error.
func IsErrInvalidByte(err error) bool {
	var errInvalidByte *ErrInvalidByte
	return errors.As(err, &errInvalidByte)
}

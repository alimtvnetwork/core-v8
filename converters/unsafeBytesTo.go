// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package converters

import (
	"unsafe"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/defaulterr"
)

func UnsafeBytesToStringWithErr(unsafeBytes []byte) (string, error) {
	if unsafeBytes == nil {
		return constants.EmptyString, defaulterr.CannotProcessNilOrEmpty
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes)), nil
}

// UnsafeBytesToStrings
//
// # Returns string arrays from unsafe bytes
//
// []byte (1-byte elements) as []string (16-byte elements) via unsafe.Pointer,
// producing corrupted memory reads. Do not use.
// Use a proper loop conversion instead.
func UnsafeBytesToStrings(unsafeBytes []byte) []string {
	if unsafeBytes == nil {
		return nil
	}

	results := make([]string, len(unsafeBytes))
	for i, b := range unsafeBytes {
		results[i] = string(b)
	}

	return results
}

func UnsafeBytesToStringPtr(unsafeBytes []byte) *string {
	if unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(&unsafeBytes))
}

func UnsafeBytesToString(unsafeBytes []byte) string {
	if unsafeBytes == nil {
		return constants.EmptyString
	}

	return *(*string)(unsafe.Pointer(&unsafeBytes))
}

// UnsafeBytesPtrToStringPtr Returns string from unsafe bytes
//
// May panic on conversion if the bytes were not in unsafe pointer.
//
// Expressions:
// - return (*string)(unsafe.Pointer(allBytes))
func UnsafeBytesPtrToStringPtr(unsafeBytes []byte) *string {
	if unsafeBytes == nil {
		return nil
	}

	return (*string)(unsafe.Pointer(&unsafeBytes))
}

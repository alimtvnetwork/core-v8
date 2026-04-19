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

package coretests

import "encoding/json"

// AnyToBytes
//
// ## Steps:
//   - If already in  []byte then return as is.
//   - If already in  string then return as []byte(string).
//   - For rest of the cases, convert to json using Marshal and then returns the bytes
//
// Panic if json marshal has error.
func AnyToBytes(anyItem any) []byte {
	switch expectedAs := anyItem.(type) {
	case []byte:
		if expectedAs == nil {
			return nil
		}

		return expectedAs
	case string:
		return []byte(expectedAs)
	default:
		toBytes, err := json.Marshal(expectedAs)

		if err != nil {
			panic(err)
		}

		return toBytes
	}
}

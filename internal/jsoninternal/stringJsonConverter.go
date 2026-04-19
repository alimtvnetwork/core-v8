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

package jsoninternal

import (
	"encoding/json"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

type stringJsonConverter struct{}

func (it stringJsonConverter) SafeDefault(anyItem any) string {
	s, _ := it.Default(anyItem)

	return s
}

// Default
//
// It is not pretty JSON
func (it stringJsonConverter) Default(
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return string(b), err
}

// Pretty
//
// Default pretty json
func (it stringJsonConverter) Pretty(
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Prefix("", b)
}

func (it stringJsonConverter) StringValue(name string) []byte {
	doubleQuoted := fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		name,
	)

	return []byte(doubleQuoted)
}

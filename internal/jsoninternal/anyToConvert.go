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

import "encoding/json"

type anyToConvert struct{}

func (it anyToConvert) SafeString(anyItem any) string {
	s, _ := it.String(anyItem)

	return s
}

// String
//
// It is not pretty JSON
func (it anyToConvert) String(
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return string(b), err
}

func (it anyToConvert) PrettyString(
	prefix string,
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Prefix(prefix, b)
}

func (it anyToConvert) PrettyStringIndent(
	prefix,
	curIndent string,
	anyItem any,
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Indent(prefix, curIndent, b)
}

func (it anyToConvert) SafePrettyString(
	prefix string,
	anyItem any,
) string {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.Safe(prefix, b)
}

func (it anyToConvert) PrettyStringDefault(
	anyItem any,
) string {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.Safe("", b)
}

func (it anyToConvert) PrettyStringDefaultMust(
	anyItem any,
) string {
	b, err := json.Marshal(anyItem)

	if err != nil {
		panic(err)
	}

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.DefaultMust(b)
}

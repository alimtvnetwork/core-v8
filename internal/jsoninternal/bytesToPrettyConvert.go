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
	"bytes"
	"encoding/json"
)

type bytesToPrettyConvert struct{}

func (it bytesToPrettyConvert) Safe(
	prefix string, b []byte,
) string {
	output, _ := it.Prefix(prefix, b)

	return output
}

func (it bytesToPrettyConvert) SafeDefault(
	b []byte,
) string {
	output, _ := it.Prefix("", b)

	return output
}

func (it bytesToPrettyConvert) Prefix(
	prefix string,
	b []byte,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, prefix, Indent)

	return prettyJSON.String(), err
}

func (it bytesToPrettyConvert) Indent(
	prefix,
	curIndent string,
	b []byte,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, prefix, curIndent)

	return prettyJSON.String(), err
}

func (it bytesToPrettyConvert) PrefixMust(prefix string, b []byte) string {
	output, err := it.Prefix(prefix, b)

	if err != nil {
		panic(err)
	}

	return output
}

func (it bytesToPrettyConvert) DefaultMust(b []byte) string {
	output, err := it.Prefix("", b)

	if err != nil {
		panic(err)
	}

	return output
}

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

type stringToPrettyConvert struct{}

func (it stringToPrettyConvert) Safe(
	prefix,
	src string,
) string {
	output, _ := it.Prefix(prefix, src)

	return output
}

func (it stringToPrettyConvert) SafeDefault(
	src string,
) string {
	output, _ := it.Prefix("", src)

	return output
}

func (it stringToPrettyConvert) Prefix(
	prefix string,
	src string,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(src), prefix, Indent)

	return prettyJSON.String(), err
}

func (it stringToPrettyConvert) Indent(
	prefix string,
	curIndent,
	src string,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(
		&prettyJSON,
		[]byte(src),
		prefix,
		curIndent,
	)

	return prettyJSON.String(), err
}

func (it stringToPrettyConvert) PrefixMust(prefix string, src string) string {
	output, err := it.Prefix(prefix, src)

	if err != nil {
		panic(err)
	}

	return output
}

func (it stringToPrettyConvert) DefaultMust(src string) string {
	output, err := it.Prefix("", src)

	if err != nil {
		panic(err)
	}

	return output
}

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

package args

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/alimtvnetwork/core/coreutils/stringutil"
)

type String string

func (it String) Concat(s ...string) String {
	toStr := stringslice.Joins("", s...)

	return it + String(toStr)
}

func (it String) Join(sep string, s ...string) String {
	toStr := stringslice.Joins(sep, s...)

	return it + String(sep) + String(toStr)
}

func (it String) Split(sep string) corestr.SimpleSlice {
	return strings.Split(string(it), sep)
}

func (it String) DoubleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintValueDoubleQuotationFormat,
			it,
		),
	)
}

func (it String) DoubleQuoteQ() String {
	return String(
		fmt.Sprintf(
			constants.SprintDoubleQuoteFormat,
			it,
		),
	)
}

func (it String) SingleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintSingleQuoteFormat,
			it,
		),
	)
}

func (it String) ValueDoubleQuote() String {
	return String(
		fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			it,
		),
	)
}

func (it String) String() string {
	return string(it)
}

func (it String) Bytes() []byte {
	return []byte(it)
}

func (it String) Runes() []rune {
	return []rune(it)
}

func (it String) Length() int {
	return utf8.RuneCountInString(it.String())
}

func (it String) Count() int {
	return utf8.RuneCountInString(it.String())
}

func (it String) IsEmptyOrWhitespace() bool {
	return len(strings.TrimSpace(it.String())) == 0
}

func (it String) TrimSpace() String {
	return String(strings.TrimSpace(it.String()))
}

func (it String) ReplaceAll(old, new string) String {
	return String(strings.ReplaceAll(it.String(), old, new))
}

func (it String) TrimReplaceMap(replacerMap map[string]string) String {
	r := stringutil.ReplaceTemplate.DirectKeyUsingMapTrim(
		it.String(),
		replacerMap,
	)

	return String(r)
}

func (it String) Substring(start, length int) String {
	return it[start:length]
}

func (it String) IsEmpty() bool {
	return len(it) == 0
}

func (it String) HasCharacter() bool {
	return len(it) > 0
}

func (it String) IsDefined() bool {
	return len(it) > 0
}

func (it String) AscIILength() int {
	return len(it)
}

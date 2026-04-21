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

package coreinstruction

import (
	"regexp"
	"strings"

	"github.com/alimtvnetwork/core-v8/internal/strutilinternal"
)

type BaseIdentifier struct {
	Id string `json:"Id"`
}

func NewIdentifier(id string) *BaseIdentifier {
	return &BaseIdentifier{Id: id}
}

func (identifier *BaseIdentifier) IdString() string {
	return identifier.Id
}

func (identifier *BaseIdentifier) IsIdEmpty() bool {
	return identifier.Id == ""
}

func (identifier *BaseIdentifier) IsIdWhitespace() bool {
	return strutilinternal.IsNullOrEmptyOrWhitespace(&identifier.Id)
}

func (identifier *BaseIdentifier) IsId(id string) bool {
	return identifier.Id == id
}

func (identifier *BaseIdentifier) IsIdCaseInsensitive(idInsensitive string) bool {
	return strings.EqualFold(identifier.Id, idInsensitive)
}

func (identifier *BaseIdentifier) IsIdContains(idContains string) bool {
	return strings.Contains(identifier.Id, idContains)
}

func (identifier *BaseIdentifier) IsIdRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(identifier.Id)
}

func (identifier *BaseIdentifier) Clone() *BaseIdentifier {
	return &BaseIdentifier{
		Id: identifier.Id,
	}
}

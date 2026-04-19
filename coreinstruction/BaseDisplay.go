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
)

type BaseDisplay struct {
	Display string `json:"Display"`
}

func (receiver BaseDisplay) IsDisplay(display string) bool {
	return receiver.Display == display
}

func (receiver BaseDisplay) IsDisplayCaseInsensitive(display string) bool {
	return strings.EqualFold(receiver.Display, display)
}

func (receiver BaseDisplay) IsDisplayContains(displayContains string) bool {
	return strings.Contains(receiver.Display, displayContains)
}

func (receiver BaseDisplay) IsDisplayRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(receiver.Display)
}

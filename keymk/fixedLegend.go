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

package keymk

import "strings"

type fixedLegend struct{}

func (it fixedLegend) FormatKeyMap(
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (format string, replacerMap map[string]string) {
	return LegendChainSample, map[string]string{
		"{root}":    rootName,
		"{package}": packageName,
		"{group}":   groupName,
		"{state}":   stateName,
		"{user}":    userName,
		"{item}":    item,
	}
}

func (it fixedLegend) Compile(
	isKeepFormatOnEmpty bool,
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (compiled string) {
	format, compilingMap := it.FormatKeyMap(
		rootName,
		packageName,
		groupName,
		stateName,
		userName,
		item)

	for searcher, replacer := range compilingMap {
		if isKeepFormatOnEmpty && replacer == "" {
			continue
		}

		format = strings.ReplaceAll(
			format,
			searcher,
			replacer)
	}

	return format
}

func (it fixedLegend) CompileKeepFormatOnEmpty(
	rootName,
	packageName,
	groupName,
	stateName,
	userName,
	item string,
) (compiled string) {
	return it.Compile(
		true,
		rootName,
		packageName,
		groupName,
		stateName,
		userName,
		item)
}

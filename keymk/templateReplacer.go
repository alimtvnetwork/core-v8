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

import (
	"strconv"
	"strings"
)

type templateReplacer struct {
	key *Key
}

func (it *templateReplacer) RequestIntRange(
	isCurly bool,
	tempReplace TempReplace,
) []string {
	return it.IntRange(
		isCurly,
		tempReplace.KeyName,
		tempReplace.Range.Start,
		tempReplace.Range.End)
}

func (it *templateReplacer) IntRange(
	isCurly bool,
	keyName string,
	startIncluding, endIncluding int,
) []string {
	keyOuts := make(
		[]string,
		0,
		endIncluding-startIncluding+1)

	keyName = curlyWrapIf(isCurly, keyName)
	templateFormat := it.key.KeyCompiled() // format may hold {key-name}

	for i := startIncluding; i <= endIncluding; i++ {
		numString := strconv.Itoa(i)

		keyOuts = append(
			keyOuts,
			strings.ReplaceAll(
				templateFormat,
				keyName,
				numString))
	}

	return keyOuts
}

func (it *templateReplacer) CompileUsingReplacerMap(
	isCurly bool,
	replacerMap map[string]string, // key ==> find, value ==> replace
) string {
	templateFormat := it.key.KeyCompiled() // format may hold {key-name}

	if templateFormat == "" || len(replacerMap) == 0 {
		return templateFormat
	}

	for finder, replacer := range replacerMap {
		finderCurly := curlyWrapIf(isCurly, finder)

		templateFormat = strings.ReplaceAll(
			templateFormat,
			finderCurly,
			replacer)
	}

	return templateFormat
}

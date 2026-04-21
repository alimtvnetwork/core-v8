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

package strsort

import (
	"sort"

	"github.com/alimtvnetwork/core-v8/coredata"
)

// QuickPtr Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickPtr(pointerStringsIn *[]*string) *[]*string {
	pointerStrings := coredata.PointerStrings(*pointerStringsIn)
	sort.Sort(pointerStrings)

	return pointerStringsIn
}

// Quick Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func Quick(stringsPointerIn *[]string) *[]string {
	sort.Strings(*stringsPointerIn)

	return stringsPointerIn
}

// QuickDscPtr Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/_OkY82E2kO9
func QuickDscPtr(pointerStringsIn *[]*string) *[]*string {
	pointerStringsDsc := coredata.PointerStringsDsc(*pointerStringsIn)
	sort.Sort(pointerStringsDsc)

	return pointerStringsIn
}

// QuickDsc Warning: Data gets mutated inside.
//
// Reference : https://play.golang.org/p/sJ8a464USeV
func QuickDsc(stringsPointerIn *[]string) *[]string {
	pointerStringsDsc := coredata.StringsDsc(*stringsPointerIn)
	sort.Sort(pointerStringsDsc)

	return stringsPointerIn
}

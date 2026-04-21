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

package chmodhelper

import "github.com/alimtvnetwork/core-v8/internal/osconstsinternal"

type tempDirGetter struct{}

// TempDefault
//
// Example:
//   - unix    : /tmp
//   - windows : %temp%
func (it tempDirGetter) TempDefault() string {
	return TempDirDefault
}

// TempPermanent
//
// Windows:
//   - c:\\windows\\temp
//
// unix:
//   - /var/tmp/
//
// Reference:
//   - Why "/var/tmp/" : https://prnt.sc/gW0DA5d4jt6R
func (it tempDirGetter) TempPermanent() string {
	if osconstsinternal.IsWindows {
		return osconstsinternal.WindowsPermanentTemp
	}

	// unix
	return osconstsinternal.LinuxPermanentTemp
}

func (it tempDirGetter) TempOption(isPermanent bool) string {
	if isPermanent {
		return it.TempPermanent()
	}

	return it.TempDefault()
}

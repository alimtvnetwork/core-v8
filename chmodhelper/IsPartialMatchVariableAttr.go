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

import "github.com/alimtvnetwork/core-v8/constants"

// IsPartialMatchVariableAttr
//
//	givenVarAttr can have wildcards "*"
//	 On wildcard present comparison will ignore for that segment.
//
//	Example (will consider this a match):
//	 - givenVarAttr: (rwx : "r*x"),
//	 - rwx         : (rwx : "r-x")
func IsPartialMatchVariableAttr(
	givenVarAttr *VarAttribute,
	rwx string,
) bool {
	r, w, x := ExpandCharRwx(rwx)

	read := givenVarAttr.isRead.ToByteCondition(
		ReadChar,
		NopChar,
		constants.WildcardChar)
	write := givenVarAttr.isWrite.ToByteCondition(
		WriteChar,
		NopChar,
		constants.WildcardChar)
	execute := givenVarAttr.isExecute.ToByteCondition(
		ExecuteChar,
		NopChar,
		constants.WildcardChar,
	)

	isRead := givenVarAttr.isRead.IsWildcard() || (read == r)
	isWrite := givenVarAttr.isWrite.IsWildcard() || (write == w)
	isExecute := givenVarAttr.isExecute.IsWildcard() || (execute == x)

	return isRead &&
		isWrite &&
		isExecute
}

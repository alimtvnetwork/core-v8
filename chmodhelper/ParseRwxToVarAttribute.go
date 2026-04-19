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

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/issetter"
)

func ParseRwxToVarAttribute(rwx string) (varAttribute *VarAttribute, err error) {
	length := len(rwx)

	if length != SingleRwxLength {
		return nil, GetRwxLengthError(rwx)
	}

	r, w, x := ExpandCharRwx(rwx)

	// any is true
	isRead := issetter.GetBool(r == ReadChar)
	isWrite := issetter.GetBool(w == WriteChar)
	isExecute := issetter.GetBool(x == ExecuteChar)

	// is any has '*' wildcard
	isReadWildcard := r == constants.WildcardChar
	isWriteWildcard := w == constants.WildcardChar
	isExecuteWildcard := x == constants.WildcardChar

	isVarType := isReadWildcard ||
		isWriteWildcard ||
		isExecuteWildcard

	if isVarType {
		readVal := issetter.GetSet(
			isReadWildcard,
			issetter.Wildcard,
			isRead)

		writeVal := issetter.GetSet(
			isWriteWildcard,
			issetter.Wildcard,
			isWrite)

		execVal := issetter.GetSet(
			isExecuteWildcard,
			issetter.Wildcard,
			isExecute)

		return &VarAttribute{
			rawInput:    rwx,
			isFixedType: !isVarType,
			isRead:      readVal,
			isWrite:     writeVal,
			isExecute:   execVal,
		}, nil
	}

	return &VarAttribute{
		rawInput:    rwx,
		isFixedType: !isVarType,
		isRead:      isRead,
		isWrite:     isWrite,
		isExecute:   isExecute,
	}, nil
}

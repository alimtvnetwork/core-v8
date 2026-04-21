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

import "github.com/alimtvnetwork/core-v8/chmodhelper/chmodins"

func ParseRwxOwnerGroupOtherToRwxVariableWrapper(
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
) (
	*RwxVariableWrapper, error,
) {
	if rwxOwnerGroupOther == nil {
		return nil, rwxInstructionNilErr
	}

	ownerVarAttr, ownerErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Owner)

	if ownerErr != nil {
		return nil, ownerErr
	}

	groupVarAttr, groupErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Group)

	if groupErr != nil {
		return nil, groupErr
	}

	otherVarAttr, otherErr := ParseRwxToVarAttribute(rwxOwnerGroupOther.Other)

	if otherErr != nil {
		return nil, otherErr
	}

	rawInput := ParseRwxInstructionToStringRwx(
		rwxOwnerGroupOther,
		false)

	isFixedType := ownerVarAttr.IsFixedType() &&
		groupVarAttr.IsFixedType() &&
		otherVarAttr.IsFixedType()

	return &RwxVariableWrapper{
		rawInput:    rawInput,
		isFixedType: isFixedType,
		Owner:       *ownerVarAttr,
		Group:       *groupVarAttr,
		Other:       *otherVarAttr,
	}, nil
}

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
	"github.com/alimtvnetwork/core-v8/issetter"
)

type VarAttribute struct {
	rawInput    string
	isFixedType bool
	isRead      issetter.Value
	isWrite     issetter.Value
	isExecute   issetter.Value
}

func (it *VarAttribute) IsFixedType() bool {
	return it.isFixedType
}

func (it *VarAttribute) HasWildcard() bool {
	return !it.isFixedType
}

// ToCompileFixAttr
//
//	must check IsFixedType, before calling.
func (it *VarAttribute) ToCompileFixAttr() *Attribute {
	if it.isFixedType {
		return &Attribute{
			IsRead:    it.isRead.IsTrue(),
			IsWrite:   it.isWrite.IsTrue(),
			IsExecute: it.isExecute.IsTrue(),
		}
	}

	return nil
}

// ToCompileAttr
//
//	if fixed type then fixed param can be nil
func (it *VarAttribute) ToCompileAttr(fixed *Attribute) Attribute {
	if it.isFixedType {
		return Attribute{
			IsRead:    it.isRead.IsTrue(),
			IsWrite:   it.isWrite.IsTrue(),
			IsExecute: it.isExecute.IsTrue(),
		}
	}

	return Attribute{
		IsRead:    it.isRead.WildcardApply(fixed.IsRead),
		IsWrite:   it.isWrite.WildcardApply(fixed.IsWrite),
		IsExecute: it.isExecute.WildcardApply(fixed.IsExecute),
	}
}

func (it *VarAttribute) Clone() *VarAttribute {
	if it == nil {
		return nil
	}

	return &VarAttribute{
		rawInput:    it.rawInput,
		isFixedType: it.IsFixedType(),
		isRead:      it.isRead,
		isWrite:     it.isWrite,
		isExecute:   it.isExecute,
	}
}

func (it *VarAttribute) IsEqualPtr(next *VarAttribute) bool {
	if it == nil && next == nil {
		return true
	}

	if it == nil || next == nil {
		return false
	}

	isRead := next.isRead == it.isRead
	isWrite := next.isWrite == it.isWrite
	isExecute := next.isExecute == it.isExecute

	return isRead &&
		isWrite &&
		isExecute
}

func (it *VarAttribute) String() string {
	return it.rawInput
}

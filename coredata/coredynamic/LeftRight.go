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

package coredynamic

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

type LeftRight struct {
	Left, Right any
}

func (it *LeftRight) IsEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left) &&
			reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) HasAnyItem() bool {
	return !it.IsEmpty()
}

func (it *LeftRight) HasLeft() bool {
	return it != nil &&
		!reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) HasRight() bool {
	return it != nil &&
		!reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) IsLeftEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Left)
}

func (it *LeftRight) IsRightEmpty() bool {
	return it == nil ||
		reflectinternal.Is.Null(it.Right)
}

func (it *LeftRight) LeftReflectSet(
	toPointerOrBytesPointer any,
) error {
	if it == nil {
		return nil
	}

	return ReflectSetFromTo(it.Left, toPointerOrBytesPointer)
}

func (it *LeftRight) RightReflectSet(
	toPointerOrBytesPointer any,
) error {
	if it == nil {
		return nil
	}

	return ReflectSetFromTo(it.Right, toPointerOrBytesPointer)
}

func (it *LeftRight) DeserializeLeft() *corejson.Result {
	if it == nil {
		return nil
	}

	return corejson.NewPtr(it.Left)
}

func (it *LeftRight) DeserializeRight() *corejson.Result {
	if it == nil {
		return nil
	}

	return corejson.NewPtr(it.Right)
}

func (it *LeftRight) LeftToDynamic() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Left, true)
}

func (it *LeftRight) RightToDynamic() *Dynamic {
	if it == nil {
		return nil
	}

	return NewDynamicPtr(it.Right, true)
}

func (it *LeftRight) TypeStatus() TypeStatus {
	if it == nil {
		return TypeSameStatus(nil, nil)
	}

	return TypeSameStatus(it.Left, it.Right)
}

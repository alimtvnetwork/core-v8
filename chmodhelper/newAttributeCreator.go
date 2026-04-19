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

import "github.com/alimtvnetwork/core/errcore"

type newAttributeCreator struct{}

func (it newAttributeCreator) Create(
	isRead, isWrite, isExecute bool,
) Attribute {
	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

func (it newAttributeCreator) Default(
	isRead, isWrite, isExecute bool,
) Attribute {
	return Attribute{
		IsRead:    isRead,
		IsWrite:   isWrite,
		IsExecute: isExecute,
	}
}

// UsingRwxString
//
// Length must be 3
// "rwx" should be put for attributes.
//
// Examples:
//   - read enable all disable    : "r--"
//   - write enable all disable   : "-w-"
//   - execute enable all disable : "--x"
//   - all enabled                : "rwx"
func (it newAttributeCreator) UsingRwxString(
	rwx string,
) Attribute {
	length := len(rwx)

	if length != SingleRwxLength {
		panic(GetRwxLengthError(rwx))
	}

	r := rwx[0]
	w := rwx[1]
	e := rwx[2]

	return Attribute{
		IsRead:    r == ReadChar,
		IsWrite:   w == WriteChar,
		IsExecute: e == ExecuteChar,
	}
}

// UsingByteMust
//
//	Byte can be at most 0 to 7
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingByteMust(v7 byte) Attribute {
	attr, err := it.UsingByte(v7)

	if err != nil {
		panic(attr)
	}

	return attr
}

// UsingByte
//
//	Byte can be at most 0 to 7
//
// 1 - Execute true
// 2 - Write true
// 3 - Write + Execute true
// 4 - Read true
// 5 - Read + Execute true
// 6 - Read + Write true
// 7 - Read + Write + Execute all true
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingByte(v7 byte) (Attribute, error) {
	if ReadWriteExecute.IsGreaterThan(v7) {
		return Attribute{}, errcore.
			ShouldBeLessThanEqualType.
			Error(
				"v7 byte should not be more than "+ReadWriteExecute.String(),
				v7)
	}

	// Use standard Unix permission bit masks:
	// bit 2 = read (4), bit 1 = write (2), bit 0 = execute (1)
	return Attribute{
		IsRead:    v7&4 != 0,
		IsWrite:   v7&2 != 0,
		IsExecute: v7&1 != 0,
	}, nil
}

// UsingVariantMust
//
// safe because converting AttrVariant should never exceed 7
//
// Warning:
//
//	Panics if more than 7
func (it newAttributeCreator) UsingVariantMust(v AttrVariant) Attribute {
	return it.UsingByteMust(v.Value())
}

func (it newAttributeCreator) UsingVariant(v AttrVariant) (Attribute, error) {
	return it.UsingByte(v.Value())
}

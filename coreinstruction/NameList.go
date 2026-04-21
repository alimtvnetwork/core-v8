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

package coreinstruction

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
)

type NameList struct {
	Name string               `json:"Name,omitempty"`
	List *corestr.SimpleSlice `json:"List,omitempty"`
}

func (it *NameList) IsNull() bool {
	return it == nil
}

func (it *NameList) IsAnyNull() bool {
	return it == nil || it.List == nil
}

func (it *NameList) IsNameEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *NameList) HasName() bool {
	return it != nil && it.Name != ""
}

func (it NameList) String() string {
	return corejson.
		Serialize.
		ToString(it)
}

func (it *NameList) Clone(
	isDeepClone bool,
) *NameList {
	if it == nil {
		return nil
	}

	return &NameList{
		Name: it.Name,
		List: it.
			List.
			ClonePtr(isDeepClone),
	}
}

func (it *NameList) DeepClone() *NameList {
	return it.Clone(true)
}

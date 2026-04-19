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

import "reflect"

type CastedResult struct {
	Casted                         any
	SourceReflectType              reflect.Type
	SourceKind                     reflect.Kind
	Error                          error
	IsNull, IsMatchingAcceptedType bool
	IsValid                        bool
	IsPointer                      bool // refers to how returned, ptr or non ptr
	IsSourcePointer                bool
}

func (it *CastedResult) IsInvalid() bool {
	return it == nil || !it.IsValid
}

func (it *CastedResult) IsNotNull() bool {
	return it != nil && !it.IsNull
}

func (it *CastedResult) IsNotPointer() bool {
	return it != nil && !it.IsPointer
}

func (it *CastedResult) IsNotMatchingAcceptedType() bool {
	return it != nil && !it.IsMatchingAcceptedType
}

func (it *CastedResult) IsSourceKind(kind reflect.Kind) bool {
	return it != nil && it.SourceKind == kind
}

func (it *CastedResult) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *CastedResult) HasAnyIssues() bool {
	return it.IsInvalid() || it.IsNull || !it.IsMatchingAcceptedType
}

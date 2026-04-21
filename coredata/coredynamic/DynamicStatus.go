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

import "github.com/alimtvnetwork/core-v8/constants"

type DynamicStatus struct {
	Dynamic
	Index   int
	Message string
}

func InvalidDynamicStatusNoMessage() *DynamicStatus {
	return InvalidDynamicStatus(constants.EmptyString)
}

func InvalidDynamicStatus(message string) *DynamicStatus {
	return &DynamicStatus{
		Dynamic: NewDynamic(nil, false),
		Index:   constants.InvalidNotFoundCase,
		Message: message,
	}
}

// Clone Warning: Cannot clone dynamic data or interface properly but set it again
//
// If it is a pointer one needs to copy it manually.
func (it DynamicStatus) Clone() DynamicStatus {
	return DynamicStatus{
		Dynamic: it.Dynamic.Clone(),
		Index:   constants.InvalidNotFoundCase,
		Message: it.Message,
	}
}

func (it *DynamicStatus) ClonePtr() *DynamicStatus {
	if it == nil {
		return nil
	}

	return &DynamicStatus{
		Dynamic: it.Dynamic.Clone(),
		Index:   constants.InvalidNotFoundCase,
		Message: it.Message,
	}
}

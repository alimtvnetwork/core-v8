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

package corestr

type newSimpleStringOnceCreator struct{}

func (it *newSimpleStringOnceCreator) Any(
	isIncludeFieldNames bool,
	value any,
	isInitialize bool,
) SimpleStringOnce {
	toString := AnyToString(
		isIncludeFieldNames,
		value)

	return SimpleStringOnce{
		value:        toString,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) Uninitialized(
	value string,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: false,
	}
}

func (it *newSimpleStringOnceCreator) Init(
	value string,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func (it *newSimpleStringOnceCreator) InitPtr(
	value string,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: true,
	}
}

func (it *newSimpleStringOnceCreator) Create(
	value string,
	isInitialize bool,
) SimpleStringOnce {
	return SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) CreatePtr(
	value string,
	isInitialize bool,
) *SimpleStringOnce {
	return &SimpleStringOnce{
		value:        value,
		isInitialize: isInitialize,
	}
}

func (it *newSimpleStringOnceCreator) Empty() SimpleStringOnce {
	return SimpleStringOnce{}
}

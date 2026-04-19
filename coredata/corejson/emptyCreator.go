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

package corejson

type emptyCreator struct{}

func (it emptyCreator) Result() Result {
	return Result{}
}

func (it emptyCreator) ResultWithErr(
	typeName string,
	err error,
) Result {
	return Result{
		Error:    err,
		TypeName: typeName,
	}
}

func (it emptyCreator) ResultPtrWithErr(
	typeName string,
	err error,
) *Result {
	return &Result{
		Error:    err,
		TypeName: typeName,
	}
}

func (it emptyCreator) ResultPtr() *Result {
	return &Result{}
}

func (it emptyCreator) BytesCollection() BytesCollection {
	return BytesCollection{}
}

func (it emptyCreator) BytesCollectionPtr() *BytesCollection {
	return &BytesCollection{}
}

func (it emptyCreator) ResultsCollection() *ResultsCollection {
	return &ResultsCollection{
		Items: []Result{},
	}
}

func (it emptyCreator) ResultsPtrCollection() *ResultsPtrCollection {
	return &ResultsPtrCollection{
		Items: []*Result{},
	}
}

func (it emptyCreator) MapResults() *MapResults {
	return &MapResults{
		Items: map[string]Result{},
	}
}

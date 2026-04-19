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

type ReturningBool struct {
	IsBreak, IsKeep bool
}

type LinkedCollectionFilterResult struct {
	Value           *LinkedCollectionNode
	IsKeep, IsBreak bool
}

type LinkedListFilterResult struct {
	Value           *LinkedListNode
	IsKeep, IsBreak bool
}

type LinkedCollectionFilterParameter struct {
	Node  *LinkedCollectionNode
	Index int
}

type LinkedListFilterParameter struct {
	Node  *LinkedListNode
	Index int
}

type LinkedListProcessorParameter struct {
	Index                       int
	CurrentNode, PrevNode       *LinkedListNode
	IsFirstIndex, IsEndingIndex bool
}

type LinkedCollectionProcessorParameter struct {
	Index                       int
	CurrentNode, PrevNode       *LinkedCollectionNode
	IsFirstIndex, IsEndingIndex bool
}

type OnCompleteCharCollectionMap func(charCollection *CharCollectionMap)
type OnCompleteLinkedCollections func(linkedCollections *LinkedCollections)
type AnyToCollectionProcessor func(any any, index int) *Collection
type OnCompleteCharHashsetMap func(charHashset *CharHashsetMap)
type IsStringFilter func(str string, index int) (result string, isKeep bool, isBreak bool)
type IsKeyAnyValueFilter func(pair KeyAnyValuePair) (result string, isKeep bool, isBreak bool)
type IsKeyValueFilter func(pair KeyValuePair) (result string, isKeep bool, isBreak bool)
type IsStringPointerFilter func(stringPointer *string, index int) (result *string, isKeep bool, isBreak bool)
type LinkedListFilter func(arg *LinkedListFilterParameter) *LinkedListFilterResult
type LinkedListSimpleProcessor func(
	arg *LinkedListProcessorParameter,
) (isBreak bool)
type LinkedCollectionFilter func(
	arg *LinkedCollectionFilterParameter,
) *LinkedCollectionFilterResult
type LinkedCollectionSimpleProcessor func(
	arg *LinkedCollectionProcessorParameter,
) (isBreak bool)

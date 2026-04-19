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

type NonChainedLinkedListNodes struct {
	items             []*LinkedListNode
	isChainingApplied bool
}

func (it *NonChainedLinkedListNodes) IsChainingApplied() bool {
	return it.isChainingApplied
}

func (it *NonChainedLinkedListNodes) Items() []*LinkedListNode {
	return it.items
}

func NewNonChainedLinkedListNodes(
	capacity int,
) *NonChainedLinkedListNodes {
	items := make([]*LinkedListNode, 0, capacity)

	return &NonChainedLinkedListNodes{
		items: items,
	}
}

func (it *NonChainedLinkedListNodes) Length() int {
	if it.items == nil {
		return 0
	}

	return len(it.items)
}

func (it *NonChainedLinkedListNodes) IsEmpty() bool {
	return it.items == nil || len(it.items) == 0
}

func (it *NonChainedLinkedListNodes) Adds(
	nodes ...*LinkedListNode,
) *NonChainedLinkedListNodes {
	if nodes == nil {
		return it
	}

	for i := range nodes {
		it.items = append(
			it.items,
			nodes[i],
		)
	}

	return it
}

func (it *NonChainedLinkedListNodes) HasItems() bool {
	return !it.IsEmpty()
}

func (it *NonChainedLinkedListNodes) First() *LinkedListNode {
	return it.items[0]
}

func (it *NonChainedLinkedListNodes) FirstOrDefault() *LinkedListNode {
	if it.IsEmpty() {
		return nil
	}

	return it.items[0]
}

func (it *NonChainedLinkedListNodes) Last() *LinkedListNode {
	return it.items[it.Length()-1]
}

func (it *NonChainedLinkedListNodes) LastOrDefault() *LinkedListNode {
	if it.IsEmpty() {
		return nil
	}

	return it.items[it.Length()-1]
}

// ApplyChaining Warning Mutates data inside.
func (it *NonChainedLinkedListNodes) ApplyChaining() *NonChainedLinkedListNodes {
	length := it.Length()
	if length == 0 {
		return it
	}

	it.isChainingApplied = true
	for i, node := range it.items {
		if i+1 >= length {
			break
		}

		nextNode := it.items[i+1]
		node.next = nextNode
	}

	if it.HasItems() {
		it.Last().next = nil
	}

	return it
}

func (it *NonChainedLinkedListNodes) ToChainedNodes() []*LinkedListNode {
	length := it.Length()
	list := make([]*LinkedListNode, length)

	if length == 0 {
		return list
	}

	for i, node := range it.items {
		if i+1 >= length {
			break
		}

		curNode := node.Clone()
		list = append(list, curNode)
		nextNode := it.items[i+1]
		nextNodeClone := nextNode.Clone()
		curNode.next = nextNodeClone
		list = append(list, nextNodeClone)
	}

	return list
}

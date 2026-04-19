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

type NonChainedLinkedCollectionNodes struct {
	items             []*LinkedCollectionNode
	isChainingApplied bool
}

func NewNonChainedLinkedCollectionNodes(
	capacity int,
) *NonChainedLinkedCollectionNodes {
	items := make([]*LinkedCollectionNode, 0, capacity)

	return &NonChainedLinkedCollectionNodes{
		items: items,
	}
}

func (it *NonChainedLinkedCollectionNodes) IsChainingApplied() bool {
	return it.isChainingApplied
}

func (it *NonChainedLinkedCollectionNodes) Items() []*LinkedCollectionNode {
	return it.items
}

func (it *NonChainedLinkedCollectionNodes) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *NonChainedLinkedCollectionNodes) IsEmpty() bool {
	return it.items == nil || len(it.items) == 0
}

func (it *NonChainedLinkedCollectionNodes) Adds(
	nodes ...*LinkedCollectionNode,
) *NonChainedLinkedCollectionNodes {
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

func (it *NonChainedLinkedCollectionNodes) HasItems() bool {
	return !it.IsEmpty()
}

func (it *NonChainedLinkedCollectionNodes) First() *LinkedCollectionNode {
	return it.items[0]
}

func (it *NonChainedLinkedCollectionNodes) FirstOrDefault() *LinkedCollectionNode {
	if it.IsEmpty() {
		return nil
	}

	return it.items[0]
}

func (it *NonChainedLinkedCollectionNodes) Last() *LinkedCollectionNode {
	return it.items[it.Length()-1]
}

func (it *NonChainedLinkedCollectionNodes) LastOrDefault() *LinkedCollectionNode {
	if it.IsEmpty() {
		return nil
	}

	return it.items[it.Length()-1]
}

// ApplyChaining Warning Mutates data inside.
func (it *NonChainedLinkedCollectionNodes) ApplyChaining() *NonChainedLinkedCollectionNodes {
	length := it.Length()
	if length == 0 || it.isChainingApplied {
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

func (it *NonChainedLinkedCollectionNodes) ToChainedNodes() *[]*LinkedCollectionNode {
	length := it.Length()
	list := make([]*LinkedCollectionNode, length)

	if length == 0 {
		return &list
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

	return &list
}

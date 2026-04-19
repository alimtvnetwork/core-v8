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

type emptyCreator struct{}

func (it *emptyCreator) Collection() *Collection {
	return &Collection{
		items: []string{},
	}
}

func (it *emptyCreator) LinkedList() *LinkedList {
	return &LinkedList{}
}

func (it *emptyCreator) SimpleSlice() *SimpleSlice {
	return New.SimpleSlice.Empty()
}

func (it *emptyCreator) KeyAnyValuePair() *KeyAnyValuePair {
	return &KeyAnyValuePair{}
}

func (it *emptyCreator) KeyValuePair() *KeyValuePair {
	return &KeyValuePair{}
}

func (it *emptyCreator) KeyValueCollection() *KeyValueCollection {
	return &KeyValueCollection{
		KeyValuePairs: []KeyValuePair{},
	}
}

func (it *emptyCreator) LinkedCollections() *LinkedCollections {
	return &LinkedCollections{}
}

func (it *emptyCreator) LeftRight() *LeftRight {
	return &LeftRight{}
}

func (it *emptyCreator) SimpleStringOnce() SimpleStringOnce {
	return SimpleStringOnce{}
}

func (it *emptyCreator) SimpleStringOncePtr() *SimpleStringOnce {
	return &SimpleStringOnce{}
}

func (it *emptyCreator) Hashset() *Hashset {
	return &Hashset{
		hasMapUpdated: false,
		items:         map[string]bool{},
	}
}

func (it *emptyCreator) HashsetsCollection() *HashsetsCollection {
	return &HashsetsCollection{
		items: nil,
	}
}

func (it *emptyCreator) Hashmap() *Hashmap {
	return &Hashmap{
		items: map[string]string{},
	}
}

func (it *emptyCreator) CharCollectionMap() *CharCollectionMap {
	return &CharCollectionMap{
		items: nil,
	}
}

func (it *emptyCreator) KeyValuesCollection() *KeyValueCollection {
	return &KeyValueCollection{
		KeyValuePairs: nil,
	}
}

func (it *emptyCreator) CollectionsOfCollection() *CollectionsOfCollection {
	return &CollectionsOfCollection{
		items: nil,
	}
}

func (it *emptyCreator) CharHashsetMap() *CharHashsetMap {
	return &CharHashsetMap{
		items: nil,
	}
}

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

package corepayload

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/defaultcapacity"
)

// PayloadsCollectionFilter.go — Filter and search methods extracted from PayloadsCollection.go

func (it *PayloadsCollection) Filter(
	filterFunc FilterFunc,
) []*PayloadWrapper {
	list := make(
		[]*PayloadWrapper, 0, it.Length())

	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
		}

		if isBreak {
			return list
		}
	}

	return list
}

func (it *PayloadsCollection) FilterWithLimit(
	limit int,
	filterFunc FilterFunc,
) []*PayloadWrapper {
	length := defaultcapacity.MaxLimit(
		it.Length(),
		limit)
	list := make(
		[]*PayloadWrapper,
		0,
		length)

	collectedItems := 0
	for _, item := range it.Items {
		isTake, isBreak := filterFunc(item)

		if isTake {
			list = append(list, item)
			collectedItems++
		}

		if isBreak {
			return list
		}

		if collectedItems >= length {
			return list
		}
	}

	return list
}

func (it *PayloadsCollection) FirstByFilter(
	findByFunc func(payloadWrapper *PayloadWrapper) (isFound bool),
) *PayloadWrapper {
	items := it.Filter(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		isTake = findByFunc(payloadWrapper)

		return isTake, isTake
	})

	if len(items) > 0 {
		return items[0]
	}

	return nil
}

func (it *PayloadsCollection) FirstById(
	id string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsIdentifier(id)
	})
}

func (it *PayloadsCollection) FirstByCategory(
	category string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsCategory(category)
	})
}

func (it *PayloadsCollection) FirstByTaskType(
	taskType string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsTaskTypeName(taskType)
	})
}

func (it *PayloadsCollection) FirstByEntityType(
	entityType string,
) *PayloadWrapper {
	return it.FirstByFilter(func(payloadWrapper *PayloadWrapper) (isFound bool) {
		return payloadWrapper.IsEntityType(entityType)
	})
}

func (it *PayloadsCollection) FilterCollection(
	filterFunc FilterFunc,
) *PayloadsCollection {
	list := it.Filter(filterFunc)

	collection := New.PayloadsCollection.UsingWrappers(
		list...)

	return collection
}

func (it *PayloadsCollection) SkipFilterCollection(
	skipFilterFunc SkipFilterFunc,
) *PayloadsCollection {
	list := make(
		[]*PayloadWrapper,
		0,
		it.Length())

	for _, item := range it.Items {
		isSkip, isBreak := skipFilterFunc(item)
		isInclude := !isSkip

		if isInclude {
			list = append(list, item)
		}

		if isBreak {
			break
		}
	}

	return New.
		PayloadsCollection.
		UsingWrappers(list...)
}

func (it *PayloadsCollection) FilterCollectionByIds(
	ids ...string,
) *PayloadsCollection {
	idsHashmap := corestr.
		New.
		Hashset.
		Strings(ids)

	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return idsHashmap.Has(payloadWrapper.Identifier), false
	})
}

func (it *PayloadsCollection) FilterNameCollection(
	name string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.Name == name, false
	})
}

func (it *PayloadsCollection) FilterCategoryCollection(
	categoryName string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.CategoryName == categoryName, false
	})
}

func (it *PayloadsCollection) FilterEntityTypeCollection(
	entityTypeName string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.EntityType == entityTypeName, false
	})
}

func (it *PayloadsCollection) FilterTaskTypeCollection(
	taskType string,
) *PayloadsCollection {
	return it.FilterCollection(func(payloadWrapper *PayloadWrapper) (isTake, isBreak bool) {
		return payloadWrapper.TaskTypeName == taskType, false
	})
}

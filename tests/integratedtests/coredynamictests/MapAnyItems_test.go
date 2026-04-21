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

package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

func Test_MapAnyItems_AddAndKeys(t *testing.T) {
	tc := mapAnyItemsAddAndKeysTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	capacity := input.GetAsIntDefault("capacity", 10)
	keys := input["keys"].([]string)

	// Act
	mapItems := coredynamic.NewMapAnyItems(capacity)
	collection := corestr.New.Collection.Cap(10)
	collection.Adds("a", "b", "c")

	for _, key := range keys {
		mapItems.Add(key, collection)
	}

	allKeys := mapItems.AllKeys()
	hasAll := true
	for _, key := range keys {
		found := false
		for _, k := range allKeys {
			if k == key {
				found = true
				break
			}
		}
		if !found {
			hasAll = false
			break
		}
	}

	actual := args.Map{
		"keyCount": len(allKeys),
		"hasAll":   hasAll,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_Paged(t *testing.T) {
	tc := mapAnyItemsPagedTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	itemCount := input.GetAsIntDefault("itemCount", 9)
	pageSize := input.GetAsIntDefault("pageSize", 2)

	// Act
	mapItems := coredynamic.NewMapAnyItems(itemCount + 5)
	collection := corestr.New.Collection.Cap(5)
	collection.Adds("a", "b")

	for i := 0; i < itemCount; i++ {
		mapItems.Add(fmt.Sprintf("key-%d", i), collection)
	}

	pagedItems := mapItems.GetPagedCollection(pageSize)

	actual := args.Map{
		"pageCount": len(pagedItems),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_JsonRoundtrip(t *testing.T) {
	tc := mapAnyItemsJsonRoundtripTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	itemCount := input.GetAsIntDefault("itemCount", 4)

	// Act
	mapItems := coredynamic.NewMapAnyItems(itemCount + 5)
	collection := corestr.New.Collection.Cap(5)
	collection.Adds("val1", "val2")

	for i := 0; i < itemCount; i++ {
		mapItems.Add(fmt.Sprintf("item-%d", i), collection)
	}

	jsonResult := mapItems.JsonPtr()
	restored := coredynamic.EmptyMapAnyItems()
	parseErr := restored.JsonParseSelfInject(jsonResult)
	errcore.HandleErr(parseErr)

	newJsonResult := restored.Json()

	actual := args.Map{
		"isEqual": jsonResult.IsEqual(newJsonResult),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_MapAnyItems_GetItemRef(t *testing.T) {
	tc := mapAnyItemsGetItemRefTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	key, _ := input.GetAsString("key")

	// Act
	mapItems := coredynamic.NewMapAnyItems(10)
	collection := corestr.New.Collection.Cap(5)
	collection.Adds("x", "y", "z")
	mapItems.Add(key, collection)

	target := corestr.Empty.Collection()
	mapItems.GetItemRef(key, target)

	actual := args.Map{
		"hasItems": target.HasAnyItem(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

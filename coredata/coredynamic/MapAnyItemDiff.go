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

import (
	"log/slog"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/internal/mapdiffinternal"
)

type MapAnyItemDiff map[string]any

func (it *MapAnyItemDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it MapAnyItemDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it MapAnyItemDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it MapAnyItemDiff) LastIndex() int {
	return it.Length() - 1
}

func (it MapAnyItemDiff) AllKeysSorted() []string {
	return mapdiffinternal.MapStringAnyDiff(it.Raw()).AllKeysSorted()
}

func (it *MapAnyItemDiff) IsRawEqual(
	isRegardlessType bool,
	rightMap map[string]any,
) bool {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.
		IsRawEqual(
			isRegardlessType,
			rightMap)
}

func (it *MapAnyItemDiff) HashmapDiffUsingRaw(
	isRegardlessType bool,
	rightMap map[string]any,
) MapAnyItemDiff {
	diffMap := it.DiffRaw(
		isRegardlessType,
		rightMap)

	if len(diffMap) == 0 {
		return map[string]any{}
	}

	return diffMap
}

func (it *MapAnyItemDiff) MapAnyItems() *MapAnyItems {
	return &MapAnyItems{
		Items: it.Raw(),
	}
}

func (it *MapAnyItemDiff) HasAnyChanges(
	isRegardlessType bool,
	rightMap map[string]any,
) bool {
	return !it.IsRawEqual(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItemDiff) RawMapDiffer() mapdiffinternal.MapStringAnyDiff {
	return it.Raw()
}

func (it *MapAnyItemDiff) DiffRaw(
	isRegardlessType bool,
	rightMap map[string]any,
) map[string]any {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.
		DiffRaw(
			isRegardlessType,
			rightMap)
}

func (it *MapAnyItemDiff) DiffJsonMessage(
	isRegardlessType bool,
	rightMap map[string]any,
) string {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.DiffJsonMessage(
		isRegardlessType,
		rightMap)
}

func (it *MapAnyItemDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]any,
) (diffSlice []string) {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.ToStringsSliceOfDiffMap(
		diffMap)
}

func (it *MapAnyItemDiff) ShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]any,
) string {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.ShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItemDiff) LogShouldDiffMessage(
	isRegardlessType bool,
	title string,
	rightMap map[string]any,
) (diffMessage string) {
	differ := mapdiffinternal.
		MapStringAnyDiff(it.Raw())

	return differ.LogShouldDiffMessage(
		isRegardlessType,
		title,
		rightMap)
}

func (it *MapAnyItemDiff) Raw() map[string]any {
	if it == nil {
		return map[string]any{}
	}

	return *it
}

func (it *MapAnyItemDiff) Clear() MapAnyItemDiff {
	if it == nil {
		return map[string]any{}
	}

	*it = map[string]any{}

	return *it
}

func (it MapAnyItemDiff) Json() corejson.Result {
	return corejson.New(it)
}

func (it MapAnyItemDiff) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it MapAnyItemDiff) PrettyJsonString() string {
	return corejson.NewPtr(it).PrettyJsonString()
}

func (it MapAnyItemDiff) LogPrettyJsonString() {
	if it.IsEmpty() {
		slog.Info("empty map")
		return
	}

	prettyJson := it.JsonPtr().PrettyJsonString()

	slog.Info("map diff", "json", prettyJson)
}

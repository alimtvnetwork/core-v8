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

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/internal/mapdiffinternal"
)

type HashmapDiff map[string]string

func (it *HashmapDiff) Length() int {
	if it == nil {
		return 0
	}

	return len(*it)
}

func (it HashmapDiff) IsEmpty() bool {
	return it.Length() == 0
}

func (it HashmapDiff) HasAnyItem() bool {
	return it.Length() > 0
}

func (it HashmapDiff) LastIndex() int {
	return it.Length() - 1
}

func (it HashmapDiff) AllKeysSorted() []string {
	return mapdiffinternal.HashmapDiff(it.Raw()).AllKeysSorted()
}

func (it *HashmapDiff) MapAnyItems() map[string]any {
	if it == nil || len(*it) == 0 {
		return map[string]any{}
	}

	newMap := make(
		map[string]any,
		it.Length()+1)

	for name, value := range *it {
		newMap[name] = value
	}

	return newMap
}

func (it *HashmapDiff) HasAnyChanges(
	rightMap map[string]string,
) bool {
	return !it.IsRawEqual(
		rightMap)
}

func (it *HashmapDiff) RawMapStringAnyDiff() mapdiffinternal.MapStringAnyDiff {
	return it.MapAnyItems()
}

func (it *HashmapDiff) IsRawEqual(
	rightMap map[string]string,
) bool {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.
		IsRawEqual(rightMap)
}

func (it *HashmapDiff) HashmapDiffUsingRaw(
	rightMap map[string]string,
) HashmapDiff {
	diffMap := it.DiffRaw(
		rightMap)

	if len(diffMap) == 0 {
		return map[string]string{}
	}

	return diffMap
}

func (it *HashmapDiff) DiffRaw(
	rightMap map[string]string,
) map[string]string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.
		DiffRaw(rightMap)
}

func (it *HashmapDiff) DiffJsonMessage(
	rightMap map[string]string,
) string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.DiffJsonMessage(
		rightMap)
}

func (it *HashmapDiff) ToStringsSliceOfDiffMap(
	diffMap map[string]string,
) (diffSlice []string) {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.ToStringsSliceOfDiffMap(
		diffMap)
}

func (it *HashmapDiff) ShouldDiffMessage(
	title string,
	rightMap map[string]string,
) string {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.ShouldDiffMessage(
		title,
		rightMap)
}

func (it *HashmapDiff) LogShouldDiffMessage(
	title string,
	rightMap map[string]string,
) (diffMessage string) {
	differ := mapdiffinternal.
		HashmapDiff(it.Raw())

	return differ.LogShouldDiffMessage(
		title,
		rightMap)
}

func (it *HashmapDiff) Raw() map[string]string {
	if it == nil {
		return map[string]string{}
	}

	return *it
}

func (it *HashmapDiff) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it.Raw())
}

func (it *HashmapDiff) Deserialize(toPtr any) (parsingErr error) {
	return corejson.NewPtr(it.Raw()).Deserialize(toPtr)
}

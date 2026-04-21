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

package coreinstruction

import (
	"regexp"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
)

type BaseTags struct {
	tagsHashset *corestr.Hashset
	Tags        []string `json:"Tags,omitempty"`
}

func NewTagsPtr(tags []string) *BaseTags {
	if len(tags) == 0 {
		return NewTags(nil)
	}

	return NewTags(tags)
}

func NewTags(tags []string) *BaseTags {
	if len(tags) == 0 {
		return &BaseTags{
			Tags: []string{},
		}
	}

	return &BaseTags{
		Tags: tags,
	}
}

func (it BaseTags) TagsLength() int {
	if it.Tags == nil {
		return constants.Zero
	}

	return len(it.Tags)
}

func (it BaseTags) IsTagsEmpty() bool {
	return it.TagsLength() == 0
}

func (it *BaseTags) TagsHashset() *corestr.Hashset {
	if it.tagsHashset != nil {
		return it.tagsHashset
	}

	it.tagsHashset = corestr.New.Hashset.Strings(
		it.Tags)

	return it.tagsHashset
}

func (it BaseTags) HasAllTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := it.TagsHashset()

	return hashset.HasAll(tags...)
}

func (it BaseTags) HasAnyTags(tags ...string) bool {
	if len(tags) == 0 {
		return true
	}

	hashset := it.TagsHashset()

	return hashset.HasAny(tags...)
}

func (it BaseTags) IsAnyTagMatchesRegex(regexp2 *regexp.Regexp) bool {
	if it.IsTagsEmpty() {
		return false
	}

	for _, s := range it.Tags {
		if regexp2.MatchString(s) {
			return true
		}
	}

	return false
}

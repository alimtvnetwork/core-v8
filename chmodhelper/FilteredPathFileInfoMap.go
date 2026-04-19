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

package chmodhelper

import (
	"os"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

type FilteredPathFileInfoMap struct {
	lazyValidLocations                   []string
	lazyValidLocationFileInfoRwxWrappers []LocationFileInfoRwxWrapper
	FilesToInfoMap                       map[string]os.FileInfo
	PathsRemainsUnProcessed              []string // TODO
	MissingOrOtherPathIssues             []string
	Error                                error
}

func (it *FilteredPathFileInfoMap) LazyValidLocationFileInfoRwxWrappers() []LocationFileInfoRwxWrapper {
	if it.lazyValidLocationFileInfoRwxWrappers != nil {
		return it.lazyValidLocationFileInfoRwxWrappers
	}

	it.lazyValidLocationFileInfoRwxWrappers =
		it.ValidLocationFileInfoRwxWrappers()

	return it.lazyValidLocationFileInfoRwxWrappers
}

func (it *FilteredPathFileInfoMap) LazyValidLocations() []string {
	if it.lazyValidLocations != nil {
		return it.lazyValidLocations
	}

	it.lazyValidLocations = it.ValidLocations()

	return it.lazyValidLocations
}

func (it *FilteredPathFileInfoMap) ValidLocations() []string {
	length := len(it.FilesToInfoMap)
	slice := make([]string, length)

	if length == 0 {
		return slice
	}

	index := 0
	for s := range it.FilesToInfoMap {
		slice[index] = s
		index++
	}

	return slice
}

func (it *FilteredPathFileInfoMap) ValidFileInfos() []os.FileInfo {
	length := len(it.FilesToInfoMap)
	slice := make([]os.FileInfo, length)

	if length == 0 {
		return slice
	}

	index := 0
	for _, fileInfo := range it.FilesToInfoMap {
		slice[index] = fileInfo
		index++
	}

	return slice
}

func (it *FilteredPathFileInfoMap) ValidLocationFileInfoRwxWrappers() []LocationFileInfoRwxWrapper {
	length := len(it.FilesToInfoMap)
	slice := make([]LocationFileInfoRwxWrapper, length)

	if length == 0 {
		return slice
	}

	index := 0
	for location, fileInfo := range it.FilesToInfoMap {
		slice[index] = LocationFileInfoRwxWrapper{
			Location:   location,
			FileInfo:   fileInfo,
			RwxWrapper: New.RwxWrapper.UsingFileModePtr(fileInfo.Mode()),
		}

		index++
	}

	return slice
}

func InvalidFilteredPathFileInfoMap() *FilteredPathFileInfoMap {
	return &FilteredPathFileInfoMap{
		FilesToInfoMap:           map[string]os.FileInfo{},
		MissingOrOtherPathIssues: []string{},
		Error:                    nil,
	}
}

func (it *FilteredPathFileInfoMap) HasAnyValidFileInfo() bool {
	return len(it.FilesToInfoMap) > 0
}

func (it *FilteredPathFileInfoMap) IsEmptyValidFileInfos() bool {
	return len(it.FilesToInfoMap) == 0
}

func (it *FilteredPathFileInfoMap) LengthOfIssues() int {
	return len(it.MissingOrOtherPathIssues)
}

func (it *FilteredPathFileInfoMap) IsEmptyIssues() bool {
	return len(it.MissingOrOtherPathIssues) == 0
}

func (it *FilteredPathFileInfoMap) HasAnyIssues() bool {
	return it.Error != nil || it.HasAnyMissingPaths()
}

func (it *FilteredPathFileInfoMap) HasError() bool {
	return it.Error != nil
}

func (it *FilteredPathFileInfoMap) HasAnyMissingPaths() bool {
	return len(it.MissingOrOtherPathIssues) > 0
}

func (it *FilteredPathFileInfoMap) MissingPathsToString() string {
	return strings.Join(it.MissingOrOtherPathIssues, constants.CommaSpace)
}

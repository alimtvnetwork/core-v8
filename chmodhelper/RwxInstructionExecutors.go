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
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/errcore"
)

type RwxInstructionExecutors struct {
	items *[]*RwxInstructionExecutor
}

func NewRwxInstructionExecutors(capacity int) *RwxInstructionExecutors {
	slice := make([]*RwxInstructionExecutor, constants.Zero, capacity)

	return &RwxInstructionExecutors{
		items: &slice,
	}
}

// Add skips nil
func (it *RwxInstructionExecutors) Add(
	rwxInstructionExecutor *RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutor == nil {
		return it
	}

	*it.items = append(*it.items, rwxInstructionExecutor)

	return it
}

// Adds skips nil
func (it *RwxInstructionExecutors) Adds(
	rwxInstructionExecutors ...*RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutors == nil {
		return it
	}

	items := *it.items

	for _, executor := range rwxInstructionExecutors {
		if executor == nil {
			continue
		}

		items = append(items, executor)
	}

	*it.items = items

	return it
}

func (it *RwxInstructionExecutors) Length() int {
	if it.items == nil {
		return constants.Zero
	}

	return len(*it.items)
}

func (it *RwxInstructionExecutors) Count() int {
	return it.Length()
}

func (it *RwxInstructionExecutors) IsEmpty() bool {
	return it.Length() == 0
}

func (it *RwxInstructionExecutors) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *RwxInstructionExecutors) LastIndex() int {
	return it.Length() - 1
}

func (it *RwxInstructionExecutors) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *RwxInstructionExecutors) VerifyRwxModifiers(
	isContinueOnErr,
	isRecursiveIgnore bool,
	locations []string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if isContinueOnErr {
		return it.verifyChmodErrorContinueOnErr(
			isRecursiveIgnore,
			locations)
	}

	for _, executor := range *it.items {
		err := executor.VerifyRwxModifiers(
			isRecursiveIgnore,
			locations)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *RwxInstructionExecutors) verifyChmodErrorContinueOnErr(
	isRecursiveIgnore bool,
	locations []string,
) error {
	var sliceErr []string

	for _, executor := range *it.items {
		err := executor.VerifyRwxModifiers(
			isRecursiveIgnore,
			locations)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}
	}

	return errcore.SliceToError(sliceErr)
}

func (it *RwxInstructionExecutors) Items() *[]*RwxInstructionExecutor {
	return it.items
}

func (it *RwxInstructionExecutors) ApplyOnPath(location string) error {
	if it.IsEmpty() {
		return nil
	}

	for _, executor := range *it.items {
		err := executor.ApplyOnPath(location)

		if err != nil {
			return err

		}
	}

	return nil
}

func (it *RwxInstructionExecutors) ApplyOnPaths(locations []string) error {
	if len(locations) == 0 {
		return nil
	}

	return it.ApplyOnPathsPtr(locations)
}

func (it *RwxInstructionExecutors) ApplyOnPathsPtr(locations []string) error {
	if it.IsEmpty() {
		return nil
	}

	for _, executor := range *it.items {
		err := executor.ApplyOnPathsPtr(&locations)

		if err != nil {
			return err
		}
	}

	return nil
}

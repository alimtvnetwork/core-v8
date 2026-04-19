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

package coreinterface

type HasAnyIssueChecker interface {
	HasAnyIssues() bool
}

type HasAnyItemChecker interface {
	HasAnyItem() bool
}

type HasErrorChecker interface {
	HasError() bool
}

type HasIndexChecker interface {
	HasIndex(index int) bool
}

type HasIssueChecker interface {
	HasAnyItemChecker
	HasIssues() bool
	IsEmptyOrIssues() bool
	// HasValidItems Has items and there is no issues
	HasValidItems() bool
}

type HasIssuesOrEmptyChecker interface {
	HasIssuesOrEmpty() bool
}

type HasItemAtChecker interface {
	HasItemAt(index int) bool
}

type HasSafeItemsChecker interface {
	// HasSafeItems
	//
	// returns true if has valid item or items and no error
	HasSafeItems() bool
}

type HasFlagByNameChecker interface {
	HasFlagByName(flagName string) bool
}

type HasKeyChecker interface {
	HasKey(key string) bool
}

type HasAllKeysChecker interface {
	HasAllKeys(keys ...string) bool
}

type HasAnyKeysChecker interface {
	HasAnyKeys(keys ...string) bool
}

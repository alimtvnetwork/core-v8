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

package coreinstructiontests

import (
	"github.com/alimtvnetwork/core-v8/coreinstruction"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
)

type stringSearchIsMatchCase struct {
	name     string
	search   *coreinstruction.StringSearch
	content  string
	expected bool
}

var stringSearchIsMatchCases = []stringSearchIsMatchCase{
	{
		name: "equal match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        "hello",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "equal no match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Equal,
			Search:        "hello",
		},
		content:  "world",
		expected: false,
	},
	{
		name: "starts with match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        "hel",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "starts with no match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.StartsWith,
			Search:        "xyz",
		},
		content:  "hello",
		expected: false,
	},
	{
		name: "ends with match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.EndsWith,
			Search:        "llo",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "contains match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Contains,
			Search:        "ell",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "regex match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.Regex,
			Search:        "^he.*o$",
		},
		content:  "hello",
		expected: true,
	},
	{
		name: "not equal match",
		search: &coreinstruction.StringSearch{
			CompareMethod: stringcompareas.NotEqual,
			Search:        "hello",
		},
		content:  "world",
		expected: true,
	},
}

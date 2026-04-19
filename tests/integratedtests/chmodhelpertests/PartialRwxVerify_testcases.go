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

package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var partialRwxVerifyTestCases = []coretestcases.CaseV1{
	{
		Title: "Same input returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwx---r*x] comparing " +
			"with [-rwx-*-r*x] returns false.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx---r*x",
			"fullRwx":    "-rwx-*-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Same [-rwx-*-r*x] comparing with " +
			"[-rwx-w-r*x] returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-r*x",
			"fullRwx":    "-rwx-w-r*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwx-*-] or [-rwx-*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns true.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwx-*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Same [-rwxr*-] or [-rwxr*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns false.",
		ArrangeInput: args.Map{
			"partialRwx": "-rwxr*-",
			"fullRwx":    "-rwx-w--*x",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

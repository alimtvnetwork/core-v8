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
	"github.com/alimtvnetwork/core-v8/chmodhelper"
	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/internal/pathinternal"
)

var (
	dirCreateBasePath = pathinternal.JoinTemp("core", "case-dir-create")

	createDirTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/if/some-dir/first.txt",
						"/if/some-dir-2/first.txt",
						"/if/some-dir-3/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/if/some-dir - isCreated : true",
				"0 - 1 : core/case-dir-create/if/some-dir-2 - isCreated : true",
				"0 - 2 : core/case-dir-create/if/some-dir-3 - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirIfMissingTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if-missing",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/if-missing/first.txt",
						"/if-missing/second.txt",
						"/if-missing/third.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/if-missing - isCreated : true",
				"0 - 1 : core/case-dir-create/if-missing - isCreated : true",
				"0 - 2 : core/case-dir-create/if-missing - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirDirectTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - direct create - if exist fails",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/first.txt",
						"/f/first.txt",
						"/s/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 1 : core/case-dir-create/f/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 2 : core/case-dir-create/s/first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirByCheckingTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - direct create - by checking",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"by-checking\\a.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core/case-dir-create/by-checking/a.txt - no error during 2nd invoke of createDir.Direct",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}
)

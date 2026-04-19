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

package main

import (
	"fmt"

	"github.com/alimtvnetwork/core/chmodhelper"
)

func pathStatTest01() {
	files := currentPackageAllFiles()
	first10 := files[:10]
	first11 := append(first10, "/something-non-exist")

	items := chmodhelper.GetPathExistStats(false, first11...)

	fmt.Println(items)
}

func pathStatTest02() {
	files := currentPackageAllFiles()
	first10 := files[:10]
	first11 := append(first10, "/something-non-exist")

	items := chmodhelper.GetPathExistStats(false, first11...)

	fmt.Println(items[0].NotExistError())
	fmt.Println(items[0].NotADirError())
	fmt.Println(items[0].NotAFileError())

	fmt.Println("----------------")
	last := len(first11) - 1
	fmt.Println(items[last].NotExistError())
	fmt.Println(items[last].NotADirError())
	fmt.Println(items[last].NotAFileError())

}

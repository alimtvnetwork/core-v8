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

	"github.com/alimtvnetwork/core/keymk"
)

func testMakerTesting() {
	key := keymk.NewKey.All(keymk.BracketJoinerOption, "alim")
	key.AppendChainStrings("2", "alim-4")

	fmt.Println(key.String())

	fmt.Println(key.Compile("hello-world"))
	fmt.Println(key.Compile("hello-5-world"))
	fmt.Println(key.Compile("hello-5-world"))
	fmt.Println(key.AppendChainStrings("alim3"))
	fmt.Println(key.AppendChainStrings("alim4"))
	fmt.Println(key.String())
	fmt.Println(key.Finalized("alim{complete}"))
	fmt.Println(key.String())
	fmt.Println(key.CompileStrings("alim4new"))

	key2 := key.ClonePtr("alim4new").AppendChain("alim5new")

	fmt.Println(key2.Finalized("|BreakPoint|"))

	fmt.Println(key2.CompileKeys(key))
	// fmt.Println(key.AppendChainStrings("alim4new"))
}

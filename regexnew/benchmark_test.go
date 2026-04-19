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

package regexnew

import (
	"testing"
)

func BenchmarkCreateLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateLock(`^\d+$`)
	}
}

func BenchmarkLazyRegex_Compile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lr := &LazyRegex{
			pattern:  `^\d+$`,
			compiler: CreateLock,
		}
		lr.Compile()
	}
}

func BenchmarkLazyRegex_IsMatch_Compiled(b *testing.B) {
	lr := &LazyRegex{
		pattern:  `^\d+$`,
		compiler: CreateLock,
	}
	lr.Compile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lr.IsMatch("12345")
	}
}

func BenchmarkLazyRegex_IsMatch_Miss(b *testing.B) {
	lr := &LazyRegex{
		pattern:  `^\d+$`,
		compiler: CreateLock,
	}
	lr.Compile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lr.IsMatch("abc")
	}
}

func BenchmarkLazyRegexMap_CreateOrExistingLock_Hit(b *testing.B) {
	// prime the cache
	lazyRegexOnceMap.CreateOrExistingLock(`^bench-hit$`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lazyRegexOnceMap.CreateOrExistingLock(`^bench-hit$`)
	}
}

func BenchmarkLazyRegexMap_CreateOrExistingLock_Miss(b *testing.B) {
	patterns := make([]string, b.N)
	for i := range patterns {
		patterns[i] = `^miss-` + string(rune(i%26+'a')) + `$`
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lazyRegexOnceMap.CreateOrExistingLock(patterns[i])
	}
}

func BenchmarkIsMatchLock(b *testing.B) {
	// prime
	IsMatchLock(`^\d+$`, "123")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMatchLock(`^\d+$`, "123")
	}
}

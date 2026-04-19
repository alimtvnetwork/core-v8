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

package mutexbykey

import (
	"strconv"
	"testing"
)

func BenchmarkGet_ExistingKey(b *testing.B) {
	Get("bench-existing")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get("bench-existing")
	}
	b.StopTimer()
	Delete("bench-existing")
}

func BenchmarkGet_NewKey(b *testing.B) {
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = "bench-new-" + strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Get(keys[i])
	}
	b.StopTimer()
	for _, k := range keys {
		Delete(k)
	}
}

func BenchmarkDelete(b *testing.B) {
	keys := make([]string, b.N)
	for i := range keys {
		keys[i] = "bench-del-" + strconv.Itoa(i)
		Get(keys[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Delete(keys[i])
	}
}

func BenchmarkGet_Contention(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			Get("contention-key-" + strconv.Itoa(i%100))
			i++
		}
	})
	b.StopTimer()
	// cleanup
	for i := 0; i < 100; i++ {
		Delete("contention-key-" + strconv.Itoa(i))
	}
}

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

package stringutil

import (
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage_internal — replaceTemplate.UsingNamerMapOptions
//
// Target: replaceTemplate.go:316-341 (both branches)
// The namer interface is unexported, so this must be an internal test.
// ══════════════════════════════════════════════════════════════════════════════

type testNamerInternal struct {
	name string
}

func (n testNamerInternal) Name() string {
	return n.name
}

func Test_UsingNamerMapOptions_CurlyKeys(t *testing.T) {
	// Arrange
	rt := replaceTemplate{}
	namerMap := map[namer]string{
		testNamerInternal{name: "host"}: "example.com",
		testNamerInternal{name: "port"}: "8080",
	}

	// Act
	result := rt.UsingNamerMapOptions(
		true,
		"https://{host}:{port}/api",
		namerMap,
	)

	// Assert
	if result != "https://example.com:8080/api" {
		t.Fatalf("expected https://example.com:8080/api, got %s", result)
	}
}

func Test_UsingNamerMapOptions_DirectKeys(t *testing.T) {
	// Arrange
	rt := replaceTemplate{}
	namerMap := map[namer]string{
		testNamerInternal{name: "HOST"}: "example.com",
	}

	// Act
	result := rt.UsingNamerMapOptions(
		false,
		"https://HOST/api",
		namerMap,
	)

	// Assert
	if result != "https://example.com/api" {
		t.Fatalf("expected https://example.com/api, got %s", result)
	}
}

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

package codestack

// skippablePackages contains Go standard library and core framework
// package names. Stack traces originating from these packages are flagged
// as skippable since they add noise and are rarely useful for debugging
// application-level issues. Uses a map for O(1) lookup.
var skippablePackages = map[string]bool{
	"net":          true,
	"net/http":     true,
	"runtime":      true,
	"reflect":      true,
	"fmt":          true,
	"strings":      true,
	"strconv":      true,
	"os":           true,
	"io":           true,
	"sync":         true,
	"encoding":     true,
	"crypto":       true,
	"math":         true,
	"testing":      true,
	"log":          true,
	"bytes":        true,
	"bufio":        true,
	"context":      true,
	"database/sql": true,
	"path":         true,
	"sort":         true,
	"time":         true,
	"regexp":       true,
	"errors":       true,
	"syscall":      true,
	"unicode":      true,
}

// isSkippablePackage checks whether the given package name matches any of the
// Go standard library packages, indicating the trace is from core framework
// code rather than application code.
func isSkippablePackage(packageName string) bool {
	return skippablePackages[packageName]
}

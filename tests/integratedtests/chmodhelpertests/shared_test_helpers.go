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
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// covTempDir creates a temporary directory for test use.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covTempDir(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

// covWriteFile writes a file in a directory and returns the full path.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covWriteFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	err := os.WriteFile(p, []byte(content), 0644)
	convey.Convey("covWriteFile creates "+name, t, func() {
		convey.So(err, should.BeNil)
	})
	return p
}

// newTestRW creates a SimpleFileReaderWriter for testing.
// Moved here from Coverage12_SimpleFileRW_test.go so split-recovery subfolders can see it.
func newTestRW(dir, file string) chmodhelper.SimpleFileReaderWriter {
	return chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               filepath.Join(dir, file),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}

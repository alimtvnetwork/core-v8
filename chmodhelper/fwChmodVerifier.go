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

package chmodhelper

type fwChmodVerifier struct {
	rw *SimpleFileReaderWriter
}

func (it fwChmodVerifier) HasMismatchFile() bool {
	return ChmodVerify.IsMismatch(
		it.rw.FilePath,
		it.rw.ChmodFile)
}

func (it fwChmodVerifier) IsEqualFile() bool {
	return ChmodVerify.IsEqual(
		it.rw.FilePath,
		it.rw.ChmodFile)
}

func (it fwChmodVerifier) HasMismatchParentDir() bool {
	return ChmodVerify.IsMismatch(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) IsEqualParentDir() bool {
	return ChmodVerify.IsEqual(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) MismatchErrorParentDir() error {
	return ChmodVerify.MismatchError(
		it.rw.ParentDir,
		it.rw.ChmodDir)
}

func (it fwChmodVerifier) MismatchErrorFile() error {
	return ChmodVerify.MismatchError(
		it.rw.FilePath,
		it.rw.ChmodFile)
}

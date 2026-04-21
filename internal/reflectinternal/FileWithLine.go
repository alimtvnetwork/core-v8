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

package reflectinternal

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/internal/jsoninternal"
)

type FileWithLine struct {
	FilePath string // absolute file reflectPath
	Line     int    // line number
}

func (it *FileWithLine) FullFilePath() string {
	return it.FilePath
}

func (it *FileWithLine) LineNumber() int {
	return it.Line
}

func (it *FileWithLine) IsNil() bool {
	return it == nil
}

func (it *FileWithLine) IsNotNil() bool {
	return it != nil
}

func (it *FileWithLine) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return it.FileWithLine()
}

func (it FileWithLine) StringUsingFmt(formatterFunc func(fileWithLine FileWithLine) string) string {
	return formatterFunc(it)
}

func (it *FileWithLine) FileWithLine() string {
	return fmt.Sprintf(
		fileWithLineFormat,
		it.FilePath,
		it.Line,
	)
}

func (it FileWithLine) JsonModel() FileWithLine {
	return it
}

func (it *FileWithLine) JsonModelAny() any {
	return it.JsonModel()
}

func (it *FileWithLine) JsonString() string {
	s, _ := jsoninternal.Pretty.AnyTo.String(it)

	return s
}

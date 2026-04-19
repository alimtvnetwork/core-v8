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

package loggerinf

import "io"

type Committer interface {
	Commit()
}

type BytesCompiler interface {
	CompileBytes() ([]byte, error)
}

type BytesCompilerIf interface {
	CompileBytesIf(isCompile bool) ([]byte, error)
}

type MustBytesCompiler interface {
	CompileBytesMust() []byte
}

type StringFinalizer interface {
	Finalize() string
}

type IfStringCompiler interface {
	CompileIf(isCompile bool) string
}

type Compiler interface {
	Compile() string
}

type FmtCompiler interface {
	CompileFmt(formatter string, v ...any) string
}

type Serializer interface {
	Serialize() ([]byte, error)
}

type NewGeneralWriter interface {
	NewGeneralWriter(writeConfigurer WriterConfigurer) io.Writer
}

type Configurer interface {
	LoggerTyper() LoggerTyper
	StackSkipIndex() int
}

type WriterConfigurer interface {
	Configurer
	AdditionalConfigProcessor
}

type AdditionalConfigProcessor interface {
	AdditionalConfigBytes() []byte
	AdditionalConfigProcess() error
}

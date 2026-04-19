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

package internalpathextender

// PathInfoer
//
// FullPath:
//   - Refers to the FullPath (dir or filename through absolute path).
//
// FileName:
//   - Refers to the file name at the end
//
// DirName:
//   - Refers to the file name at the end
//
// Name:
//   - Refers to the end of name could be file or dir.
//
// Extension:
//   - Refers to dot extension (.db, .back etc)
//
// RootDir:
//   - Refers to root dir where things started from.
//
// Relative:
//   - Refers to relative from RootDir
//
// ParentDir:
//   - Refers to parent dir of the FullPath and different from RootDir
type PathInfoer interface {
	FullPath() string
	FileName() string
	DirName() string
	Name() string
	// Extension
	//
	//  Refers to dot extension always
	Extension() string
	// RootDir
	//
	//  Refers to start of the dir
	//  For example a repo start point
	RootDir() string
	// Relative
	//
	//  Refers to relative path from root
	//  Joining Root + Relative should give absolute or FullPath()
	Relative() string
	// ParentDir
	//
	//  Refers to current full-path's parent dir.
	//  ParentDir is different from RootDir.
	ParentDir() string
	// String
	//
	//  returns full summary info
	String() string

	Size() uint64
}

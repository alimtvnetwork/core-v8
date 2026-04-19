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

package enumtype

var (
	rangesMap = [...]string{
		Invalid:           "Invalid",
		Boolean:           "Boolean",
		Byte:              "Byte",
		UnsignedInteger16: "UnsignedInteger16",
		UnsignedInteger32: "UnsignedInteger32",
		UnsignedInteger64: "UnsignedInteger64",
		Integer8:          "Integer8",
		Integer16:         "Integer16",
		Integer32:         "Integer32",
		Integer64:         "Integer64",
		Integer:           "Integer",
		String:            "String",
	}

	stringToVariantMap = map[string]Variant{
		"Invalid":           Invalid,
		"Boolean":           Boolean,
		"Byte":              Byte,
		"UnsignedInteger16": UnsignedInteger16,
		"UnsignedInteger32": UnsignedInteger32,
		"UnsignedInteger64": UnsignedInteger64,
		"Integer8":          Integer8,
		"Integer16":         Integer16,
		"Integer32":         Integer32,
		"Integer64":         Integer64,
		"Integer":           Integer,
		"String":            String,
	}

	unSignedMap = map[Variant]bool{
		Byte:              true,
		UnsignedInteger16: true,
		UnsignedInteger32: true,
		UnsignedInteger64: true,
	}

	numbersMap = map[Variant]bool{
		Byte:              true,
		UnsignedInteger16: true,
		UnsignedInteger32: true,
		UnsignedInteger64: true,
		Integer8:          true,
		Integer16:         true,
		Integer32:         true,
		Integer64:         true,
		Integer:           true,
	}

	integersMap = map[Variant]bool{
		Integer8:  true,
		Integer16: true,
		Integer32: true,
		Integer64: true,
		Integer:   true,
	}
)

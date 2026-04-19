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

package coreinterface

// ReflectSetter
//
// ReflectSetTo
//
//	sets current object to something else by casting,
//	reflection, by unmarshalling or by marshalling
//
// # Set any object from to toPointer object
//
// Valid Inputs or Supported (https://t.ly/SGWUx):
//   - From, To: (null, null)                          -- do nothing
//   - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
//   - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
//   - From, To: ([]byte, otherType)                   -- try unmarshal, reflect
//   - From, To: (otherType, *[]byte)                  -- try marshal, reflect
//
// Validations:
//   - Check null, if both null no error return quickly.
//   - NotSupported returns as error.
//   - NotSupported: (from, to) - (..., not pointer)
//   - NotSupported: (from, to) - (null, notNull)
//   - NotSupported: (from, to) - (notNull, null)
//   - NotSupported: (from, to) - not same type and not bytes on any
//   - `From` null or nil is not supported and will return error.
//
// Reference:
//   - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
//   - Method document screenshot    : https://prnt.sc/26dmf5g
type ReflectSetter interface {
	// ReflectSetTo
	//
	// ReflectSetter
	//  sets current object to something else by casting,
	//  reflection, by unmarshalling or by marshalling
	//
	// Set any object from to toPointer object
	//
	// Valid Inputs or Supported (https://t.ly/SGWUx):
	//  - From, To: (null, null)                          -- do nothing
	//  - From, To: (sameTypePointer, sameTypePointer)    -- try reflection
	//  - From, To: (sameTypeNonPointer, sameTypePointer) -- try reflection
	//  - From, To: ([]byte, otherType)                   -- try unmarshal, reflect
	//  - From, To: (otherType, *[]byte)                  -- try marshal, reflect
	//
	// Validations:
	//  - Check null, if both null no error return quickly.
	//  - NotSupported returns as error.
	//      - NotSupported: (from, to) - (..., not pointer)
	//      - NotSupported: (from, to) - (null, notNull)
	//      - NotSupported: (from, to) - (notNull, null)
	//      - NotSupported: (from, to) - not same type and not bytes on any
	//  - `From` null or nil is not supported and will return error.
	//
	// Reference:
	//  - Reflection String Set Example : https://go.dev/play/p/fySLYuOvoRK.go?download=true
	//  - Method document screenshot    : https://prnt.sc/26dmf5g
	ReflectSetTo(toPointer any) error
}

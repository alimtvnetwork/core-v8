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

package corestrtests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// =============================================================================
// SimpleStringOnce — JSON/Serialization coverage gaps (15 methods)
// =============================================================================

func Test_SimpleStringOnce_Json(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Json", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		r := sso.Json()

		// Act
		actual := args.Map{"noErr": r.Error == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO Json no error", actual)
	})
}

func Test_SimpleStringOnce_JsonPtr(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_JsonPtr", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		r := sso.JsonPtr()

		// Act
		actual := args.Map{
			"nonNil": r != nil,
			"noErr": r.Error == nil,
		}

		// Assert
		expected := args.Map{
			"nonNil": true,
			"noErr": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO JsonPtr", actual)
	})
}

func Test_SimpleStringOnce_JsonModel_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_JsonModel", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		m := sso.JsonModel()

		// Act
		actual := args.Map{
			"val": m.Value,
			"init": m.IsInitialize,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"init": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO JsonModel", actual)
	})
}

func Test_SimpleStringOnce_JsonModelAny_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_JsonModelAny", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("x")

		// Act
		actual := args.Map{"nonNil": sso.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonModelAny", actual)
	})
}

func Test_SimpleStringOnce_MarshalJSON(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_MarshalJSON", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, err := sso.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO MarshalJSON", actual)
	})
}

func Test_SimpleStringOnce_UnmarshalJSON(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_UnmarshalJSON", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, _ := sso.MarshalJSON()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		err := sso2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO UnmarshalJSON", actual)
	})
}

func Test_SimpleStringOnce_UnmarshalJSON_Error(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_UnmarshalJSON_Error", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("")
		err := sso.UnmarshalJSON([]byte("invalid"))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SSO UnmarshalJSON error", actual)
	})
}

func Test_SimpleStringOnce_ParseInjectUsingJson_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ParseInjectUsingJson", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		r, err := sso2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonNil": r != nil,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonNil": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJson", actual)
	})
}

func Test_SimpleStringOnce_ParseInjectUsingJson_Error(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ParseInjectUsingJson_Error", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("")
		jr := &corejson.Result{Error: errors.New("fail")}
		_, err := sso.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJson error", actual)
	})
}

func Test_SimpleStringOnce_ParseInjectUsingJsonMust_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ParseInjectUsingJsonMust", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		r := sso2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"nonNil": r != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJsonMust", actual)
	})
}

func Test_SimpleStringOnce_ParseInjectUsingJsonMust_Panics(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_ParseInjectUsingJsonMust_Panics", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("")
		jr := &corejson.Result{Error: errors.New("fail")}
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			sso.ParseInjectUsingJsonMust(jr)
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "SSO ParseInjectUsingJsonMust panics", actual)
	})
}

func Test_SimpleStringOnce_JsonParseSelfInject_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_JsonParseSelfInject", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		jr := sso.JsonPtr()
		sso2 := corestr.New.SimpleStringOnce.Init("")
		err := sso2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO JsonParseSelfInject", actual)
	})
}

func Test_SimpleStringOnce_AsJsonContractsBinder_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_AsJsonContractsBinder", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"nonNil": sso.AsJsonContractsBinder() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonContractsBinder", actual)
	})
}

func Test_SimpleStringOnce_AsJsoner_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_AsJsoner", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"nonNil": sso.AsJsoner() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsoner", actual)
	})
}

func Test_SimpleStringOnce_AsJsonParseSelfInjector_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_AsJsonParseSelfInjector", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"nonNil": sso.AsJsonParseSelfInjector() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonParseSelfInjector", actual)
	})
}

func Test_SimpleStringOnce_AsJsonMarshaller_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_AsJsonMarshaller", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"nonNil": sso.AsJsonMarshaller() != nil}

		// Assert
		expected := args.Map{"nonNil": true}
		expected.ShouldBeEqual(t, 0, "SSO AsJsonMarshaller", actual)
	})
}

func Test_SimpleStringOnce_Serialize_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Serialize", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		b, err := sso.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"nonEmpty": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"nonEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SSO Serialize", actual)
	})
}

func Test_SimpleStringOnce_Deserialize_SsoJsonRemaininggaps(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_Deserialize", func() {
		// Arrange
		sso := corestr.New.SimpleStringOnce.Init("hello")
		var target corestr.SimpleStringOnceModel
		err := sso.Deserialize(&target)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "SSO Deserialize", actual)
	})
}

// =============================================================================
// LinkedCollectionNode — Next, isNextEqual, isNextChainEqual
// =============================================================================

func Test_LCN_Next(t *testing.T) {
	safeTest(t, "Test_LCN_Next", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		head := lc.Head()

		// Act
		actual := args.Map{"hasNext": head.Next() != nil}

		// Assert
		expected := args.Map{"hasNext": true}
		expected.ShouldBeEqual(t, 0, "LCN Next", actual)
	})
}

func Test_LCN_Next_Nil(t *testing.T) {
	safeTest(t, "Test_LCN_Next_Nil", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Create()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		head := lc.Head()

		// Act
		actual := args.Map{"noNext": head.Next() == nil}

		// Assert
		expected := args.Map{"noNext": true}
		expected.ShouldBeEqual(t, 0, "LCN Next nil tail", actual)
	})
}

func Test_LCN_IsChainEqual_SameChain(t *testing.T) {
	safeTest(t, "Test_LCN_IsChainEqual_SameChain", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc1.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"b"}))

		// Act
		actual := args.Map{"eq": lc1.Head().IsChainEqual(lc2.Head())}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "LCN IsChainEqual same", actual)
	})
}

func Test_LCN_IsChainEqual_Different(t *testing.T) {
	safeTest(t, "Test_LCN_IsChainEqual_Different", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc1.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2.Add(corestr.New.Collection.Strings([]string{"c"}))

		// Act
		actual := args.Map{"eq": lc1.Head().IsChainEqual(lc2.Head())}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "LCN IsChainEqual diff", actual)
	})
}

func Test_LCN_IsEqual_EqualElements(t *testing.T) {
	safeTest(t, "Test_LCN_IsEqual_EqualElements", func() {
		// Arrange
		lc1 := corestr.New.LinkedCollection.Create()
		lc1.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc2 := corestr.New.LinkedCollection.Create()
		lc2.Add(corestr.New.Collection.Strings([]string{"a"}))

		// Act
		actual := args.Map{"eq": lc1.Head().IsEqual(lc2.Head())}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "LCN IsEqual same elements", actual)
	})
}

// =============================================================================
// KeyValueCollection — KeysHashset (panics with "implement me")
// =============================================================================

func Test_KVC_KeysHashset_Panics(t *testing.T) {
	safeTest(t, "Test_KVC_KeysHashset_Panics", func() {
		// Arrange
		kvc := corestr.New.KeyValues.Empty()
		kvc.Add("k", "v")
		panicked := false
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicked = true
				}
			}()
			kvc.KeysHashset()
		}()

		// Act
		actual := args.Map{"panicked": panicked}

		// Assert
		expected := args.Map{"panicked": true}
		expected.ShouldBeEqual(t, 0, "KVC KeysHashset panics", actual)
	})
}

// =============================================================================
// CollectionsOfCollection — JsonPtr
// =============================================================================

func Test_COC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_COC_JsonPtr", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.Empty()
		coc.Add(corestr.New.Collection.Strings([]string{"a", "b"}))
		r := coc.JsonPtr()

		// Act
		actual := args.Map{
			"nonNil": r != nil,
			"noErr": r.Error == nil,
		}

		// Assert
		expected := args.Map{
			"nonNil": true,
			"noErr": true,
		}
		expected.ShouldBeEqual(t, 0, "COC JsonPtr", actual)
	})
}

// =============================================================================
// HashsetsCollection — JsonPtr
// =============================================================================

func Test_HC_JsonPtr(t *testing.T) {
	safeTest(t, "Test_HC_JsonPtr", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Empty()
		hs := corestr.New.Hashset.Strings([]string{"x", "y"})
		hc.Add(hs)
		r := hc.JsonPtr()

		// Act
		actual := args.Map{
			"nonNil": r != nil,
			"noErr": r.Error == nil,
		}

		// Assert
		expected := args.Map{
			"nonNil": true,
			"noErr": true,
		}
		expected.ShouldBeEqual(t, 0, "HC JsonPtr", actual)
	})
}

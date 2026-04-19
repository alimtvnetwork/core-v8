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
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── SimpleSlice ──

func Test_SS_Add(t *testing.T) {
	safeTest(t, "Test_SS_Add", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()
		ss.Add("a")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS_AddSplit(t *testing.T) {
	safeTest(t, "Test_SS_AddSplit", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddSplit("a,b", ",")
	})
}

func Test_SS_AddIf(t *testing.T) {
	safeTest(t, "Test_SS_AddIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddIf(false, "skip")
		ss.AddIf(true, "add")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SS_Adds(t *testing.T) {
	safeTest(t, "Test_SS_Adds", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.Adds("a", "b")
		ss.Adds()
	})
}

func Test_SS_Append(t *testing.T) {
	safeTest(t, "Test_SS_Append", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.Append("a")
		ss.Append()
	})
}

func Test_SS_AppendFmt(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmt", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AppendFmt("hello %s", "world")
		ss.AppendFmt("")
	})
}

func Test_SS_AppendFmtIf(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmtIf", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AppendFmtIf(true, "x%d", 1)
		ss.AppendFmtIf(false, "skip")
	})
}

func Test_SS_Length(t *testing.T) {
	safeTest(t, "Test_SS_Length", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		var n *corestr.SimpleSlice
		actual = args.Map{"result": n.Length() != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SS_IsEmpty_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_IsEmpty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Empty()

		// Act
		actual := args.Map{"result": ss.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SS_HasAnyItem_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_HasAnyItem", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.HasAnyItem()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SS_First(t *testing.T) {
	safeTest(t, "Test_SS_First", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.First() != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_SS_Last(t *testing.T) {
	safeTest(t, "Test_SS_Last", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a", "b")

		// Act
		actual := args.Map{"result": ss.Last() != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_SS_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_SS_FirstOrDefault", func() {
		ss := corestr.New.SimpleSlice.Empty()
		_ = ss.FirstOrDefault()
		ss.Add("a")
		_ = ss.FirstOrDefault()
	})
}

func Test_SS_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_SS_LastOrDefault", func() {
		ss := corestr.New.SimpleSlice.Empty()
		_ = ss.LastOrDefault()
		ss.Add("a")
		_ = ss.LastOrDefault()
	})
}

func Test_SS_Strings_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Strings", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Strings()
	})
}

func Test_SS_Join(t *testing.T) {
	safeTest(t, "Test_SS_Join", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.Join(",")
	})
}

func Test_SS_JoinLine_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_JoinLine", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.JoinLine()
	})
}

func Test_SS_Take(t *testing.T) {
	safeTest(t, "Test_SS_Take", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		_ = ss.Take(2)
		_ = ss.Take(5)
	})
}

func Test_SS_Skip(t *testing.T) {
	safeTest(t, "Test_SS_Skip", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b", "c")
		_ = ss.Skip(1)
	})
}

func Test_SS_Collection_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Collection", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Collection(false)
	})
}

func Test_SS_Hashset_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Hashset", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Hashset()
	})
}

func Test_SS_Clear_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Clear", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Clear()
	})
}

func Test_SS_Dispose_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Dispose", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		ss.Dispose()
	})
}

func Test_SS_Clone_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Clone", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.Clone(false)
	})
}

func Test_SS_DeepClone_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_DeepClone", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.DeepClone()
	})
}

func Test_SS_AddAsTitleValue_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_AddAsTitleValue", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddAsTitleValue("key", "val")
	})
}

func Test_SS_AddAsTitleValueIf(t *testing.T) {
	safeTest(t, "Test_SS_AddAsTitleValueIf", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddAsTitleValueIf(true, "key", "val")
		ss.AddAsTitleValueIf(false, "key", "val")
	})
}

func Test_SS_AddAsCurlyTitleWrap(t *testing.T) {
	safeTest(t, "Test_SS_AddAsCurlyTitleWrap", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddAsCurlyTitleWrap("title", "content")
	})
}

func Test_SS_AddAsCurlyTitleWrapIf(t *testing.T) {
	safeTest(t, "Test_SS_AddAsCurlyTitleWrapIf", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddAsCurlyTitleWrapIf(true, "title", "content")
		ss.AddAsCurlyTitleWrapIf(false, "title", "content")
	})
}

func Test_SS_AddError(t *testing.T) {
	safeTest(t, "Test_SS_AddError", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddError(nil)
	})
}

func Test_SS_AddIf2(t *testing.T) {
	safeTest(t, "Test_SS_AddIf2", func() {
		ss := corestr.New.SimpleSlice.Empty()
		ss.AddIf(true, "a")
		ss.AddIf(false, "c")
	})
}

func Test_SS_Sort_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Sort", func() {
		ss := corestr.New.SimpleSlice.Lines("b", "a")
		_ = ss.Sort()
	})
}

func Test_SS_Reverse_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_Reverse", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		ss.Reverse()
	})
}

func Test_SS_IsContains(t *testing.T) {
	safeTest(t, "Test_SS_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Lines("a")

		// Act
		actual := args.Map{"result": ss.IsContains("a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SS_IndexOf_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_IndexOf", func() {
		ss := corestr.New.SimpleSlice.Lines("a", "b")
		_ = ss.IndexOf("a")
	})
}

func Test_SS_HasIndex_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SS_HasIndex", func() {
		ss := corestr.New.SimpleSlice.Lines("a")
		_ = ss.HasIndex(0)
		_ = ss.HasIndex(99)
	})
}

// ── newSimpleSliceCreator ──

func Test_NSSC_Cap(t *testing.T)     { _ = corestr.New.SimpleSlice.Cap(5) }
func Test_NSSC_Default(t *testing.T) { _ = corestr.New.SimpleSlice.Default() }
func Test_NSSC_Deserialize(t *testing.T) {
	safeTest(t, "Test_NSSC_Deserialize", func() {
		// Arrange
		_, _ = corestr.New.SimpleSlice.Deserialize([]byte(`["a"]`))
		_, _ = corestr.New.SimpleSlice.Deserialize([]byte(`invalid`))
	})
}
func Test_NSSC_DeserializeJsoner(t *testing.T) {
	safeTest(t, "Test_NSSC_DeserializeJsoner", func() {
		r := corejson.New([]string{"a"})
		_, _ = corestr.New.SimpleSlice.DeserializeJsoner(&r)
	})
}
func Test_NSSC_UsingLines(t *testing.T) {
	safeTest(t, "Test_NSSC_UsingLines", func() {
		_ = corestr.New.SimpleSlice.UsingLines(true, "a")
		_ = corestr.New.SimpleSlice.UsingLines(false, "a")
		_ = corestr.New.SimpleSlice.UsingLines(false)
	})
}
func Test_NSSC_Lines(t *testing.T)         { _ = corestr.New.SimpleSlice.Lines("a") }
func Test_NSSC_Split(t *testing.T)         { _ = corestr.New.SimpleSlice.Split("a,b", ",") }
func Test_NSSC_SplitLines(t *testing.T)    { _ = corestr.New.SimpleSlice.SplitLines("a\nb") }
func Test_NSSC_SpreadStrings(t *testing.T) { _ = corestr.New.SimpleSlice.SpreadStrings("a") }
func Test_NSSC_Hashset(t *testing.T) {
	safeTest(t, "Test_NSSC_Hashset", func() {
		_ = corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.StringsSpreadItems("a"))
		_ = corestr.New.SimpleSlice.Hashset(corestr.New.Hashset.Empty())
	})
}
func Test_NSSC_Map(t *testing.T) {
	safeTest(t, "Test_NSSC_Map", func() {
		_ = corestr.New.SimpleSlice.Map(map[string]string{"k": "v"})
		_ = corestr.New.SimpleSlice.Map(map[string]string{})
	})
}
func Test_NSSC_Create(t *testing.T)        { _ = corestr.New.SimpleSlice.Create([]string{"a"}) }
func Test_NSSC_Strings(t *testing.T)       { _ = corestr.New.SimpleSlice.Strings([]string{"a"}) }
func Test_NSSC_StringsPtr(t *testing.T)    { _ = corestr.New.SimpleSlice.StringsPtr(nil) }
func Test_NSSC_StringsOptions(t *testing.T) {
	safeTest(t, "Test_NSSC_StringsOptions", func() {
		_ = corestr.New.SimpleSlice.StringsOptions(true, []string{"a"})
		_ = corestr.New.SimpleSlice.StringsOptions(false, []string{})
	})
}
func Test_NSSC_StringsClone(t *testing.T) {
	safeTest(t, "Test_NSSC_StringsClone", func() {
		_ = corestr.New.SimpleSlice.StringsClone([]string{"a"})
		_ = corestr.New.SimpleSlice.StringsClone(nil)
	})
}
func Test_NSSC_Direct(t *testing.T) {
	safeTest(t, "Test_NSSC_Direct", func() {
		_ = corestr.New.SimpleSlice.Direct(true, []string{"a"})
		_ = corestr.New.SimpleSlice.Direct(false, nil)
	})
}
func Test_NSSC_UsingSeparatorLine(t *testing.T) { _ = corestr.New.SimpleSlice.UsingSeparatorLine(",", "a,b") }
func Test_NSSC_UsingLine(t *testing.T)          { _ = corestr.New.SimpleSlice.UsingLine("a|b") }
func Test_NSSC_Empty(t *testing.T)              { _ = corestr.New.SimpleSlice.Empty() }
func Test_NSSC_ByLen(t *testing.T)              { _ = corestr.New.SimpleSlice.ByLen([]string{"a"}) }

// ── SimpleStringOnce ──

func Test_SSO_Value(t *testing.T) {
	safeTest(t, "Test_SSO_Value", func() {
		sso := corestr.New.SimpleStringOnce.Init("hello")

		// Act
		actual := args.Map{"result": sso.Value() != "hello"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
	})
}

func Test_SSO_IsInitialized_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_IsInitialized", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		actual := args.Map{"result": sso.IsInitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SSO_IsDefined_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_IsDefined", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		actual := args.Map{"result": sso.IsDefined()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SSO_IsUninitialized_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_IsUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		actual := args.Map{"result": sso.IsUninitialized()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SSO_Invalidate_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Invalidate", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		sso.Invalidate()
	})
}

func Test_SSO_Reset_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Reset", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		sso.Reset()
	})
}

func Test_SSO_IsInvalid_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_IsInvalid", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		actual := args.Map{"result": sso.IsInvalid()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SSO_ValueBytes_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_ValueBytes", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		_ = sso.ValueBytes()
		_ = sso.ValueBytesPtr()
	})
}

func Test_SSO_SetOnUninitialized_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_SetOnUninitialized", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.SetOnUninitialized("v")
		err := sso.SetOnUninitialized("v2")
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SSO_GetSetOnce_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_GetSetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		v := sso.GetSetOnce("hello")
		actual := args.Map{"result": v != "hello"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello", actual)
		v2 := sso.GetSetOnce("world")
		actual = args.Map{"result": v2 != "hello"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected hello still", actual)
	})
}

func Test_SSO_GetOnce_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_GetOnce", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		_ = sso.GetOnce()
	})
}

func Test_SSO_SetInitialize(t *testing.T) {
	safeTest(t, "Test_SSO_SetInitialize", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		sso.SetInitialize()
	})
}

func Test_SSO_String_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_String", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		_ = sso.String()
	})
}

func Test_SSO_Dispose_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Dispose", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		sso.Dispose()
	})
}

func Test_SSO_Boolean_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Boolean", func() {
		sso := corestr.New.SimpleStringOnce.Init("true")
		_ = sso.Boolean(true)
	})
}

func Test_SSO_Int_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Int", func() {
		sso := corestr.New.SimpleStringOnce.Init("42")
		_ = sso.Int()
	})
}

func Test_SSO_ValueDefFloat64_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_ValueDefFloat64", func() {
		sso := corestr.New.SimpleStringOnce.Init("3.14")
		_ = sso.ValueDefFloat64()
	})
}

func Test_SSO_JsonMethods(t *testing.T) {
	safeTest(t, "Test_SSO_JsonMethods", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		_ = sso.Json()
		_ = sso.JsonPtr()
		_ = sso.JsonModel()
		_ = sso.JsonModelAny()
		_, _ = sso.MarshalJSON()
		_ = sso.AsJsonContractsBinder()
		_ = sso.AsJsoner()
		_ = sso.AsJsonParseSelfInjector()
		_ = sso.AsJsonMarshaller()
		_, _ = sso.Serialize()
	})
}

func Test_SSO_UnmarshalJSON_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_UnmarshalJSON", func() {
		sso := corestr.SimpleStringOnce{}
		_ = sso.UnmarshalJSON([]byte(`"hello"`))
	})
}

func Test_SSO_ParseInjectUsingJson_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_ParseInjectUsingJson", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		r := corejson.New("hello")
		_, _ = sso.ParseInjectUsingJson(&r)
	})
}

func Test_SSO_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	safeTest(t, "Test_SSO_ParseInjectUsingJsonMust_Panic", func() {
		defer func() { recover() }()
		sso := corestr.New.SimpleStringOnce.Empty()
		bad := corejson.NewResult.UsingString(`invalid`)
		sso.ParseInjectUsingJsonMust(bad)
	})
}

func Test_SSO_JsonParseSelfInject_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_JsonParseSelfInject", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		r := corejson.New("hello")
		_ = sso.JsonParseSelfInject(&r)
	})
}

func Test_SSO_Deserialize_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Deserialize", func() {
		sso := corestr.New.SimpleStringOnce.Init("x")
		var s string
		_ = sso.Deserialize(&s)
	})
}

func Test_SSO_IsEmpty_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_IsEmpty", func() {
		sso := corestr.New.SimpleStringOnce.Empty()
		_ = sso.IsEmpty()
	})
}

func Test_SSO_IsWhitespace(t *testing.T) {
	safeTest(t, "Test_SSO_IsWhitespace", func() {
		sso := corestr.New.SimpleStringOnce.Init("  ")
		_ = sso.IsWhitespace()
	})
}

func Test_SSO_Trim_SimplesliceSso(t *testing.T) {
	safeTest(t, "Test_SSO_Trim", func() {
		sso := corestr.New.SimpleStringOnce.Init("  x  ")
		_ = sso.Trim()
	})
}

// ── newSimpleStringOnceCreator ──

func Test_NSSOC_Any(t *testing.T)      { _ = corestr.New.SimpleStringOnce.Any(false, 42, true) }
func Test_NSSOC_Create(t *testing.T)    { _ = corestr.New.SimpleStringOnce.Create("x", true) }
func Test_NSSOC_CreatePtr(t *testing.T) { _ = corestr.New.SimpleStringOnce.CreatePtr("x", true) }
func Test_NSSOC_Empty(t *testing.T)      { _ = corestr.New.SimpleStringOnce.Empty() }
func Test_NSSOC_Init(t *testing.T)       { _ = corestr.New.SimpleStringOnce.Init("x") }
func Test_NSSOC_InitPtr(t *testing.T)    { _ = corestr.New.SimpleStringOnce.InitPtr("x") }
func Test_NSSOC_Uninitialized(t *testing.T) { _ = corestr.New.SimpleStringOnce.Uninitialized("x") }

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
	"sync"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── LinkedList ──

func Test_LL_Empty(t *testing.T) {
	safeTest(t, "Test_LL_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
	})
}

func Test_LL_Add_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_Add", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("a")
		ll.Add("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LL_Head(t *testing.T) {
	safeTest(t, "Test_LL_Head", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LL_Tail(t *testing.T) {
	safeTest(t, "Test_LL_Tail", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Tail().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LL_LengthLock_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_LengthLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.LengthLock()
	})
}

func Test_LL_IsEquals(t *testing.T) {
	safeTest(t, "Test_LL_IsEquals", func() {
		// Arrange
		a := corestr.New.LinkedList.Strings([]string{"a"})
		b := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_LL_IsEqualsWithSensitive(t *testing.T) {
	safeTest(t, "Test_LL_IsEqualsWithSensitive", func() {
		// Arrange
		a := corestr.New.LinkedList.Strings([]string{"A"})
		b := corestr.New.LinkedList.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": a.IsEqualsWithSensitive(b, false)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_LL_IsEquals_Nil(t *testing.T) {
	safeTest(t, "Test_LL_IsEquals_Nil", func() {
		// Arrange
		var a, b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_LL_AddBack(t *testing.T) {
	safeTest(t, "Test_LL_AddBack", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.PushBack("a")
		ll.PushBack("b")
	})
}

func Test_LL_AddLock_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_AddLock", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.AddLock("a")
	})
}

func Test_LL_AddCollection_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_AddCollection", func() {
		ll := corestr.New.LinkedList.Empty()
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(c)
		ll.AddCollection(corestr.New.Collection.Empty())
	})
}

func Test_LL_AddStrings_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_AddStrings", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.AddStrings([]string{"a"})
		ll.AddStrings(nil)
	})
}

func Test_LL_Adds_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_Adds", func() {
		ll := corestr.New.LinkedList.Empty()
		ll.Adds("a", "b")
	})
}

func Test_LL_AddAfterNode(t *testing.T) {
	safeTest(t, "Test_LL_AddAfterNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.AddAfterNode(ll.Head(), "x")
	})
}

func Test_LL_Remove(t *testing.T) {
	safeTest(t, "Test_LL_Remove", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByElementValue("b", true, false)
		ll.RemoveNodeByElementValue("a", true, false)
		ll.RemoveNodeByElementValue("c", true, false)
	})
}

func Test_LL_RemoveAt(t *testing.T) {
	safeTest(t, "Test_LL_RemoveAt", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b", "c"})
		ll.RemoveNodeByIndex(1)
		ll.RemoveNodeByIndex(0)
	})
}

func Test_LL_List_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		r := ll.List()

		// Act
		actual := args.Map{"result": len(r) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		ll2 := corestr.New.LinkedList.Empty()
		_ = ll2.List()
	})
}

func Test_LL_ListLock_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_ListLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.ListLock()
	})
}

func Test_LL_ToCollection_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		c := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": c == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil collection", actual)
	})
}

func Test_LL_Loop_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_Loop", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			return false
		})
	})
}

func Test_LL_String_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.String()
		_ = corestr.New.LinkedList.Empty().String()
	})
}

func Test_LL_StringLock_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_StringLock", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.StringLock()
	})
}

func Test_LL_Clear_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_Clear", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Clear()
	})
}

func Test_LL_JsonMethods(t *testing.T) {
	safeTest(t, "Test_LL_JsonMethods", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Json()
		_ = ll.JsonPtr()
		_ = ll.JsonModel()
		_ = ll.JsonModelAny()
		_, _ = ll.MarshalJSON()
		_ = ll.AsJsonMarshaller()
	})
}

func Test_LL_AddCollectionToNode_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.AddCollectionToNode(false, ll.Head(), corestr.New.Collection.Strings([]string{"b"}))
	})
}

func Test_LL_AddStringsPtrToNode_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LL_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)
	})
}

// ── newLinkedListCreator ──

func Test_NLLC_Empty(t *testing.T)   { _ = corestr.New.LinkedList.Empty() }
func Test_NLLC_Create(t *testing.T)  { _ = corestr.New.LinkedList.Create() }
func Test_NLLC_Strings(t *testing.T) { _ = corestr.New.LinkedList.Strings([]string{"a"}) }
func Test_NLLC_SpreadStrings(t *testing.T) {
	safeTest(t, "Test_NLLC_SpreadStrings", func() {
		// Arrange
		_ = corestr.New.LinkedList.SpreadStrings("a", "b")
		_ = corestr.New.LinkedList.SpreadStrings()
	})
}
func Test_NLLC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_NLLC_PointerStringsPtr", func() {
		s := "a"
		_ = corestr.New.LinkedList.PointerStringsPtr(&[]*string{&s})
		_ = corestr.New.LinkedList.PointerStringsPtr(nil)
	})
}

// ── LinkedCollections ──

func Test_LC_Empty(t *testing.T) {
	safeTest(t, "Test_LC_Empty", func() {
		lc := corestr.New.LinkedCollection.Empty()

		// Act
		actual := args.Map{"result": lc.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_LC_Add_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_Add", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		actual := args.Map{"result": lc.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LC_AddLock_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddLock", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.AddLock(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_LC_AddAsync_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddAsync", func() {
		lc := corestr.New.LinkedCollection.Empty()
		var wg sync.WaitGroup
		wg.Add(1)
		lc.AddAsync(corestr.New.Collection.Strings([]string{"a"}), &wg)
		wg.Wait()
	})
}

func Test_LC_AddAnother_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddAnother", func() {
		lc := corestr.New.LinkedCollection.Empty()
		other := corestr.New.LinkedCollection.Strings("a")
		lc.AddAnother(other)
		lc.AddAnother(corestr.New.LinkedCollection.Empty())
	})
}

func Test_LC_AddStrings_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddStrings", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.AddStrings("a")
	})
}

func Test_LC_AddAfterNode_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddAfterNode", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddAfterNode(lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
	})
}

func Test_LC_PushBack_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_PushBack", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.PushBack(corestr.New.Collection.Strings([]string{"a"}))
	})
}

func Test_LC_AllIndividualItemsLength_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AllIndividualItemsLength", func() {
		lc := corestr.New.LinkedCollection.Strings("a", "b")
		_ = lc.AllIndividualItemsLength()
	})
}

func Test_LC_First(t *testing.T) {
	safeTest(t, "Test_LC_First", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.First()
	})
}

func Test_LC_Single_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_Single", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Single()
	})
}

func Test_LC_Last(t *testing.T) {
	safeTest(t, "Test_LC_Last", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Last()
	})
}

func Test_LC_FirstOrDefault(t *testing.T) {
	safeTest(t, "Test_LC_FirstOrDefault", func() {
		lc := corestr.New.LinkedCollection.Empty()
		_ = lc.FirstOrDefault()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.FirstOrDefault()
	})
}

func Test_LC_LastOrDefault(t *testing.T) {
	safeTest(t, "Test_LC_LastOrDefault", func() {
		lc := corestr.New.LinkedCollection.Empty()
		_ = lc.LastOrDefault()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		_ = lc.LastOrDefault()
	})
}

func Test_LC_RemoveNodeByIndex(t *testing.T) {
	safeTest(t, "Test_LC_RemoveNodeByIndex", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.Add(corestr.New.Collection.Strings([]string{"b"}))
		lc.RemoveNodeByIndex(0)
	})
}

func Test_LC_List_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_List", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.List()
	})
}

func Test_LC_Loop_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_Loop", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Loop(func(arg *corestr.LinkedCollectionProcessorParameter) bool {
			return false
		})
	})
}

func Test_LC_String_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_String", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.String()
		_ = corestr.New.LinkedCollection.Empty().String()
	})
}

func Test_LC_Clear_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_Clear", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Clear()
	})
}

func Test_LC_JsonMethods(t *testing.T) {
	safeTest(t, "Test_LC_JsonMethods", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Json()
		_ = lc.JsonPtr()
		_ = lc.JsonModel()
		_ = lc.JsonModelAny()
		_, _ = lc.MarshalJSON()
		_ = lc.AsJsonContractsBinder()
		_ = lc.AsJsoner()
		_ = lc.AsJsonMarshaller()
		_ = lc.AsJsonParseSelfInjector()
	})
}

func Test_LC_SimpleSlice_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_SimpleSlice", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.SimpleSlice()
	})
}

func Test_LC_AddCollectionToNode_LinkedlistLinkedcoll(t *testing.T) {
	safeTest(t, "Test_LC_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Empty()
		lc.Add(corestr.New.Collection.Strings([]string{"a"}))
		lc.AddCollectionToNode(false, lc.Head(), corestr.New.Collection.Strings([]string{"x"}))
	})
}

// ── newLinkedListCollectionsCreator ──

func Test_NLLCC_Empty(t *testing.T)  { _ = corestr.New.LinkedCollection.Empty() }
func Test_NLLCC_Create(t *testing.T) { _ = corestr.New.LinkedCollection.Create() }
func Test_NLLCC_Strings(t *testing.T) {
	safeTest(t, "Test_NLLCC_Strings", func() {
		_ = corestr.New.LinkedCollection.Strings("a")
	})
}
func Test_NLLCC_PointerStringsPtr(t *testing.T) {
	safeTest(t, "Test_NLLCC_PointerStringsPtr", func() {
		s := "a"
		_ = corestr.New.LinkedCollection.PointerStringsPtr(&[]*string{&s})
		_ = corestr.New.LinkedCollection.PointerStringsPtr(nil)
	})
}
func Test_NLLCC_UsingCollections(t *testing.T) {
	safeTest(t, "Test_NLLCC_UsingCollections", func() {
		_ = corestr.New.LinkedCollection.UsingCollections(
			corestr.New.Collection.Strings([]string{"a"}),
		)
		_ = corestr.New.LinkedCollection.UsingCollections()
	})
}

// ── LinkedListNode ──

func Test_LLN_HasNext(t *testing.T) {
	safeTest(t, "Test_LLN_HasNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		actual := args.Map{"result": ll.Head().HasNext()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_LLN_Next(t *testing.T) {
	safeTest(t, "Test_LLN_Next", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_ = ll.Head().Next()
	})
}

func Test_LLN_EndOfChain(t *testing.T) {
	safeTest(t, "Test_LLN_EndOfChain", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_, l := ll.Head().EndOfChain()
		actual := args.Map{"result": l != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LLN_LoopEndOfChain(t *testing.T) {
	safeTest(t, "Test_LLN_LoopEndOfChain", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a", "b"})
		_, _ = ll.Head().LoopEndOfChain(func(arg *corestr.LinkedListProcessorParameter) bool {
			return false
		})
	})
}

func Test_LLN_Clone(t *testing.T) {
	safeTest(t, "Test_LLN_Clone", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Head().Clone()
	})
}

func Test_LLN_String(t *testing.T) {
	safeTest(t, "Test_LLN_String", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		_ = ll.Head().String()
	})
}

func Test_LLN_AddNext(t *testing.T) {
	safeTest(t, "Test_LLN_AddNext", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddNext(ll, "x")
	})
}

func Test_LLN_AddNextNode(t *testing.T) {
	safeTest(t, "Test_LLN_AddNextNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddNextNode(ll, &corestr.LinkedListNode{Element: "x"})
	})
}

func Test_LLN_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_LLN_AddStringsToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddStringsToNode(ll, false, []string{"b"})
	})
}

func Test_LLN_AddStringsPtrToNode(t *testing.T) {
	safeTest(t, "Test_LLN_AddStringsPtrToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		items := []string{"b"}
		ll.Head().AddStringsPtrToNode(ll, false, &items)
	})
}

func Test_LLN_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_LLN_AddCollectionToNode", func() {
		ll := corestr.New.LinkedList.Strings([]string{"a"})
		ll.Head().AddCollectionToNode(ll, false, corestr.New.Collection.Strings([]string{"b"}))
	})
}

// ── LinkedCollectionNode ──

func Test_LCN_IsEmpty(t *testing.T) {
	safeTest(t, "Test_LCN_IsEmpty", func() {
		node := &corestr.LinkedCollectionNode{}
		_ = node.IsEmpty()
	})
}

func Test_LCN_HasElement(t *testing.T) {
	safeTest(t, "Test_LCN_HasElement", func() {
		node := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"a"})}
		_ = node.HasElement()
	})
}

func Test_LCN_AddNext(t *testing.T) {
	safeTest(t, "Test_LCN_AddNext", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddNext(lc, corestr.New.Collection.Strings([]string{"x"}))
	})
}

func Test_LCN_AddNextNode(t *testing.T) {
	safeTest(t, "Test_LCN_AddNextNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		newNode := &corestr.LinkedCollectionNode{Element: corestr.New.Collection.Strings([]string{"x"})}
		lc.Head().AddNextNode(lc, newNode)
	})
}

func Test_LCN_AddStringsToNode(t *testing.T) {
	safeTest(t, "Test_LCN_AddStringsToNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddStringsToNode(lc, false, []string{"b"}, false)
	})
}

func Test_LCN_AddCollectionToNode(t *testing.T) {
	safeTest(t, "Test_LCN_AddCollectionToNode", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		lc.Head().AddCollectionToNode(lc, false, corestr.New.Collection.Strings([]string{"b"}))
	})
}

func Test_LCN_Clone(t *testing.T) {
	safeTest(t, "Test_LCN_Clone", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Head().Clone()
	})
}

func Test_LCN_String(t *testing.T) {
	safeTest(t, "Test_LCN_String", func() {
		lc := corestr.New.LinkedCollection.Strings("a")
		_ = lc.Head().String()
	})
}

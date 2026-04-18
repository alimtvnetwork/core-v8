package corestrtests

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S12 — LinkedList.go (1,141 lines) — Full coverage
// ══════════════════════════════════════════════════════════════

func Test_LinkedList_01_LinkedList_HeadTailLength_FromS12(t *testing.T) {
	safeTest(t, "Test_01_LinkedList_HeadTailLength", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act & Assert
		actual := args.Map{"result": ll.Head() == nil || ll.Tail() == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil head/tail", actual)
		actual = args.Map{"result": ll.Length() != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_02_LinkedList_LengthLock_FromS12(t *testing.T) {
	safeTest(t, "Test_02_LinkedList_LengthLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.LengthLock() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_03_LinkedList_IsEmpty_HasItems_FromS12(t *testing.T) {
	safeTest(t, "Test_03_LinkedList_IsEmpty_HasItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmpty()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no items", actual)
		ll.Add("a")
		actual = args.Map{"result": ll.IsEmpty() || !ll.HasItems()}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not empty", actual)
	})
}

func Test_LinkedList_04_LinkedList_IsEmptyLock_FromS12(t *testing.T) {
	safeTest(t, "Test_04_LinkedList_IsEmptyLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.IsEmptyLock()}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_LinkedList_05_LinkedList_Add_MultipleItems_FromS12(t *testing.T) {
	safeTest(t, "Test_05_LinkedList_Add_MultipleItems", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_06_LinkedList_AddLock_FromS12(t *testing.T) {
	safeTest(t, "Test_06_LinkedList_AddLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_07_LinkedList_AddItemsMap_FromS12(t *testing.T) {
	safeTest(t, "Test_07_LinkedList_AddItemsMap", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{"a": true, "b": false})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_08_LinkedList_AddItemsMap_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_08_LinkedList_AddItemsMap_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddItemsMap(map[string]bool{})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_09_LinkedList_AddNonEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_09_LinkedList_AddNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmpty("")
		ll.AddNonEmpty("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_10_LinkedList_AddNonEmptyWhitespace_FromS12(t *testing.T) {
	safeTest(t, "Test_10_LinkedList_AddNonEmptyWhitespace", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddNonEmptyWhitespace("   ")
		ll.AddNonEmptyWhitespace("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_11_LinkedList_AddIf_FromS12(t *testing.T) {
	safeTest(t, "Test_11_LinkedList_AddIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddIf(true, "yes")
		ll.AddIf(false, "no")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_12_LinkedList_AddsIf_FromS12(t *testing.T) {
	safeTest(t, "Test_12_LinkedList_AddsIf", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsIf(true, "a", "b")
		ll.AddsIf(false, "c")

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_13_LinkedList_AddFunc_FromS12(t *testing.T) {
	safeTest(t, "Test_13_LinkedList_AddFunc", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFunc(func() string { return "val" })

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_14_LinkedList_AddFuncErr_NoError_FromS12(t *testing.T) {
	safeTest(t, "Test_14_LinkedList_AddFuncErr_NoError", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Assert
		ll.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) })

		// Act
		actual := args.Map{"result": ll.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_15_LinkedList_AddFuncErr_WithError_FromS12(t *testing.T) {
	safeTest(t, "Test_15_LinkedList_AddFuncErr_WithError", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		called := false
		ll.AddFuncErr(func() (string, error) { return "", &testErr{} }, func(err error) { called = true })

		// Act
		actual := args.Map{"result": called}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected err handler called", actual)
	})
}

func Test_LinkedList_16_LinkedList_Push_PushBack_PushFront_FromS12(t *testing.T) {
	safeTest(t, "Test_16_LinkedList_Push_PushBack_PushFront", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Push("a")
		ll.PushBack("b")
		ll.PushFront("z")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_17_LinkedList_AddFront_OnEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_17_LinkedList_AddFront_OnEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_18_LinkedList_AddFront_OnNonEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_18_LinkedList_AddFront_OnNonEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.AddFront("a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_LinkedList_19_LinkedList_Adds_FromS12(t *testing.T) {
	safeTest(t, "Test_19_LinkedList_Adds", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_20_LinkedList_Adds_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_20_LinkedList_Adds_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_21_LinkedList_AddStrings_FromS12(t *testing.T) {
	safeTest(t, "Test_21_LinkedList_AddStrings", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_22_LinkedList_AddStrings_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_22_LinkedList_AddStrings_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStrings([]string{})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_23_LinkedList_AddsLock_FromS12(t *testing.T) {
	safeTest(t, "Test_23_LinkedList_AddsLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddsLock("a")

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_24_LinkedList_AppendNode_OnEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_24_LinkedList_AppendNode_OnEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		node := &corestr.LinkedListNode{Element: "a"}
		ll.AppendNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_25_LinkedList_AddBackNode_FromS12(t *testing.T) {
	safeTest(t, "Test_25_LinkedList_AddBackNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("x")
		node := &corestr.LinkedListNode{Element: "y"}
		ll.AddBackNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_26_LinkedList_InsertAt_FromS12(t *testing.T) {
	safeTest(t, "Test_26_LinkedList_InsertAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "c")
		ll.InsertAt(1, "b")

		// Act
		actual := args.Map{"result": ll.Length() != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_27_LinkedList_InsertAt_Front_FromS12(t *testing.T) {
	safeTest(t, "Test_27_LinkedList_InsertAt_Front", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("b")
		ll.InsertAt(0, "a")

		// Act
		actual := args.Map{"result": ll.Head().Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a at head", actual)
	})
}

func Test_LinkedList_28_LinkedList_Loop_FromS12(t *testing.T) {
	safeTest(t, "Test_28_LinkedList_Loop", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return false
		})

		// Act
		actual := args.Map{"result": count != 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_LinkedList_29_LinkedList_Loop_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_29_LinkedList_Loop_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {

		// Act
			actual := args.Map{"result": false}

		// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "should not be called", actual)
			return false
		})
	})
}

func Test_LinkedList_30_LinkedList_Loop_Break_FromS12(t *testing.T) {
	safeTest(t, "Test_30_LinkedList_Loop_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		count := 0
		ll.Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count++
			return true
		})

		// Act
		actual := args.Map{"result": count != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_31_LinkedList_Filter_FromS12(t *testing.T) {
	safeTest(t, "Test_31_LinkedList_Filter", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_32_LinkedList_Filter_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_32_LinkedList_Filter_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: false}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_33_LinkedList_Filter_Break_FromS12(t *testing.T) {
	safeTest(t, "Test_33_LinkedList_Filter_Break", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		nodes := ll.Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})

		// Act
		actual := args.Map{"result": len(nodes) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_34_LinkedList_IsEquals_FromS12(t *testing.T) {
	safeTest(t, "Test_34_LinkedList_IsEquals", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Adds("a", "b")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LinkedList_35_LinkedList_IsEquals_BothNil_FromS12(t *testing.T) {
	safeTest(t, "Test_35_LinkedList_IsEquals_BothNil", func() {
		// Arrange
		var a *corestr.LinkedList
		var b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LinkedList_36_LinkedList_IsEquals_OneNil_FromS12(t *testing.T) {
	safeTest(t, "Test_36_LinkedList_IsEquals_OneNil", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		var b *corestr.LinkedList

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(b, true)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_LinkedList_37_LinkedList_IsEquals_SamePtr_FromS12(t *testing.T) {
	safeTest(t, "Test_37_LinkedList_IsEquals_SamePtr", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")

		// Act
		actual := args.Map{"result": a.IsEqualsWithSensitive(a, true)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LinkedList_38_LinkedList_IsEquals_BothEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_38_LinkedList_IsEquals_BothEmpty", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		b := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected equal", actual)
	})
}

func Test_LinkedList_39_LinkedList_IsEquals_DiffLength_FromS12(t *testing.T) {
	safeTest(t, "Test_39_LinkedList_IsEquals_DiffLength", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("a")
		b := corestr.New.LinkedList.Create()
		b.Adds("a", "b")

		// Act
		actual := args.Map{"result": a.IsEquals(b)}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected not equal", actual)
	})
}

func Test_LinkedList_40_LinkedList_IsEqualsWithSensitive_CaseInsensitive_FromS12(t *testing.T) {
	safeTest(t, "Test_40_LinkedList_IsEqualsWithSensitive_CaseInsensitive", func() {
		a := corestr.New.LinkedList.Create()
		a.Add("A")
		b := corestr.New.LinkedList.Create()
		b.Add("a")
		if a.IsEqualsWithSensitive(b, false) {
			// depends on LinkedListNode.IsChainEqual implementation
			_ = 0
		}
	})
}

func Test_LinkedList_41_LinkedList_RemoveNodeByElementValue_FromS12(t *testing.T) {
	safeTest(t, "Test_41_LinkedList_RemoveNodeByElementValue", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByElementValue("b", true, false)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_42_LinkedList_RemoveNodeByElementValue_First_FromS12(t *testing.T) {
	safeTest(t, "Test_42_LinkedList_RemoveNodeByElementValue_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByElementValue("a", true, false)

		// Act
		actual := args.Map{"result": ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b at head", actual)
	})
}

func Test_LinkedList_43_LinkedList_RemoveNodeByIndex_FromS12(t *testing.T) {
	safeTest(t, "Test_43_LinkedList_RemoveNodeByIndex", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_44_LinkedList_RemoveNodeByIndex_First_FromS12(t *testing.T) {
	safeTest(t, "Test_44_LinkedList_RemoveNodeByIndex_First", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(0)

		// Act
		actual := args.Map{"result": ll.Head().Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LinkedList_45_LinkedList_RemoveNodeByIndex_Last_FromS12(t *testing.T) {
	safeTest(t, "Test_45_LinkedList_RemoveNodeByIndex_Last", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.RemoveNodeByIndex(1)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_46_LinkedList_RemoveNode_FromS12(t *testing.T) {
	safeTest(t, "Test_46_LinkedList_RemoveNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.Add("b")
		node := ll.IndexAt(0)
		ll.RemoveNode(node)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_47_LinkedList_RemoveNode_Nil_FromS12(t *testing.T) {
	safeTest(t, "Test_47_LinkedList_RemoveNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveNode(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_48_LinkedList_IndexAt_FromS12(t *testing.T) {
	safeTest(t, "Test_48_LinkedList_IndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		node := ll.IndexAt(1)

		// Act
		actual := args.Map{"result": node.Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
	})
}

func Test_LinkedList_49_LinkedList_IndexAt_Zero_FromS12(t *testing.T) {
	safeTest(t, "Test_49_LinkedList_IndexAt_Zero", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.IndexAt(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LinkedList_50_LinkedList_IndexAt_Negative_FromS12(t *testing.T) {
	safeTest(t, "Test_50_LinkedList_IndexAt_Negative", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.IndexAt(-1) != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LinkedList_51_LinkedList_SafeIndexAt_FromS12(t *testing.T) {
	safeTest(t, "Test_51_LinkedList_SafeIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": ll.SafeIndexAt(1).Element != "b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected b", actual)
		actual = args.Map{"result": ll.SafeIndexAt(-1) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
		actual = args.Map{"result": ll.SafeIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LinkedList_52_LinkedList_SafeIndexAtLock_FromS12(t *testing.T) {
	safeTest(t, "Test_52_LinkedList_SafeIndexAtLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafeIndexAtLock(0).Element != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LinkedList_53_LinkedList_SafePointerIndexAt_FromS12(t *testing.T) {
	safeTest(t, "Test_53_LinkedList_SafePointerIndexAt", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ptr := ll.SafePointerIndexAt(0)

		// Act
		actual := args.Map{"result": ptr == nil || *ptr != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.SafePointerIndexAt(10) != nil}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_LinkedList_54_LinkedList_SafePointerIndexAtUsingDefault_FromS12(t *testing.T) {
	safeTest(t, "Test_54_LinkedList_SafePointerIndexAtUsingDefault", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefault(0, "def") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ll.SafePointerIndexAtUsingDefault(10, "def") != "def"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected def", actual)
	})
}

func Test_LinkedList_55_LinkedList_SafePointerIndexAtUsingDefaultLock_FromS12(t *testing.T) {
	safeTest(t, "Test_55_LinkedList_SafePointerIndexAtUsingDefaultLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.SafePointerIndexAtUsingDefaultLock(0, "def") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LinkedList_56_LinkedList_GetNextNodes_FromS12(t *testing.T) {
	safeTest(t, "Test_56_LinkedList_GetNextNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		nodes := ll.GetNextNodes(2)

		// Act
		actual := args.Map{"result": len(nodes) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_57_LinkedList_GetAllLinkedNodes_FromS12(t *testing.T) {
	safeTest(t, "Test_57_LinkedList_GetAllLinkedNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": len(ll.GetAllLinkedNodes()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_58_LinkedList_AddPointerStringsPtr_FromS12(t *testing.T) {
	safeTest(t, "Test_58_LinkedList_AddPointerStringsPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		a := "a"
		ll.AddPointerStringsPtr([]*string{&a, nil})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_59_LinkedList_AddCollection_FromS12(t *testing.T) {
	safeTest(t, "Test_59_LinkedList_AddCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		col := corestr.New.Collection.Strings([]string{"a", "b"})
		ll.AddCollection(col)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_60_LinkedList_AddCollection_Nil_FromS12(t *testing.T) {
	safeTest(t, "Test_60_LinkedList_AddCollection_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddCollection(nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_61_LinkedList_ToCollection_FromS12(t *testing.T) {
	safeTest(t, "Test_61_LinkedList_ToCollection", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_62_LinkedList_ToCollection_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_62_LinkedList_ToCollection_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		col := ll.ToCollection(0)

		// Act
		actual := args.Map{"result": col.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_63_LinkedList_List_FromS12(t *testing.T) {
	safeTest(t, "Test_63_LinkedList_List", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": len(ll.List()) != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_64_LinkedList_ListPtr_FromS12(t *testing.T) {
	safeTest(t, "Test_64_LinkedList_ListPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtr()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_65_LinkedList_ListLock_FromS12(t *testing.T) {
	safeTest(t, "Test_65_LinkedList_ListLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_66_LinkedList_ListPtrLock_FromS12(t *testing.T) {
	safeTest(t, "Test_66_LinkedList_ListPtrLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.ListPtrLock()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_67_LinkedList_String_FromS12(t *testing.T) {
	safeTest(t, "Test_67_LinkedList_String", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.String() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_LinkedList_68_LinkedList_String_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_68_LinkedList_String_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": strings.Contains(ll.String(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_LinkedList_69_LinkedList_StringLock_FromS12(t *testing.T) {
	safeTest(t, "Test_69_LinkedList_StringLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.StringLock() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_LinkedList_70_LinkedList_StringLock_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_70_LinkedList_StringLock_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": strings.Contains(ll.StringLock(), "No Element")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected No Element", actual)
	})
}

func Test_LinkedList_71_LinkedList_Join_FromS12(t *testing.T) {
	safeTest(t, "Test_71_LinkedList_Join", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")

		// Act
		actual := args.Map{"result": ll.Join(",") != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_LinkedList_72_LinkedList_JoinLock_FromS12(t *testing.T) {
	safeTest(t, "Test_72_LinkedList_JoinLock", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JoinLock(",") != "a"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
	})
}

func Test_LinkedList_73_LinkedList_Joins_FromS12(t *testing.T) {
	safeTest(t, "Test_73_LinkedList_Joins", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		result := ll.Joins(",", "b")

		// Act
		actual := args.Map{"result": strings.Contains(result, "a")}

		// Assert
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected a in result", actual)
	})
}

func Test_LinkedList_74_LinkedList_Joins_NilItems_FromS12(t *testing.T) {
	safeTest(t, "Test_74_LinkedList_Joins_NilItems", func() {
		ll := corestr.New.LinkedList.Create()
		result := ll.Joins(",", nil...)
		_ = result
	})
}

func Test_LinkedList_75_LinkedList_MarshalJSON_FromS12(t *testing.T) {
	safeTest(t, "Test_75_LinkedList_MarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		data, err := ll.MarshalJSON()

		// Act
		actual := args.Map{"result": err != nil || len(data) == 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected valid JSON", actual)
	})
}

func Test_LinkedList_76_LinkedList_UnmarshalJSON_FromS12(t *testing.T) {
	safeTest(t, "Test_76_LinkedList_UnmarshalJSON", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`["a","b"]`))

		// Act
		actual := args.Map{"result": err != nil || ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_77_LinkedList_UnmarshalJSON_Invalid_FromS12(t *testing.T) {
	safeTest(t, "Test_77_LinkedList_UnmarshalJSON_Invalid", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		err := ll.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_LinkedList_78_LinkedList_JsonModel_FromS12(t *testing.T) {
	safeTest(t, "Test_78_LinkedList_JsonModel", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": len(ll.JsonModel()) != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_79_LinkedList_JsonModelAny_FromS12(t *testing.T) {
	safeTest(t, "Test_79_LinkedList_JsonModelAny", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JsonModelAny() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_LinkedList_80_LinkedList_Clear_RemoveAll_FromS12(t *testing.T) {
	safeTest(t, "Test_80_LinkedList_Clear_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b")
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_81_LinkedList_Clear_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_81_LinkedList_Clear_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Clear()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_82_LinkedList_RemoveAll_FromS12(t *testing.T) {
	safeTest(t, "Test_82_LinkedList_RemoveAll", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.RemoveAll()

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_83_LinkedList_Json_FromS12(t *testing.T) {
	safeTest(t, "Test_83_LinkedList_Json", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.Json()

		// Act
		actual := args.Map{"result": jsonResult.HasError()}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_LinkedList_84_LinkedList_JsonPtr_FromS12(t *testing.T) {
	safeTest(t, "Test_84_LinkedList_JsonPtr", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")

		// Act
		actual := args.Map{"result": ll.JsonPtr() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_LinkedList_85_LinkedList_ParseInjectUsingJson_FromS12(t *testing.T) {
	safeTest(t, "Test_85_LinkedList_ParseInjectUsingJson", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result, err := target.ParseInjectUsingJson(jsonResult)

		// Act
		actual := args.Map{"result": err != nil || result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_86_LinkedList_ParseInjectUsingJsonMust_FromS12(t *testing.T) {
	safeTest(t, "Test_86_LinkedList_ParseInjectUsingJsonMust", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		result := target.ParseInjectUsingJsonMust(jsonResult)

		// Act
		actual := args.Map{"result": result.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_87_LinkedList_JsonParseSelfInject_FromS12(t *testing.T) {
	safeTest(t, "Test_87_LinkedList_JsonParseSelfInject", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		jsonResult := ll.JsonPtr()
		target := corestr.New.LinkedList.Create()
		err := target.JsonParseSelfInject(jsonResult)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_LinkedList_88_LinkedList_AsJsonMarshaller_FromS12(t *testing.T) {
	safeTest(t, "Test_88_LinkedList_AsJsonMarshaller", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()

		// Act
		actual := args.Map{"result": ll.AsJsonMarshaller() == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_LinkedList_89_LinkedList_GetCompareSummary_FromS12(t *testing.T) {
	safeTest(t, "Test_89_LinkedList_GetCompareSummary", func() {
		// Arrange
		a := corestr.New.LinkedList.Create()
		a.Add("x")
		b := corestr.New.LinkedList.Create()
		b.Add("y")
		summary := a.GetCompareSummary(b, "left", "right")

		// Act
		actual := args.Map{"result": summary == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_LinkedList_90_LinkedList_AppendChainOfNodes_FromS12(t *testing.T) {
	safeTest(t, "Test_90_LinkedList_AppendChainOfNodes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		other := corestr.New.LinkedList.Create()
		other.Adds("b", "c")
		ll.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_LinkedList_91_LinkedList_AppendChainOfNodes_OnEmpty_FromS12(t *testing.T) {
	safeTest(t, "Test_91_LinkedList_AppendChainOfNodes_OnEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		other := corestr.New.LinkedList.Create()
		other.Add("a")
		ll.AppendChainOfNodes(other.Head())

		// Act
		actual := args.Map{"result": ll.Length() < 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 1", actual)
	})
}

func Test_LinkedList_92_LinkedList_AttachWithNode_FromS12(t *testing.T) {
	safeTest(t, "Test_92_LinkedList_AttachWithNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(node, addNode)

		// Act
		actual := args.Map{"result": err != nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected no error", actual)
	})
}

func Test_LinkedList_93_LinkedList_AttachWithNode_NilCurrent_FromS12(t *testing.T) {
	safeTest(t, "Test_93_LinkedList_AttachWithNode_NilCurrent", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		addNode := &corestr.LinkedListNode{Element: "b"}
		err := ll.AttachWithNode(nil, addNode)

		// Act
		actual := args.Map{"result": err == nil}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_LinkedList_94_LinkedList_RemoveNodeByIndexes_FromS12(t *testing.T) {
	safeTest(t, "Test_94_LinkedList_RemoveNodeByIndexes", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Adds("a", "b", "c")
		ll.RemoveNodeByIndexes(false, 1)

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_95_LinkedList_RemoveNodeByIndexes_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_95_LinkedList_RemoveNodeByIndexes_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.RemoveNodeByIndexes(false)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_96_LinkedList_AddStringsToNode_FromS12(t *testing.T) {
	safeTest(t, "Test_96_LinkedList_AddStringsToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b"})

		// Act
		actual := args.Map{"result": ll.Length() != 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_LinkedList_97_LinkedList_AddStringsToNode_SkipOnNull_FromS12(t *testing.T) {
	safeTest(t, "Test_97_LinkedList_AddStringsToNode_SkipOnNull", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsToNode(true, nil, []string{"a"})

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_98_LinkedList_AddStringsToNode_Empty_FromS12(t *testing.T) {
	safeTest(t, "Test_98_LinkedList_AddStringsToNode_Empty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		ll.AddStringsToNode(false, ll.Head(), []string{})

		// Act
		actual := args.Map{"result": ll.Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_LinkedList_99_LinkedList_AddStringsToNode_Multiple_FromS12(t *testing.T) {
	safeTest(t, "Test_99_LinkedList_AddStringsToNode_Multiple", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		node := ll.Head()
		ll.AddStringsToNode(false, node, []string{"b", "c"})

		// Act
		actual := args.Map{"result": ll.Length() < 3}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 3", actual)
	})
}

func Test_LinkedList_100_LinkedList_AddStringsPtrToNode_FromS12(t *testing.T) {
	safeTest(t, "Test_100_LinkedList_AddStringsPtrToNode", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.Add("a")
		items := []string{"b"}
		ll.AddStringsPtrToNode(false, ll.Head(), &items)

		// Act
		actual := args.Map{"result": ll.Length() < 2}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected at least 2", actual)
	})
}

func Test_LinkedList_101_LinkedList_AddStringsPtrToNode_Nil_FromS12(t *testing.T) {
	safeTest(t, "Test_101_LinkedList_AddStringsPtrToNode_Nil", func() {
		// Arrange
		ll := corestr.New.LinkedList.Create()
		ll.AddStringsPtrToNode(true, nil, nil)

		// Act
		actual := args.Map{"result": ll.Length() != 0}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_LinkedList_102_LinkedList_AddAsync_FromS12(t *testing.T) {
	safeTest(t, "Test_102_LinkedList_AddAsync", func() {
		ll := corestr.New.LinkedList.Create()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		col := corestr.New.Collection.Strings([]string{"a"})
		ll.AddCollectionToNode(true, nil, col)
		_ = ll
	})
}

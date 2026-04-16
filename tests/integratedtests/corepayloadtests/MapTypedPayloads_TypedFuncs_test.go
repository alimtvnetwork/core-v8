package corepayloadtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corepayload"
	"github.com/alimtvnetwork/core/coretests/args"
)

// testUserTyped, makeTypedWrapper, makeTypedCollection are defined in shared_typed_helpers.go

// ── MapTypedPayloads ──

func Test_MapTypedPayloads(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	names := corepayload.MapTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) string {
			return item.Data().Name
		},
	)

	// Assert
	actual := args.Map{
		"length": len(names),
		"first": names[0],
	}
	expected := args.Map{
		"length": 3,
		"first": "Alice",
	}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns names -- 3 items", actual)
}

func Test_MapTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	result := corepayload.MapTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) string {
			return item.Data().Name
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloads returns empty -- empty source", actual)
}

// ── MapTypedPayloadData ──

func Test_MapTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	emails := corepayload.MapTypedPayloadData[testUserTyped, string](col,
		func(u testUserTyped) string { return u.Email },
	)

	// Assert
	actual := args.Map{
		"length": len(emails),
		"first": emails[0],
	}
	expected := args.Map{
		"length": 3,
		"first": "a@a.com",
	}
	expected.ShouldBeEqual(t, 0, "MapTypedPayloadData returns emails -- 3 items", actual)
}

// ── FlatMapTypedPayloads ──

func Test_FlatMapTypedPayloads(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.FlatMapTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) []string {
			return []string{item.Data().Name, item.Data().Email}
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 6}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloads returns flattened -- 3x2", actual)
}

func Test_FlatMapTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	result := corepayload.FlatMapTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) []string {
			return []string{item.Data().Name}
		},
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloads returns empty -- empty source", actual)
}

// ── FlatMapTypedPayloadData ──

func Test_FlatMapTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.FlatMapTypedPayloadData[testUserTyped, string](col,
		func(u testUserTyped) []string { return []string{u.Name} },
	)

	// Assert
	actual := args.Map{"length": len(result)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "FlatMapTypedPayloadData returns names -- 3 items", actual)
}

// ── ReduceTypedPayloads ──

func Test_ReduceTypedPayloads(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	totalAge := corepayload.ReduceTypedPayloads[testUserTyped, int](col, 0,
		func(acc int, item *corepayload.TypedPayloadWrapper[testUserTyped]) int {
			return acc + item.Data().Age
		},
	)

	// Assert
	actual := args.Map{"totalAge": totalAge}
	expected := args.Map{"totalAge": 90}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns sum -- 30+25+35", actual)
}

func Test_ReduceTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	result := corepayload.ReduceTypedPayloads[testUserTyped, int](col, 42,
		func(acc int, item *corepayload.TypedPayloadWrapper[testUserTyped]) int {
			return acc + 1
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloads returns initial -- empty", actual)
}

// ── ReduceTypedPayloadData ──

func Test_ReduceTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.ReduceTypedPayloadData[testUserTyped, int](col, 0,
		func(acc int, u testUserTyped) int { return acc + u.Age },
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 90}
	expected.ShouldBeEqual(t, 0, "ReduceTypedPayloadData returns sum -- 90", actual)
}

// ── GroupTypedPayloads ──

func Test_GroupTypedPayloads(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	groups := corepayload.GroupTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) string {
			if item.Data().Age >= 30 {
				return "senior"
			}
			return "junior"
		},
	)

	// Assert
	actual := args.Map{
		"groupCount":   len(groups),
		"seniorCount":  groups["senior"].Length(),
		"juniorCount":  groups["junior"].Length(),
	}
	expected := args.Map{
		"groupCount":   2,
		"seniorCount":  2,
		"juniorCount":  1,
	}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns grouped -- 2 groups", actual)
}

func Test_GroupTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	groups := corepayload.GroupTypedPayloads[testUserTyped, string](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) string { return "x" },
	)

	// Assert
	actual := args.Map{"count": len(groups)}
	expected := args.Map{"count": 0}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloads returns empty -- empty source", actual)
}

// ── GroupTypedPayloadData ──

func Test_GroupTypedPayloadData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	groups := corepayload.GroupTypedPayloadData[testUserTyped, string](col,
		func(u testUserTyped) string {
			if u.Age >= 30 {
				return "old"
			}
			return "young"
		},
	)

	// Assert
	actual := args.Map{"groupCount": len(groups)}
	expected := args.Map{"groupCount": 2}
	expected.ShouldBeEqual(t, 0, "GroupTypedPayloadData returns grouped -- 2", actual)
}

// ── PartitionTypedPayloads ──

func Test_PartitionTypedPayloads(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	matching, notMatching := corepayload.PartitionTypedPayloads[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
			return item.Data().Age >= 30
		},
	)

	// Assert
	actual := args.Map{
		"matchingLen":    matching.Length(),
		"notMatchingLen": notMatching.Length(),
	}
	expected := args.Map{
		"matchingLen":    2,
		"notMatchingLen": 1,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns split -- 2+1", actual)
}

func Test_PartitionTypedPayloads_Empty(t *testing.T) {
	// Arrange
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	matching, notMatching := corepayload.PartitionTypedPayloads[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool { return true },
	)

	// Assert
	actual := args.Map{
		"matchingLen":    matching.Length(),
		"notMatchingLen": notMatching.Length(),
	}
	expected := args.Map{
		"matchingLen":    0,
		"notMatchingLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "PartitionTypedPayloads returns empty -- empty source", actual)
}

// ── AnyTypedPayload / AllTypedPayloads ──

func Test_AnyTypedPayload_True(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.AnyTypedPayload[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
			return item.Data().Name == "Bob"
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns true -- Bob exists", actual)
}

func Test_AnyTypedPayload_False(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.AnyTypedPayload[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
			return item.Data().Name == "Nobody"
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AnyTypedPayload returns false -- Nobody missing", actual)
}

func Test_AllTypedPayloads_True(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.AllTypedPayloads[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
			return item.Data().Age > 0
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns true -- all age>0", actual)
}

func Test_AllTypedPayloads_False(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := corepayload.AllTypedPayloads[testUserTyped](col,
		func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
			return item.Data().Age >= 30
		},
	)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllTypedPayloads returns false -- Bob is 25", actual)
}

// ── ConvertTypedPayloads ──

func Test_ConvertTypedPayloads_Valid(t *testing.T) {
	// Arrange
	type simpleUser struct {
		Name string `json:"Name"`
	}
	col := makeTypedCollection()

	// Act
	converted, err := corepayload.ConvertTypedPayloads[testUserTyped, simpleUser](col)

	// Assert
	actual := args.Map{
		"length":  converted.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvertTypedPayloads returns converted -- 3 items", actual)
}

func Test_ConvertTypedPayloads_Empty(t *testing.T) {
	// Arrange
	type other struct{ X int }
	col := corepayload.EmptyTypedPayloadCollection[testUserTyped]()

	// Act
	converted, err := corepayload.ConvertTypedPayloads[testUserTyped, other](col)

	// Assert
	actual := args.Map{
		"isEmpty": converted.IsEmpty(),
		"noError": err == nil,
	}
	expected := args.Map{
		"isEmpty": true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConvertTypedPayloads returns empty -- empty source", actual)
}

// ── TypedPayloadCollection methods ──

func Test_TypedPayloadCollection_ForEachData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()
	var names []string

	// Act
	col.ForEachData(func(index int, data testUserTyped) {
		names = append(names, data.Name)
	})

	// Assert
	actual := args.Map{"length": len(names)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "ForEachData iterates all -- 3 items", actual)
}

func Test_TypedPayloadCollection_ForEachBreak(t *testing.T) {
	// Arrange
	col := makeTypedCollection()
	count := 0

	// Act
	col.ForEachBreak(func(index int, item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
		count++
		return index >= 1
	})

	// Assert
	actual := args.Map{"count": count}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "ForEachBreak stops early -- 2 iterations", actual)
}

func Test_TypedPayloadCollection_FilterByData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.FilterByData(func(u testUserTyped) bool {
		return u.Age >= 30
	})

	// Assert
	actual := args.Map{"length": result.Length()}
	expected := args.Map{"length": 2}
	expected.ShouldBeEqual(t, 0, "FilterByData returns filtered -- 2 seniors", actual)
}

func Test_TypedPayloadCollection_FirstByData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.FirstByData(func(u testUserTyped) bool {
		return u.Name == "Bob"
	})

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"name":   result.Data().Name,
	}
	expected := args.Map{
		"notNil": true,
		"name":   "Bob",
	}
	expected.ShouldBeEqual(t, 0, "FirstByData returns Bob -- found", actual)
}

func Test_TypedPayloadCollection_FirstByName(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.FirstByName("user")

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "FirstByName returns item -- name user", actual)
}

func Test_TypedPayloadCollection_FirstById(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.FirstById("2")

	// Assert
	actual := args.Map{
		"notNil": result != nil,
		"name":   result.Data().Name,
	}
	expected := args.Map{
		"notNil": true,
		"name":   "Bob",
	}
	expected.ShouldBeEqual(t, 0, "FirstById returns Bob -- id 2", actual)
}

func Test_TypedPayloadCollection_CountFunc(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.CountFunc(func(item *corepayload.TypedPayloadWrapper[testUserTyped]) bool {
		return item.Data().Age >= 30
	})

	// Assert
	actual := args.Map{"count": result}
	expected := args.Map{"count": 2}
	expected.ShouldBeEqual(t, 0, "CountFunc returns 2 -- seniors", actual)
}

func Test_TypedPayloadCollection_SkipTake(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	skipped := col.Skip(1)
	taken := col.Take(2)

	// Assert
	actual := args.Map{
		"skippedLen": len(skipped),
		"takenLen":   len(taken),
	}
	expected := args.Map{
		"skippedLen": 2,
		"takenLen":   2,
	}
	expected.ShouldBeEqual(t, 0, "Skip/Take return correct -- slice ops", actual)
}

func Test_TypedPayloadCollection_AllData(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	data := col.AllData()

	// Assert
	actual := args.Map{
		"length": len(data),
		"first":  data[0].Name,
	}
	expected := args.Map{
		"length": 3,
		"first":  "Alice",
	}
	expected.ShouldBeEqual(t, 0, "AllData returns typed slice -- 3 items", actual)
}

func Test_TypedPayloadCollection_AllNames(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	names := col.AllNames()

	// Assert
	actual := args.Map{"length": len(names)}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "AllNames returns names -- 3 items", actual)
}

func Test_TypedPayloadCollection_AllIdentifiers(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	ids := col.AllIdentifiers()

	// Assert
	actual := args.Map{
		"length": len(ids),
		"first": ids[0],
	}
	expected := args.Map{
		"length": 3,
		"first": "1",
	}
	expected.ShouldBeEqual(t, 0, "AllIdentifiers returns ids -- 3 items", actual)
}

func Test_TypedPayloadCollection_ToPayloadsCollection(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	pc := col.ToPayloadsCollection()

	// Assert
	actual := args.Map{"length": pc.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "ToPayloadsCollection returns PC -- 3 items", actual)
}

func Test_TypedPayloadCollection_Clone_FromMapTypedPayloadsType(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	cloned, err := col.Clone()

	// Assert
	actual := args.Map{
		"length":  cloned.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "Clone returns copy -- 3 items", actual)
}

func Test_TypedPayloadCollection_CloneMust(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	cloned := col.CloneMust()

	// Assert
	actual := args.Map{"length": cloned.Length()}
	expected := args.Map{"length": 3}
	expected.ShouldBeEqual(t, 0, "CloneMust returns copy -- 3 items", actual)
}

func Test_TypedPayloadCollection_ConcatNew_FromMapTypedPayloadsType(t *testing.T) {
	// Arrange
	col := makeTypedCollection()
	extra := makeTypedWrapper("user", "4", testUserTyped{Name: "Dave", Age: 40})

	// Act
	result, err := col.ConcatNew(extra)

	// Assert
	actual := args.Map{
		"length":  result.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  4,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "ConcatNew returns new -- 4 items", actual)
}

func Test_TypedPayloadCollection_RemoveAt(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	ok := col.RemoveAt(1)

	// Assert
	actual := args.Map{
		"ok":     ok,
		"length": col.Length(),
	}
	expected := args.Map{
		"ok":     true,
		"length": 2,
	}
	expected.ShouldBeEqual(t, 0, "RemoveAt removes item -- index 1", actual)
}

func Test_TypedPayloadCollection_RemoveAt_OutOfBounds(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	ok := col.RemoveAt(99)

	// Assert
	actual := args.Map{"ok": ok}
	expected := args.Map{"ok": false}
	expected.ShouldBeEqual(t, 0, "RemoveAt returns false -- out of bounds", actual)
}

// ── Paging ──

func Test_TypedPayloadCollection_GetPagesSize(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	pages := col.GetPagesSize(2)

	// Assert
	actual := args.Map{"pages": pages}
	expected := args.Map{"pages": 2}
	expected.ShouldBeEqual(t, 0, "GetPagesSize returns 2 -- 3 items / 2", actual)
}

func Test_TypedPayloadCollection_GetPagedCollection(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	pages := col.GetPagedCollection(2)

	// Assert
	actual := args.Map{
		"pageCount":   len(pages),
		"page1Length": pages[0].Length(),
		"page2Length": pages[1].Length(),
	}
	expected := args.Map{
		"pageCount":   2,
		"page1Length": 2,
		"page2Length": 1,
	}
	expected.ShouldBeEqual(t, 0, "GetPagedCollection returns pages -- 2+1", actual)
}

func Test_TypedPayloadCollection_GetPagedCollectionWithInfo(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	pages := col.GetPagedCollectionWithInfo(2)

	// Assert
	actual := args.Map{
		"pageCount":        len(pages),
		"page1CurrentPage": pages[0].Paging.CurrentPageIndex,
		"page1TotalPages":  pages[0].Paging.TotalPages,
		"page1TotalItems":  pages[0].Paging.TotalItems,
	}
	expected := args.Map{
		"pageCount":        2,
		"page1CurrentPage": 1,
		"page1TotalPages":  2,
		"page1TotalItems":  3,
	}
	expected.ShouldBeEqual(t, 0, "GetPagedCollectionWithInfo returns paging -- 2 pages", actual)
}

// ── IsValid / HasErrors / Errors / MergedError ──

func Test_TypedPayloadCollection_IsValid(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.IsValid()

	// Assert
	actual := args.Map{"isValid": result}
	expected := args.Map{"isValid": true}
	expected.ShouldBeEqual(t, 0, "IsValid returns true -- all parsed", actual)
}

func Test_TypedPayloadCollection_HasErrors(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.HasErrors()

	// Assert
	actual := args.Map{"hasErrors": result}
	expected := args.Map{"hasErrors": false}
	expected.ShouldBeEqual(t, 0, "HasErrors returns false -- no errors", actual)
}

func Test_TypedPayloadCollection_Errors(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	errs := col.Errors()

	// Assert
	actual := args.Map{"length": len(errs)}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "Errors returns empty -- no errors", actual)
}

func Test_TypedPayloadCollection_FirstError(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.FirstError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "FirstError returns nil -- no errors", actual)
}

func Test_TypedPayloadCollection_MergedError(t *testing.T) {
	// Arrange
	col := makeTypedCollection()

	// Act
	result := col.MergedError()

	// Assert
	actual := args.Map{"isNil": result == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "MergedError returns nil -- no errors", actual)
}

// ── NewTypedPayloadCollectionSingle / FromData ──

func Test_NewTypedPayloadCollectionSingle(t *testing.T) {
	// Arrange
	tw := makeTypedWrapper("user", "1", testUserTyped{Name: "Alice"})

	// Act
	col := corepayload.NewTypedPayloadCollectionSingle[testUserTyped](tw)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionSingle returns 1 -- single item", actual)
}

func Test_NewTypedPayloadCollectionSingle_Nil(t *testing.T) {
	// Arrange & Act
	col := corepayload.NewTypedPayloadCollectionSingle[testUserTyped](nil)

	// Assert
	actual := args.Map{"isEmpty": col.IsEmpty()}
	expected := args.Map{"isEmpty": true}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionSingle returns empty -- nil", actual)
}

func Test_NewTypedPayloadCollectionFromData(t *testing.T) {
	// Arrange
	data := []testUserTyped{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	// Act
	col, err := corepayload.NewTypedPayloadCollectionFromData[testUserTyped]("user", data)

	// Assert
	actual := args.Map{
		"length":  col.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromData returns 2 -- from data", actual)
}

func Test_NewTypedPayloadCollectionFromData_Empty(t *testing.T) {
	// Arrange & Act
	col, err := corepayload.NewTypedPayloadCollectionFromData[testUserTyped]("user", []testUserTyped{})

	// Assert
	actual := args.Map{
		"isEmpty": col.IsEmpty(),
		"noError": err == nil,
	}
	expected := args.Map{
		"isEmpty": true,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromData returns empty -- no data", actual)
}

func Test_NewTypedPayloadCollectionFromDataMust(t *testing.T) {
	// Arrange
	data := []testUserTyped{{Name: "Alice"}}

	// Act
	col := corepayload.NewTypedPayloadCollectionFromDataMust[testUserTyped]("user", data)

	// Assert
	actual := args.Map{"length": col.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "NewTypedPayloadCollectionFromDataMust returns 1 -- from data", actual)
}

// ── TypedPayloadWrapperCreator funcs ──

func Test_TypedPayloadWrapperRecord(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperRecord[testUserTyped](
		"user-create", "usr-1", "task", "cat",
		testUserTyped{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"name":    tw.Data().Name,
		"noError": err == nil,
	}
	expected := args.Map{
		"name":    "Alice",
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperRecord returns typed -- valid", actual)
}

func Test_TypedPayloadWrapperRecords(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperRecords[[]testUserTyped](
		"users", "batch-1", "task", "cat",
		[]testUserTyped{{Name: "A"}, {Name: "B"}},
	)

	// Assert
	actual := args.Map{
		"dataLen": len(tw.Data()),
		"noError": err == nil,
	}
	expected := args.Map{
		"dataLen": 2,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperRecords returns slice -- 2 items", actual)
}

func Test_TypedPayloadWrapperNameIdRecord(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperNameIdRecord[testUserTyped](
		"user", "1", testUserTyped{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"name":    tw.Data().Name,
		"noError": err == nil,
	}
	expected := args.Map{
		"name":    "Alice",
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperNameIdRecord returns typed -- valid", actual)
}

func Test_TypedPayloadWrapperNameIdCategory(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperNameIdCategory[testUserTyped](
		"user", "1", "cat", testUserTyped{Name: "Alice"},
	)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"catName": tw.CategoryName(),
	}
	expected := args.Map{
		"noError": true,
		"catName": "cat",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperNameIdCategory returns typed -- valid", actual)
}

func Test_TypedPayloadWrapperAll(t *testing.T) {
	// Arrange & Act
	tw, err := corepayload.TypedPayloadWrapperAll[testUserTyped](
		"name", "id", "task", "entity", "cat", false,
		testUserTyped{Name: "Alice"}, nil,
	)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    tw.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperAll returns typed -- valid", actual)
}

func Test_TypedPayloadWrapperDeserialize(t *testing.T) {
	// Arrange
	tw := makeTypedWrapper("user", "1", testUserTyped{Name: "Alice"})
	jsonBytes := tw.SerializeMust()

	// Act
	result, err := corepayload.TypedPayloadWrapperDeserialize[testUserTyped](jsonBytes)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    result.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserialize returns typed -- valid", actual)
}

func Test_TypedPayloadWrapperDeserializeUsingJsonResult(t *testing.T) {
	// Arrange
	tw := makeTypedWrapper("user", "1", testUserTyped{Name: "Alice"})
	jsonResult := tw.JsonPtr()

	// Act
	result, err := corepayload.TypedPayloadWrapperDeserializeUsingJsonResult[testUserTyped](jsonResult)

	// Assert
	actual := args.Map{
		"noError": err == nil,
		"name":    result.Data().Name,
	}
	expected := args.Map{
		"noError": true,
		"name":    "Alice",
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadWrapperDeserializeUsingJsonResult returns typed -- valid", actual)
}

func Test_TypedPayloadCollectionDeserialize(t *testing.T) {
	// Arrange
	col := makeTypedCollection()
	pc := col.ToPayloadsCollection()
	jsonBytes, _ := corejson.Serialize.Raw(pc.Items)

	// Act
	result, err := corepayload.TypedPayloadCollectionDeserialize[testUserTyped](jsonBytes)

	// Assert
	actual := args.Map{
		"length":  result.Length(),
		"noError": err == nil,
	}
	expected := args.Map{
		"length":  3,
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "TypedPayloadCollectionDeserialize returns 3 -- from bytes", actual)
}

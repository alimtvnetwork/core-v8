# Phase 2 — S## / Seg# File Merge Plan

## Status: 🟢 In progress — Batch 2.1 ✅ green (`run.ps1 -tc` 10/11, corestr 98.5%); Batch 2.2 ✅ code complete, awaiting `run.ps1 -tc`

### Batch 2.2 (executed 2026-04-18)
Source files removed (3) → target files created (11):
- `S20_001_S20_NonChainedNodes_HashmapDiff_test.go` (52 funcs) →
  - `NonChainedLinkedListNodes_AllMethods_test.go` (14)
  - `NonChainedLinkedCollectionNodes_AllMethods_test.go` (14)
  - `HashmapDiff_AllMethods_test.go` (24)
- `S21_001_S21_CloneSlice_Empty_Reflect_test.go` (31 funcs) →
  - `CloneSlice_Behaviour_test.go` (6)
  - `Empty_Constructors_test.go` (17)
  - `AnyToString_Behaviour_test.go` (8)
- `S22_001_S22_FromSplit_Creators_test.go` (103 funcs) →
  - `LeftRightFromSplit_Factories_test.go` (6)
  - `LeftMiddleRightFromSplit_Factories_test.go` (6)
  - `NewCollection_Factories_test.go` (13)
  - `NewHashset_Factories_test.go` (16)
  - `NewHashmap_Factories_test.go` (62)

Symbol pattern: `Test_<Type>_<Behaviour>_FromS20|S21|S22`. Zero duplicates verified via `grep | sort | uniq -d`.

Parent plan: `spec/01-app/28-corestrtests-cleanup-plan.md`
Scope: `tests/integratedtests/corestrtests/` only.

---

## 1. Inventory (27 files)

| # | Source file | Primary subject(s) | Test count |
|---|-------------|-------------------|------------|
| 1 | `S01_InvalidValueStatus_S01_ValueStatus_TWLN_test.go` | `ValueStatus`, `TextWithLineNumber` | ~14 |
| 2 | `S01_NewValidValue_S01_ValidValue_test.go` | `ValidValue` constructors + Bool/Int/Byte/Float converters | ~30 |
| 3 | `S08_01_Collection_S08_test.go` | `Collection` JSON / RemoveAt / Length / Capacity / IsEquals | ~30 |
| 4 | `S09_01_Hashmap_S09_test.go` | `Hashmap` IsEmpty / AddOrUpdate variants / Set | ~25 |
| 5 | `S10_01_Hashset_S10a_test.go` | `Hashset` IsEmpty / AddCapacities / Resize / Concat | ~25 |
| 6 | `S10_88_Hashset_S10b_test.go` | `Hashset` SortedList / Filter / Lines / SimpleSlice / GetFiltered | ~20 |
| 7 | `S11_01_SimpleSlice_S11a_test.go` | `SimpleSlice` Add / AddIf / Adds / Append / AppendFmt | ~25 |
| 8 | `S11_66_SimpleSlice_S11b_test.go` | `SimpleSlice` (later methods) | ~20 |
| 9 | `S12_01_LinkedList_S12_test.go` | `LinkedList` core methods | ~25 |
| 10 | `S13_01_LinkedCollections_S13_test.go` | `LinkedCollection(s)` core methods | ~25 |
| 11 | `S20_001_S20_NonChainedNodes_HashmapDiff_test.go` | `NonChainedLinkedListNodes`, `NonChainedLinkedCollectionNodes`, `Hashmap.Diff` | ~30 |
| 12 | `S21_001_S21_CloneSlice_Empty_Reflect_test.go` | `CloneSlice`, `CloneSliceIf`, type `.Empty()` constructors via reflect | ~30 |
| 13 | `S22_001_S22_FromSplit_Creators_test.go` | `LeftRightFromSplit*`, `LeftMiddleRightFromSplit*`, `NewCollection` factories | ~30 |
| 14 | `Seg1_AllIndividualStringsOfStringsLength_Seg1_Utilities_test.go` | `AllIndividualStringsOfStringsLength` + misc utils | ~10 |
| 15 | `Seg2_Collection_Seg2_CollectionMid_test.go` | `Collection` middle batch | ~20 |
| 16 | `Seg3_Collection_Seg3_CollectionEnd_test.go` | `Collection` end batch | ~20 |
| 17 | `Seg4_SS_Seg4_SimpleSlice_test.go` | `SimpleSlice` extra batch | ~20 |
| 18 | `Seg5_HM_Seg5_Hashmap_test.go` | `Hashmap` extra batch | ~20 |
| 19 | `Seg5_HS_Seg5_Hashset_test.go` | `Hashset` extra batch | ~20 |
| 20 | `Seg6_CCM_Seg6_CharCollectionMap_test.go` | `CharCollectionMap` | ~20 |
| 21 | `Seg6_CHM_Seg6_CharHashsetMap_test.go` | `CharHashsetMap` | ~20 |
| 22 | `Seg6_COC_Seg6_CollOfColl_KV_test.go` | `CollectionsOfCollection` + `KeyValues` mixed | ~20 |
| 23 | `Seg7_LC_Seg7_LinkedCollections_test.go` | `LinkedCollections` | ~20 |
| 24 | `Seg7_LL_Seg7_LinkedList_test.go` | `LinkedList` extra batch | ~20 |
| 25 | `Seg7_LR_Seg7_LeftRight_test.go` | `LeftRight` family | ~25 |
| 26 | `Seg8_SSO_Seg8_SimpleStringOnce_test.go` | `SimpleStringOnce` | ~20 |
| 27 | `Seg8_VV_Seg8_ValidValue_VV_VS_TWLN_test.go` | `ValidValue` extra + `ValueStatus` + `TextWithLineNumber` mixed | ~30 |

## 2. Merge Strategy

### Guiding principles
1. **Group by primary type, not by batch ID.** Every file ends up under `<Type>_…_test.go`.
2. **Prefer new files over giant merges.** A new `Hashmap_AddOrUpdate_test.go` is clearer than appending 25 tests to an existing `Hashmap_BasicOps_test.go`.
3. **Preserve every test function.** Renames only — no deletions.
4. **Strip `_S##` / `_Seg#` / `_<initialism>` tokens from test symbols** when they appear (e.g. `Test_Seg7_LR_NewLeftRight` → `Test_LeftRight_NewLeftRight_StoresBothSides`).
5. **One file may be split into multiple new files** when it covers two distinct types (files #1, #11, #12, #13, #22, #27).
6. **Do not collide with existing names** — check `tests/integratedtests/corestrtests/` directory before each rename.

### Per-file action

| # | Source file | Action | Target file(s) |
|---|-------------|--------|----------------|
| 1 | `S01_InvalidValueStatus_S01_ValueStatus_TWLN_test.go` | **SPLIT** | `ValueStatus_InvalidStatus_test.go` + `TextWithLineNumber_LineNumberAndEmptiness_test.go` |
| 2 | `S01_NewValidValue_S01_ValidValue_test.go` | **SPLIT** | `ValidValue_Constructors_test.go` + `ValidValue_PrimitiveConverters_test.go` |
| 3 | `S08_01_Collection_S08_test.go` | rename | `Collection_JsonAndIndexing_test.go` |
| 4 | `S09_01_Hashmap_S09_test.go` | rename | `Hashmap_AddOrUpdate_test.go` |
| 5 | `S10_01_Hashset_S10a_test.go` | rename | `Hashset_AddCapacitiesAndResize_test.go` |
| 6 | `S10_88_Hashset_S10b_test.go` | rename | `Hashset_FilterAndSortedList_test.go` |
| 7 | `S11_01_SimpleSlice_S11a_test.go` | rename | `SimpleSlice_AddAndAppend_test.go` |
| 8 | `S11_66_SimpleSlice_S11b_test.go` | rename | `SimpleSlice_LaterMethods_test.go` *(name to be tightened after reading body)* |
| 9 | `S12_01_LinkedList_S12_test.go` | rename | `LinkedList_CoreMethods_test.go` *(or split if two clear topics emerge)* |
| 10 | `S13_01_LinkedCollections_S13_test.go` | rename | `LinkedCollections_CoreMethods_test.go` |
| 11 | `S20_001_S20_NonChainedNodes_HashmapDiff_test.go` | **SPLIT** | `NonChainedLinkedListNodes_LifecycleAndChaining_test.go` + `NonChainedLinkedCollectionNodes_Lifecycle_test.go` + `Hashmap_Diff_test.go` |
| 12 | `S21_001_S21_CloneSlice_Empty_Reflect_test.go` | **SPLIT** | `CloneSlice_AllVariants_test.go` + `Constructors_EmptyFactories_test.go` |
| 13 | `S22_001_S22_FromSplit_Creators_test.go` | **SPLIT** | `LeftRight_FromSplitFactories_test.go` + `LeftMiddleRight_FromSplitFactories_test.go` + `Collection_NewFactories_test.go` |
| 14 | `Seg1_AllIndividualStringsOfStringsLength_Seg1_Utilities_test.go` | rename | `AllIndividualStringsOfStringsLength_AllInputs_test.go` |
| 15 | `Seg2_Collection_Seg2_CollectionMid_test.go` | rename | `Collection_MidBatchMethods_test.go` *(tighten name after reading)* |
| 16 | `Seg3_Collection_Seg3_CollectionEnd_test.go` | rename | `Collection_EndBatchMethods_test.go` *(tighten name after reading)* |
| 17 | `Seg4_SS_Seg4_SimpleSlice_test.go` | rename | `SimpleSlice_ExtraMethods_test.go` *(tighten)* |
| 18 | `Seg5_HM_Seg5_Hashmap_test.go` | rename | `Hashmap_ExtraMethods_test.go` *(tighten)* |
| 19 | `Seg5_HS_Seg5_Hashset_test.go` | rename | `Hashset_ExtraMethods_test.go` *(tighten)* |
| 20 | `Seg6_CCM_Seg6_CharCollectionMap_test.go` | rename | `CharCollectionMap_CoreMethods_test.go` |
| 21 | `Seg6_CHM_Seg6_CharHashsetMap_test.go` | rename | `CharHashsetMap_CoreMethods_test.go` |
| 22 | `Seg6_COC_Seg6_CollOfColl_KV_test.go` | **SPLIT** | `CollectionsOfCollection_CoreMethods_test.go` + `KeyValues_CoreMethods_test.go` |
| 23 | `Seg7_LC_Seg7_LinkedCollections_test.go` | rename | `LinkedCollections_ExtraMethods_test.go` |
| 24 | `Seg7_LL_Seg7_LinkedList_test.go` | rename | `LinkedList_ExtraMethods_test.go` |
| 25 | `Seg7_LR_Seg7_LeftRight_test.go` | rename | `LeftRight_ConstructorsAndUsingSlice_test.go` |
| 26 | `Seg8_SSO_Seg8_SimpleStringOnce_test.go` | rename | `SimpleStringOnce_CoreMethods_test.go` |
| 27 | `Seg8_VV_Seg8_ValidValue_VV_VS_TWLN_test.go` | **SPLIT** | `ValidValue_ExtraMethods_test.go` + `ValueStatus_ExtraMethods_test.go` + `TextWithLineNumber_ExtraMethods_test.go` |

**Rename total:** 19 straight renames + 8 splits = 27 source files → ~36 target files.

### Symbol cleanup inside each merged file

Drop noise tokens from `Test_…` symbols and assertion titles:
- `Test_Seg7_LR_NewLeftRight` → `Test_LeftRight_New_StoresBothSides`
- `Test_Seg6_COC_IsEmpty` → `Test_CollectionsOfCollection_IsEmpty_NewIsTrue`
- `Test_001_NewNonChainedLinkedListNodes_creates_with_capacity` → `Test_NonChainedLinkedListNodes_NewCap_CreatesWithCapacity`
- `Test_NewValidValue_S01NewvalidvalueS01Validvalue` → `Test_ValidValue_New_IsValidWithValue`
- Auto-numbered `Test_NN_<Type>_<Method>` → `Test_<Type>_<Method>_<Behaviour>`

Where the bare descriptive name would collide with a sibling test file, append a behaviour suffix (same rule used in Phase 1).

## 3. Execution Plan (sequential batches)

To keep `run.ps1 -tc` green at every checkpoint, execute in **5 batches**, validating after each.

| Batch | Files | Action |
|-------|-------|--------|
| **2.1** | 1, 2, 27 | `ValidValue` + `ValueStatus` + `TextWithLineNumber` splits — three smallest related types |
| **2.2** | 11, 12, 13 | `S20`/`S21`/`S22` multi-topic splits (NonChainedNodes, CloneSlice/Empty, FromSplit) |
| **2.3** | 22, 25, 26 | CollectionsOfCollection / KeyValues / LeftRight / SimpleStringOnce |
| **2.4** | 4, 5, 6, 18, 19, 20, 21 | Hashmap + Hashset + CharCollectionMap + CharHashsetMap renames |
| **2.5** | 3, 7, 8, 9, 10, 14, 15, 16, 17, 23, 24 | Collection + SimpleSlice + LinkedList + LinkedCollections + utilities renames |

**Validation gate (mandatory after every batch):**
1. User runs `.\run.ps1 -tc`
2. Status must be `✓ READY TO COMMIT`
3. Coverage for `coredata/corestr` must stay ≥ 98.5%
4. Update batch row in this file with ✅ + date
5. Only then proceed to next batch

## 4. Open Decisions Before Batch 2.1 Starts

1. **Batch granularity** — OK with 5 batches as above, or prefer smaller (10 batches of ~3 files each) for tighter rollback? Default: **5 batches**.
2. **"_ExtraMethods" / "_LaterMethods" naming** — for files #8, #15, #16, #17, #18, #19, #23, #24, #27 the temporary name is a placeholder. Should I:
   - (a) Open each file, classify the methods, and pick a precise descriptor (slow but final), **or**
   - (b) Use `_ExtraMethods` now, refine in Phase 4 ("eliminate vague suffixes")?
   Default: **(b)** — keeps Phase 2 scope narrow.
3. **Splits with cross-references** — when a split moves a test that uses a helper variable defined in the original file, I will move the helper to `testhelpers.go` (Phase 6 territory) **or** copy it into both new files with a `_helpers.go` suffix. Default: **copy** in Phase 2, consolidate in Phase 6.

## 5. Phase 2 Status Tracker

| Batch | Status | Started | Finished | Coverage delta |
|-------|--------|---------|----------|----------------|
| 2.1 — ValidValue/ValueStatus/TWLN splits | ✅ code | 2026-04-18 | 2026-04-18 | pending `run.ps1 -tc` |
| 2.2 — NonChainedNodes/CloneSlice/FromSplit splits | ⬜ | — | — | — |
| 2.3 — CollectionsOfCollection/KeyValues/LeftRight/SSO | ⬜ | — | — | — |
| 2.4 — Hashmap/Hashset/CharMaps renames | ⬜ | — | — | — |
| 2.5 — Collection/SimpleSlice/LinkedList/utilities renames | ⬜ | — | — | — |

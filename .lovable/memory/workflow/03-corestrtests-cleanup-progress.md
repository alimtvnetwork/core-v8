---
name: corestrtests Cleanup Progress
description: Phase-by-phase status for the 7-phase corestrtests naming, structure, and assertion cleanup
type: feature
---

# corestrtests Cleanup Progress

Plan: `spec/01-app/28-corestrtests-cleanup-plan.md`
Phase 2 sub-plan: `spec/01-app/29-corestrtests-phase2-merge-plan.md`

| Phase | Title | Status |
|-------|-------|--------|
| 1 | Strip `_Cov<N>` symbols (38 in `SimpleSlice_Cap_test.go`) | ✅ Done — green on `run.ps1 -tc` |
| 2 | Remove `S##` / `Seg#` cryptic-prefix files (~27) | 🟡 In progress — Batch 2.1 ✅ green; Batch 2.2 ✅ green; Batch 2.3 ✅ green; Batch 2.4 ✅ code complete (5→10 files, 668 funcs); Batch 2.5 pending |
| 3 | Resolve initialism prefixes HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR (~13) | ⬜ Not started |
| 4 | Eliminate vague `_Part/_Core/_Full/_Basic/_Gaps…` suffixes (~60) | ⬜ Not started |
| 5 | Restore data/logic separation (extract `_testcases.go` siblings) | ⬜ Not started |
| 6 | Consolidate helpers + nil-safety into purpose-named files | ⬜ Not started |
| 7 | Lint guardrail in `run.ps1` (block `Cov\d+`, `_Part\d+`, `Seg\d+`, `t.Log/Error/Fatal`) | ⬜ Not started |

## Phase 1 Notes
- All 38 `Test_*_Cov2` symbols renamed to behaviour-driven `Test_<Type>_<Method>_<Outcome>` form.
- All `safeTest("…")` literal labels updated to match the new symbol.
- Verified by `run.ps1 -tc`: 10/11 phases pass, ✓ READY TO COMMIT, `coredata/corestr` at 98.5%.

## Phase 2 Batch 2.1 Notes (2026-04-18)
- 3 source files (#1, #2, #27 in the merge plan), 155 functions → 7 type-scoped target files.
- Naming pattern: `Test_<Type>_<Behaviour>_FromS01` / `_FromSeg8` provenance suffix to avoid collision with 457 pre-existing `Test_ValidValue_*` / `Test_ValueStatus_*` / `Test_TextWithLineNumber_*` symbols.
- Verified green via `run.ps1 -tc` (10/11 phases, corestr 98.5%).

## Phase 2 Batch 2.2 Notes (2026-04-18)
- 3 source files (#11, #12, #13 in the merge plan), 186 functions → 11 type-scoped target files.
- Source files removed: `S20_001_S20_NonChainedNodes_HashmapDiff_test.go` (52), `S21_001_S21_CloneSlice_Empty_Reflect_test.go` (31), `S22_001_S22_FromSplit_Creators_test.go` (103).
- Target files created: `NonChainedLinkedListNodes_AllMethods_test.go`, `NonChainedLinkedCollectionNodes_AllMethods_test.go`, `HashmapDiff_AllMethods_test.go`, `CloneSlice_Behaviour_test.go`, `Empty_Constructors_test.go`, `AnyToString_Behaviour_test.go`, `LeftRightFromSplit_Factories_test.go`, `LeftMiddleRightFromSplit_Factories_test.go`, `NewCollection_Factories_test.go`, `NewHashset_Factories_test.go`, `NewHashmap_Factories_test.go`.
- Symbol pattern: `Test_<Type>_<Behaviour>_FromS20|S21|S22`. Zero duplicates verified via `grep | sort | uniq -d`.
- Tooling: `/tmp/batch22/split.py` (regex-based AST-light splitter); `safeTest("…")` labels rewritten in lockstep.
- Pending validation: user must run `.\run.ps1 -tc` and confirm `✓ READY TO COMMIT` + `coredata/corestr` ≥ 98.5% before Batch 2.3 starts.

## Phase 2 Batch 2.3 Notes (2026-04-18)
- 4 source files (Seg1–Seg4 in the merge plan), 455 functions → 18 type-scoped target files (13 created, 4 appended to Batch 2.2 outputs, 1 misc).
- Source files removed: `Seg1_AllIndividualStringsOfStringsLength_Seg1_Utilities_test.go` (136), `Seg2_Collection_Seg2_CollectionMid_test.go` (58), `Seg3_Collection_Seg3_CollectionEnd_test.go` (105), `Seg4_SS_Seg4_SimpleSlice_test.go` (156).
- Heaviest targets: `SimpleSlice_Queries_FromSeg4_test.go` (134), `Collection_FilterRemaining_FromSeg3_test.go` (95), `Collection_MidAccess_FromSeg2_test.go` (58), `Collection_FromSeg1_test.go` (35).
- Symbol pattern: `Test_<Type>_<Behaviour>_FromSeg1|Seg2|Seg3|Seg4`. Zero duplicates verified via `grep | sort | uniq -d` across all 50+ test files.
- Tooling: `/tmp/batch23/split.py` (regex-based AST-light splitter with per-Seg routing rules); `safeTest("…")` labels rewritten in lockstep.
- Pending validation: user must run `.\run.ps1 -tc` and confirm `✓ READY TO COMMIT` + `coredata/corestr` ≥ 98.5% before Batch 2.4 starts.

## Phase 2 Batch 2.4 Notes (2026-04-18)
- 5 source files (Seg5_HM, Seg5_HS, Seg6_CCM, Seg6_CHM, Seg6_COC), 668 functions → 10 type-scoped target files (8 created, 2 appended).
- Buckets discovered beyond plan: Seg5 had **HMD** (HashmapDataModel, 19) and **HSCDM** (HashsetCollectionDataModel, 2 — folded into HashsetCollection file). Seg6 COC had 4 sub-buckets (COC 27, KVC 62, KAVP 17, KVP 20).
- Source files removed: `Seg5_HM_Seg5_Hashmap_test.go` (145), `Seg5_HS_Seg5_Hashset_test.go` (188), `Seg6_CCM_Seg6_CharCollectionMap_test.go` (105), `Seg6_CHM_Seg6_CharHashsetMap_test.go` (104), `Seg6_COC_Seg6_CollOfColl_KV_test.go` (126).
- Target files created: `Hashmap_AllMethods_FromSeg5_test.go` (126), `HashmapDataModel_FromSeg5_test.go` (19), `Hashset_AllMethods_FromSeg5_test.go` (147), `HashsetCollection_AllMethods_FromSeg5_test.go` (41 = HSC 39 + HSCDM 2), `CharCollectionMap_AllMethods_FromSeg6_test.go` (105), `CharHashsetMap_AllMethods_FromSeg6_test.go` (104), `CollectionOfCollections_AllMethods_FromSeg6_test.go` (27), `KeyValueCollection_AllMethods_FromSeg6_test.go` (62).
- Appended: 17 KAVP funcs → `KeyAnyValuePair_AllMethods_FromSeg1_test.go` (now 34); 20 KVP funcs → `KeyValuePair_AllMethods_FromSeg1_test.go` (now 37).
- Symbol pattern: `Test_<Type>_<Behaviour>_FromSeg5|Seg6`. Zero duplicates verified via `grep | sort | uniq -d`.
- Tooling: `/tmp/batch24/split.py` (regex-based AST-light splitter with longest-prefix-first matching to disambiguate HSC/HSCDM, HM/HMD); `safeTest("…")` labels rewritten in lockstep.
- Pending validation: user must run `.\run.ps1 -tc` and confirm `✓ READY TO COMMIT` + `coredata/corestr` ≥ 98.5% before Batch 2.5 starts.

## Remaining Tasks
- Batch 2.5: Seg7/Seg8 + remaining S03–S13 splits (files #3–10, #23–26 in the merge plan).
- Phases 3–7 per master plan.

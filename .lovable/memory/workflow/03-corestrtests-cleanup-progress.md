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
| 2 | Remove `S##` / `Seg#` cryptic-prefix files (~27) | ✅ Done — Batches 2.1–2.5 green |
| 3 | Resolve initialism prefixes HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR (~13) | 🟡 Code complete (10 files, 1185 funcs, 340 disambiguated symbols) — pending `-tc` |
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

## Phase 2 Batch 2.5 Notes (2026-04-18)
- 13 source files (Seg7 LL/LC/LR, Seg8 SSO, S08, S08b, S09, S10a, S10b, S11a, S11b, S12, S13), 1,295 functions → 13 type-scoped target files (1:1 mapping; no per-method-family sub-split this round per user direction).
- Source files removed and replaced by: `LinkedList_AllMethods_FromSeg7_test.go` (88), `LinkedCollections_AllMethods_FromSeg7_test.go` (126), `LeftRight_AllMethods_FromSeg7_test.go` (50), `SimpleStringOnce_AllMethods_FromSeg8_test.go` (99), `Collection_AllMethods_FromS08_test.go` (99), `Collection_AllMethods_FromS08b_test.go` (133), `Hashmap_AllMethods_FromS09_test.go` (154), `Hashset_AllMethods_FromS10a_test.go` (87), `Hashset_AllMethods_FromS10b_test.go` (89), `SimpleSlice_AllMethods_FromS11a_test.go` (65), `SimpleSlice_AllMethods_FromS11b_test.go` (91), `LinkedList_AllMethods_FromS12_test.go` (102), `LinkedCollections_AllMethods_FromS13_test.go` (112).
- Symbol pattern: `Test_<Type>_<Behaviour>_FromSeg7|Seg8|S08|S08b|S09|S10a|S10b|S11a|S11b|S12|S13`. Zero duplicates verified via `grep | sort | uniq -d` across 206 test files / 14,104 Test funcs.
- Tooling: `/tmp/batch25_split.py` — single regex-based pass: rename Test_ symbols, rewrite `safeTest("…")` labels, prune unused candidate imports per-file (lesson learned from Batch 2.4).
- Pending validation: user runs `.\run.ps1 -tc`; on green, Phase 2 closes and Phase 3 (initialism prefixes) begins.

## Remaining Tasks
- Phase 3: Resolve any leftover initialism-prefixed files (HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR) introduced as targets — most already type-scoped, audit for stragglers.
- Phases 4–7 per master plan (vague suffixes, data/logic separation, helper consolidation, lint guardrail).

## Batch 2.4 Post-Verification Fix (2026-04-18)

**Issue**: 15 of the new Batch 2.3/2.4 split files failed compile-check during `.\run.ps1 -tc` split-recovery due to **unused imports** (`errors`, `regexp`, `corejson`, `sync`) inherited from the original Seg source files.

**Root cause**: When splitting Seg1/Seg3/Seg4/Seg5 into smaller per-type files, the full import block was duplicated to every child file even when the new file no longer referenced those packages. The full `corestrtests` package compiled because some other file used them, but the per-file split-recovery surfaced the unused-import errors.

**Fix**: Pruned each file's import block to keep only packages actually referenced in its body (preserved `errors` in `Collection_FromSeg1`, `corejson` in `KeyAnyValuePair_AllMethods_FromSeg1`, `regexp` in `LeftRight_Behaviour_FromSeg1`).

**Files cleaned** (15): AllIndividualsLengthOfSimpleSlices_Behaviour, AllIndividualStringsOfStringsLength_Behaviour, Collection_FilterLock_FromSeg3, Collection_FilterPtr_FromSeg3, Collection_FromSeg1, HashmapDataModel_FromSeg5, HashsetCollection_AllMethods_FromSeg5, KeyAnyValuePair_AllMethods_FromSeg1, KeyValuePair_AllMethods_FromSeg1, LeftMiddleRight_Behaviour_FromSeg1, LeftRight_Behaviour_FromSeg1, SimpleSlice_AddAppend_FromSeg4, SimpleSlice_Mutators_FromSeg4, TextWithLineNumber_FromSeg1, ValueStatus_FromSeg1.

**Lesson**: After every split that produces sub-50-func files, run an automated unused-import prune step before claiming green — the per-file recovery check will catch what the package-level build hides.

## Phase 3 Notes (2026-04-18)
- Audit found 10 initialism-prefixed files (1,185 funcs total): `KVP_Basic`, `LC_Add`, `LC_Head_LinkedCollSeg1`, `LL_Add`, `SSO_Core`, `SSO_Extended`, `SSO_Json_RemainingGaps`, `SSO_SplitLeftRight`, `SSO_Value`, `SS_WrapDoubleQuote_SimpleSlice_S03`.
- HM/HS/CHM/CCM/COC/LR/VV initialism prefixes: zero matches — already on full names.
- Renamed to: `KeyValuePair_Basic`, `LinkedCollections_Add`, `LinkedCollections_Head_FromSeg1`, `LinkedList_Add`, `SimpleStringOnce_{Core,Extended,Json_RemainingGaps,SplitLeftRight,Value}`, `SimpleSlice_WrapDoubleQuote_FromS03`.
- Symbol-prefix rewrite (`Test_KVP_*` → `Test_KeyValuePair_*` etc.) caused 340 collisions with pre-existing symbols. Resolved by appending per-file disambiguator tags: `_Basic`, `_AddV2`, `_HeadSeg1`, `_Core`, `_Extended`, `_WrapDQ`, etc.
- Verified zero duplicates package-wide: 14,104 unique Test funcs across 237 files.
- Tooling: `/tmp/phase3.py` (rename + prefix rewrite + safeTest label sync), `/tmp/phase3_fix.py` (collision disambiguator).
- Lesson learned: mass prefix expansion requires upfront collision detection AND a disambiguator strategy — `_<short-tag>` per source file is sufficient and preserves provenance.
- Pending validation: user must run `.\run.ps1 -tc` and confirm `✓ READY TO COMMIT`.

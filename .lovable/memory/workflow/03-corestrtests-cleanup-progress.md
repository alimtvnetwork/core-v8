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
| 1 | Strip `_Cov<N>` symbols (38 in `SimpleSlice_Cap_test.go`) | ✅ Code done — awaiting `run.ps1 -tc` |
| 2 | Remove `S##` / `Seg#` cryptic-prefix files (~27) | 🟡 In progress — Batch 2.1 (3 of 27 files) ✅ code, batches 2.2–2.5 pending |
| 3 | Resolve initialism prefixes HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR (~13) | ⬜ Not started |
| 4 | Eliminate vague `_Part/_Core/_Full/_Basic/_Gaps…` suffixes (~60) | ⬜ Not started |
| 5 | Restore data/logic separation (extract `_testcases.go` siblings) | ⬜ Not started |
| 6 | Consolidate helpers + nil-safety into purpose-named files | ⬜ Not started |
| 7 | Lint guardrail in `run.ps1` (block `Cov\d+`, `_Part\d+`, `Seg\d+`, `t.Log/Error/Fatal`) | ⬜ Not started |

## Phase 1 Notes
- All 38 `Test_*_Cov2` symbols renamed to `Test_<Type>_<Behaviour>` form (e.g. `Test_SimpleSlice_NewCap_InitialStateIsEmpty`).
- All `safeTest("…")` literal labels updated to match the new symbol.
- All `ShouldBeEqual` assertion titles rewritten in behaviour-driven `"<Method> <verb> <object> -- <context>"` form.
- File name `SimpleSlice_Cap_test.go` is now misleading (covers ~25 types); will be split/merged in Phase 2 or 3.

## Phase 2 Batch 2.1 Notes (2026-04-18)
- Source: 3 files (#1, #2, #27 in the merge plan), 155 test functions total.
- Target: 7 new files with type-scoped names — `ValueStatus_InvalidStatus_test.go`, `TextWithLineNumber_LineNumberAndEmptiness_test.go`, `ValidValue_Constructors_test.go`, `ValidValue_PrimitiveConverters_test.go`, `ValidValue_ExtraMethods_test.go`, `ValueStatus_ExtraMethods_test.go`, `TextWithLineNumber_ExtraMethods_test.go`.
- Symbol renaming: behaviour-driven names with `_FromS01` / `_FromSeg8` provenance suffix to avoid collisions with the 457 pre-existing `Test_ValidValue_*` / `Test_ValueStatus_*` / `Test_TextWithLineNumber_*` symbols in sibling files. Suffix will be re-evaluated in Phase 6 once helper consolidation deduplicates semantically identical tests.
- Tooling: split performed by `/tmp/batch21/split.py` (regex-based AST-light splitter). `safeTest("…")` labels rewritten in lockstep with function declarations.
- Verification: `grep | sort | uniq -d` on all `^func Test_` declarations across `corestrtests/` returned zero duplicates.
- Defaults adopted: 5 batches, `_ExtraMethods` placeholder for now (Phase 4 will refine), helpers copied per-file in Phase 2 (consolidation deferred to Phase 6).
- Pending validation: user must run `.\run.ps1 -tc` and confirm `✓ READY TO COMMIT` + `coredata/corestr` ≥ 98.5% before Batch 2.2 starts.

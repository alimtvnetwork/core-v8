---
name: corestrtests Cleanup Progress
description: Phase-by-phase status for the 7-phase corestrtests naming, structure, and assertion cleanup
type: feature
---

# corestrtests Cleanup Progress

Plan: `spec/01-app/28-corestrtests-cleanup-plan.md`

| Phase | Title | Status |
|-------|-------|--------|
| 1 | Strip `_Cov<N>` symbols (38 in `SimpleSlice_Cap_test.go`) | ✅ Code done — awaiting `run.ps1 -tc` |
| 2 | Remove `S##` / `Seg#` cryptic-prefix files (~27) | ⬜ Not started |
| 3 | Resolve initialism prefixes HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR (~13) | ⬜ Not started |
| 4 | Eliminate vague `_Part/_Core/_Full/_Basic/_Gaps…` suffixes (~60) | ⬜ Not started |
| 5 | Restore data/logic separation (extract `_testcases.go` siblings) | ⬜ Not started |
| 6 | Consolidate helpers + nil-safety into purpose-named files | ⬜ Not started |
| 7 | Lint guardrail in `run.ps1` (block `Cov\d+`, `_Part\d+`, `Seg\d+`, `t.Log/Error/Fatal`) | ⬜ Not started |

## Phase 1 Notes
- All 38 `Test_*_Cov2` symbols renamed to `Test_<Type>_<Behaviour>` form (e.g. `Test_SimpleSlice_NewCap_InitialStateIsEmpty`).
- All `safeTest("…")` literal labels updated to match the new symbol.
- All `ShouldBeEqual` assertion titles rewritten in behaviour-driven `"<Method> <verb> <object> -- <context>"` form.
- No external file references — `_Cov2` only existed inside this single file. Zero cross-file rename impact.
- Renames chosen to avoid collision with sibling test packages: bare names like `Test_SimpleSlice_Add` already exist elsewhere, so each Phase 1 symbol carries a behaviour suffix.
- File name `SimpleSlice_Cap_test.go` is now misleading (covers ~25 types); will be split/merged in Phase 2 or 3.

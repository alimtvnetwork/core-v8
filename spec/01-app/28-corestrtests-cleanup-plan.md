# corestrtests — Naming, Structure & Assertion Cleanup Plan

## Status: 🔄 In Progress — Phase 1 ✅ done, awaiting `run.ps1 -tc` validation before Phase 2
## Scope: `tests/integratedtests/corestrtests/` only (211 `.go` files: 180 `_test.go`, 29 `_testcases.go`, 2 helpers)

---

## 1. Deep-Dive Findings

### 1.1 Naming Convention Issues

#### A. Cryptic file-name prefixes (40 files)
Files whose names start with an opaque short code instead of the function/feature under test. The reader cannot tell what is being verified without opening the file.

| Prefix | Meaning (inferred) | Example files | Count |
|--------|-------------------|---------------|-------|
| `S01_…S22_` | "Section" / split-recovery batch IDs from the Phase 1 cleanup | `S01_NewValidValue_S01_ValidValue_test.go`, `S20_001_S20_NonChainedNodes_HashmapDiff_test.go` | 13 |
| `Seg1…Seg8` | "Segment" — same as above, different generation pass | `Seg2_Collection_Seg2_CollectionMid_test.go`, `Seg6_CCM_Seg6_CharCollectionMap_test.go` | 14 |
| `HM01_`, `HS01_`, `SS01_`, `LL_`, `LC_`, `VV_`, `KVP_`, `SSO_`, `CCM_`, `CHM_`, `COC_`, `LR_` | Type initialisms | `HM01_IsEmpty_Hashmap_I8_test.go`, `LL_Add_test.go`, `KVP_Basic_test.go` | 13 |

Several names also **repeat the same token twice** (`S01_…_S01_…`, `Seg6_CCM_Seg6_CharCollectionMap`), which is pure noise from automated rename passes.

#### B. Vague descriptor suffixes (≈60+ files)
Words such as `_Part1`, `_Part2`, `_Methods`, `_Core`, `_Full`, `_Basic`, `_Extended`, `_Comprehensive`, `_Remaining`, `_Mid`, `_End`, `_Gaps`, `_MoreGaps`, `_FinalGaps`, `_AllGaps`. They tell us *nothing* about what behaviour is verified.

Examples: `Collection_Part1_test.go`, `Collection_Part2_test.go`, `Hashmap_BasicOps_test.go`, `Utils_WrapDouble_MoreGaps_test.go`, `SSO_Json_RemainingGaps_test.go`.

#### C. Coverage-motivated test-symbol suffixes (38 unique symbols, 1 file)
`SimpleSlice_Cap_test.go` still contains 38 `Test_*_Cov2` functions (e.g., `Test_SimpleSlice_Cap_Cov2`, `Test_AnyToString_Nil_Cov2`). The `_Cov2` token violates the project rule against `cov*` / `coverage*` markers in test names (see `.lovable/strictly-avoid.md`).

#### D. Mixed dual-naming inside a single file
Names like `Collection_IndexAt_Collection_Full_test.go` and `LeftRightFromSplit_Basic_S02_FromSplit_test.go` fuse two different naming schemes into one file, making search and ownership unclear.

### 1.2 Test-Structure Issues

1. **Inconsistent file/testcase pairing.** The guideline (`spec/testing-guidelines/01-folder-structure.md`) says every `Foo_test.go` should have a sibling `Foo_testcases.go`. Currently there are 180 test files but only 29 testcase files — the majority of test files inline their data, breaking the data/logic separation rule.
2. **Shared helper sprawl.** `Helpers_Methods_test.go`, `DataModels_Helpers_test.go`, `CloneSlice_Nil_Helpers_test.go`, `Helpers_testcases.go`, `Shared_testcases.go`, `Extended_testcases.go`, `Iteration_testcases.go`, `Remaining_testcases.go`, `Full_testcases.go`, `shared_compat_helpers.go`, `trydo_safe_test_helper_test.go` — overlapping responsibilities, no single owner per type.
3. **`DeadCode_doc.go`** — placeholder doc file with no clear purpose; should be removed or repurposed into a real package-doc file.
4. **Split-recovery residue.** The `S##` and `Seg#` files exist because Phase 1 split a monolithic file for compile-isolation; they should be merged back into per-function files now that compile is green.
5. **No `NilReceiver_test.go` consolidation.** Nil-safety is scattered across `BugfixRegression_NilReceiver_testcases.go`, `Hashset_NilReceiver_testcases.go`, `CloneSlice_Nil_Helpers_test.go`, and `NilReceiver_test.go` instead of a single canonical `NilReceiver_test.go` per the guideline.

### 1.3 Assertion-Usage Issues

1. **No `t.Log` / `t.Logf` calls** inside `corestrtests` (good — already clean). The user concern about `t.logger` is therefore preventive: keep it that way.
2. **No raw `t.Error` / `t.Errorf` / `t.Fatal` / `t.Fatalf`** inside `corestrtests` (good).
3. **Risk area:** any *new* tests added during the cleanup must continue to use framework assertions (`CaseV1.ShouldBeEqual`, `ShouldBeEqualMap`, `ShouldBeSafe`, GoConvey `convey.So`) — never raw `t.Error`. This rule must be re-asserted in every phase.
4. **`safeTest(...)` wrapper**: heavily used (e.g., `SimpleSlice_Cap_test.go`). This is acceptable but does not replace assertions — every `safeTest` body must end in a framework assertion.

### 1.4 Cross-Reference / Breakage Risk

Renaming a `Test_…` function or moving it to another file is safe (Go test discovery is by prefix), **but**:
- Any helper that is referenced by another test file must keep the same exported symbol or be updated everywhere.
- `_testcases.go` variables (e.g., `intSortQuickTestCases`) are package-scoped — renaming them requires updating every `_test.go` consumer in the same directory.
- `shared_compat_helpers.go` exports symbols consumed by sibling tests — verify before renaming.
- The split-recovery PowerShell tooling (`scripts/CoverageSplitRecovery.psm1`) keys off file names; large file rename batches must run with `run.ps1 -tc` after each batch to prove compile + coverage stay green.

---

## 2. Phased Improvement Plan

**Total: 7 phases.** Each phase is independently shippable, gated by `run.ps1 -tc` returning ✓ READY TO COMMIT and coverage ≥ current 98.5% for `coredata/corestr`.

| Phase | Title | Scope | Files touched (est.) | Exit criteria |
|-------|-------|-------|---------------------|---------------|
| **1** | Strip `_Cov<N>` symbols | Rename 38 `Test_*_Cov2` functions in `SimpleSlice_Cap_test.go` to descriptive `Test_<Function>_<Behaviour>` names. Update internal `safeTest("…")` strings to match. | 1 | 0 occurrences of `Cov[0-9]+` in any test symbol; `run.ps1 -tc` ✓; coverage unchanged |
| **2** | Remove cryptic prefix files (S##, Seg#) | Rename or merge the 27 `S##_…` and `Seg#_…` files into existing per-function files. Strip the duplicate-token noise (`S01_…_S01_…`). Update any cross-file helper references. | ~27 renamed, ~5 merged | No filename starts with `S\d` or `Seg\d`; `run.ps1 -tc` ✓ |
| **3** | Resolve initialism prefixes (HM/HS/SS/LL/LC/VV/KVP/SSO/CCM/CHM/COC/LR) | Rename 13 files to use the full type name (`HM01_IsEmpty_Hashmap_I8_test.go` → `Hashmap_IsEmpty_Int8_test.go`, `KVP_Basic_test.go` → `KeyValuePair_Basic_test.go`, etc.). | ~13 | All filenames begin with the full Go type name; `run.ps1 -tc` ✓ |
| **4** | Eliminate vague descriptor suffixes | Replace `_Part1/_Part2/_Core/_Full/_Methods/_Basic/_Extended/_Comprehensive/_Remaining/_Mid/_End/_Gaps/_MoreGaps/_FinalGaps/_AllGaps` with behaviour-describing suffixes (e.g., `_Part1` → split into `_Add`, `_Clone`, `_Length`, etc., per the actual tests inside). Where the file truly tests "all methods", rename to `_AllMethods_Smoke`. | ~60 | Grep for the listed banned suffixes returns 0; `run.ps1 -tc` ✓ |
| **5** | Restore data/logic separation | For every `_test.go` lacking a sibling `_testcases.go`, extract its inline `args.Map`/struct literals into the matching `_testcases.go`. Apply AAA comments. | ~150 | Every `_test.go` either has a sibling `_testcases.go` or is documented as no-data (e.g., panic-recovery). |
| **6** | Consolidate helpers and nil-safety | Merge `Helpers_Methods_test.go`, `DataModels_Helpers_test.go`, `CloneSlice_Nil_Helpers_test.go`, `Helpers_testcases.go`, `Shared_testcases.go`, `Extended_testcases.go`, `Iteration_testcases.go`, `Remaining_testcases.go`, `Full_testcases.go` into purpose-named files (`testhelpers.go`, `nilreceiver_helpers.go`). Move all nil-receiver tests under a single `NilReceiver_test.go` + per-type `<Type>_NilReceiver_testcases.go`. Remove `DeadCode_doc.go` (or convert to `doc.go`). | ~12 | Single helpers file per concern; `run.ps1 -tc` ✓ |
| **7** | Assertion guardrail + lint | Add a CI lint (extend the existing in-package import check in `run.ps1`) that fails the build if any test under `tests/integratedtests/` uses `t.Log`, `t.Logf`, `t.Error`, `t.Errorf`, `t.Fatal`, `t.Fatalf`, or any test symbol matches `Cov\d+|Coverage\d+`, `_Part\d+`, `_Seg\d+`, `^S\d+_`. | 1 PS module + 1 spec doc | Lint phase added; intentional violation fails fast |

### Per-phase Workflow (applies to every phase)

1. **Read** every file before editing.
2. **Search** for cross-references with `code--search_files` before any rename.
3. **Rename** in batches of ≤10 files.
4. **Compile-check** with `go build ./...` (or run `run.ps1 -tc`) after each batch.
5. **Run coverage** at the end of the phase; require coverage for `coredata/corestr` to stay ≥ 98.5%.
6. **Mark phase done** in this file's status table.
7. **Bump minor version** in `CHANGELOG`.
8. **Update memory** (`mem://workflow/01-coverage-fix-progress`) at the end of each phase.

### Phase Status Tracker

| Phase | Status | Started | Finished | Coverage delta |
|-------|--------|---------|----------|----------------|
| 1 — Strip `_Cov<N>` | ✅ Code done — pending `run.ps1 -tc` | 2026-04-18 | 2026-04-18 | expected unchanged (rename only) |
| 2 — Remove S##/Seg# | ⬜ Not started | — | — | — |
| 3 — Resolve initialisms | ⬜ Not started | — | — | — |
| 4 — Vague suffixes | ⬜ Not started | — | — | — |
| 5 — Data/logic split | ⬜ Not started | — | — | — |
| 6 — Helpers + nil-safety | ⬜ Not started | — | — | — |
| 7 — Lint guardrail | ⬜ Not started | — | — | — |

---

## 3. Out of Scope (this plan)

- Other test directories under `tests/integratedtests/` (will get a sibling plan once corestrtests is done).
- Production code under `coredata/corestr/` — only test files are renamed.
- Coverage *increase* work — the goal here is structural cleanup at constant coverage.

## 4. Open Questions for the User

1. For Phase 4, when a `_Part1/_Part2` file genuinely groups N unrelated tiny tests, do you prefer (a) split into N tiny files, or (b) merge into a single `_Smoke_test.go`?
2. For Phase 7, should the lint be a hard failure or a warning during a grace period?

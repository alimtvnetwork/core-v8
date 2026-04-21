# Project Overview

## Name
`github.com/alimtvnetwork/core-v8` — Go shared utility framework (auk-go ecosystem)

## Purpose
Reusable primitives, data structures, interfaces, converters, validators, file-system helpers, testing utilities, and code-generation tools for downstream Go packages.

## Tech Stack
- **Language:** Go 1.24+
- **Testing:** Custom `coretests` framework with `CaseV1`, `CaseNilSafe`, `GenericGherkins`
- **CI/CD:** GitHub Actions + GoReleaser
- **Tooling:** PowerShell `run.ps1` dispatcher for testing, coverage, and deployment
- **Cross-compile:** Windows, Linux, Darwin (amd64, arm64)

## Current State (as of session start)
- **182 blocked test packages** (compilation errors preventing coverage)
- **Coverage:** 58.7% overall (target: 100% reachable)
- **Major regression:** `corestrtests` — 176 files blocked, `coredata/corestr` at 8.4%
- **Other blocked:** `errcoretests` (3 files), `corepayloadtests` (2 files), `coredynamictests` (1 file)
- **Root cause:** Phase 1 cleanup (suffix stripping, function renames) caused 900+ function name redeclaration collisions

## Key Conventions
- Tests in `/tests/integratedtests/<package>tests/`
- AAA pattern (Arrange, Act, Assert)
- `args.Map` + `ShouldBeEqual` assertion style
- No boolean flag parameters
- Pointer receivers for nil-safety
- Never modify exported APIs
- `.release` folder is read-only

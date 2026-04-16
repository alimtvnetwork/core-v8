---
name: macos runner fix workflow
description: Solution record for the macOS PowerShell runner escape-sequence bug and error classification cleanup.
type: workflow
---
Solution applied for macOS `run.ps1 -tc` issue:

Files changed:
- `scripts/DashboardTheme.psm1`
- `scripts/CoverageCompileCheck.psm1`
- `scripts/CoverageReportJson.psm1`
- `scripts/CoverageRunner.psm1`

Fix summary:
- Added interactive-console guards before probing terminal background color with OSC 11.
- Restored terminal state safely with `finally` after `stty` usage.
- Removed compile-blocked package insertion into runtime failure aggregation during pre-coverage compile checks.
- Updated runtime failure report generation to exclude blocked packages and only report real runtime failures or missing coverage profiles.

Reason this memory exists:
- The user explicitly wants root cause analysis preserved before/after fixes.
- Future sessions should first classify failures correctly instead of treating all package problems as runtime crashes.

Verification target for future run:
- Run `./run.ps1 -tc` on macOS.
- Confirm no trailing `]11;rgb...` shell error appears.
- Confirm blocked compile packages remain only in `blocked-packages.txt` / `build-errors.txt`.
- Confirm `runtime-failures.txt` contains only true runtime failures.
